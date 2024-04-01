package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/onuraltuntasb/e-commerce/internal/auth"
	"github.com/onuraltuntasb/e-commerce/internal/driver"
	"github.com/onuraltuntasb/e-commerce/internal/models"
	"github.com/onuraltuntasb/e-commerce/internal/repository"
	"github.com/onuraltuntasb/e-commerce/internal/repository/dbrepo"
	"github.com/onuraltuntasb/e-commerce/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	DB repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgressRepo(db.SQL),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func AuthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "authCheck")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var requestPayload struct {
		FirstName       string `json:"firstName"`
		LastName        string `json:"lastName"`
		Password        string `json:"password"`
		PasswordConfirm string `json:"passwordConfirm"`
		Email           string `json:"email"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if requestPayload.Password != requestPayload.PasswordConfirm {
		utils.ErrorJSON(w, errors.New("given passwords are not matching"), http.StatusBadRequest)
		return
	}

	_, err = Repo.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {

		bytes, err := bcrypt.GenerateFromPassword([]byte(requestPayload.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
		}

		var user = &models.User{}
		user.FirstName = requestPayload.FirstName
		user.LastName = requestPayload.LastName
		user.Email = requestPayload.Email
		user.Password = string(bytes) // hashing password
		user.Status = 0               // only customer
		user.IsSeller = false
		user.Address = nil
		user.AuthType = 0 // only jwt
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		userDbId, err := Repo.DB.SaveUser(user)
		if err != nil {

			utils.ErrorJSON(w, err, http.StatusBadRequest)
		}

		u := auth.JwtUser{
			ID:        userDbId,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}

		tokens, err := auth.GenerateTokenPair(&u)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
		}

		// refresh token set cookie
		refreshCookie := auth.GetRefreshTokenCookie(tokens.RefreshToken)
		http.SetCookie(w, refreshCookie)

		// fmt.Printf("user %v", u)
		_, err = Repo.DB.SaveRefreshToken(tokens.RefreshToken, u.ID)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		// access token set cookie
		accessCookie := auth.GetAccessTokenCookie(tokens.AccessToken)
		http.SetCookie(w, accessCookie)

		//TODO dont return user.password
		utils.WriteJSON(w, http.StatusAccepted, "success")
		//http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		utils.ErrorJSON(w, errors.New("user already signed up"), http.StatusBadRequest)
	}

}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// fmt.Printf("request payload : %v \n", requestPayload)

	//get user by email
	user, err := Repo.DB.GetUserByEmail(requestPayload.Email)

	if err != nil {
		utils.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	u := auth.JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		utils.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	tokens, err := auth.GenerateTokenPair(&u)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	// refresh token set cookie
	refreshCookie := auth.GetRefreshTokenCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	_, err = Repo.DB.SaveRefreshToken(tokens.RefreshToken, u.ID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	// access token set cookie
	accessCookie := auth.GetAccessTokenCookie(tokens.AccessToken)
	http.SetCookie(w, accessCookie)

	utils.WriteJSON(w, http.StatusAccepted, tokens)

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	for _, cookie := range r.Cookies() {
		if cookie.Name == auth.Auth.CookieRefreshTokenName {
			claims := &auth.Claims{}
			refreshToken := cookie.Value

			//parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(auth.Auth.Secret), nil
			})
			if err != nil {
				utils.ErrorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				utils.ErrorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := Repo.DB.GetUserByID(userID)
			if err != nil {
				utils.ErrorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := auth.JwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			tokenPairs, err := auth.GenerateTokenPair(&u)
			if err != nil {
				utils.ErrorJSON(w, errors.New("error generation tokens"), http.StatusUnauthorized)
				return
			}

			//TODO delete old refresh token and save new one

			http.SetCookie(w, auth.GetRefreshTokenCookie(tokenPairs.RefreshToken))
			utils.WriteJSON(w, http.StatusOK, tokenPairs)

		}
	}

}

func AccessToken(w http.ResponseWriter, r *http.Request) {

	// fmt.Printf("access token hit!!")

	w.Header().Set("Access-Control-Allow-Credentials", "true")

	for _, cookie := range r.Cookies() {
		if cookie.Name == auth.Auth.CookieAccessTokenName {
			claims := &auth.Claims{}
			accessToken := cookie.Value

			//parse the token to get the claims
			_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(auth.Auth.Secret), nil
			})
			if err != nil {
				utils.ErrorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				utils.ErrorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := Repo.DB.GetUserByID(userID)
			if err != nil {
				utils.ErrorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := auth.JwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			tokenPairs, err := auth.GenerateTokenPair(&u)
			if err != nil {
				utils.ErrorJSON(w, errors.New("error generation tokens"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, auth.GetAccessTokenCookie(tokenPairs.AccessToken))
			utils.WriteJSON(w, http.StatusOK, tokenPairs)

		}
	}
}

func RenewAccessTokenByRefreshToken(w http.ResponseWriter, r *http.Request) {

	// fmt.Printf("RenewAccessTokenByRefreshToken \n")

	w.Header().Set("Access-Control-Allow-Credentials", "true")

	for _, cookie := range r.Cookies() {
		if cookie.Name == auth.Auth.CookieRefreshTokenName {
			claims := &auth.Claims{}
			refreshToken := cookie.Value

			//parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(auth.Auth.Secret), nil
			})
			if err != nil {
				utils.ErrorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				utils.ErrorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := Repo.DB.GetUserByID(userID)
			if err != nil {
				utils.ErrorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := auth.JwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			}

			tokenPairs, err := auth.GenerateTokenPair(&u)
			if err != nil {
				utils.ErrorJSON(w, errors.New("error generation tokens"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, auth.GetAccessTokenCookie(tokenPairs.AccessToken))
			utils.WriteJSON(w, http.StatusOK, tokenPairs)

		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Credentials", "true")

	http.SetCookie(w, auth.GetExpiredTokenCookie(auth.Auth.CookieAccessTokenName))
	http.SetCookie(w, auth.GetExpiredTokenCookie(auth.Auth.CookieRefreshTokenName))

	utils.WriteJSON(w, http.StatusOK, "success")
}

func Hello(w http.ResponseWriter, r *http.Request) {

	// fmt.Printf(" hello cookies : %v", r.Cookies())

	for _, cookie := range r.Cookies() {
		if cookie.Name == auth.Auth.CookieAccessTokenName {
			// accessToken := cookie.Value

		}
	}

	utils.WriteJSON(w, http.StatusOK, "hello")
}

func GetAllCities(w http.ResponseWriter, r *http.Request) {

	cityList, err := Repo.DB.GetAllCities()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, cityList)
}

func GetAllDistrictsByCityId(w http.ResponseWriter, r *http.Request) {

	cityId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	districtList, err := Repo.DB.GetAllDistrictsByCityId(cityId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, districtList)
}

func GetAllNeighborhoodByDistrictId(w http.ResponseWriter, r *http.Request) {

	districtId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	districtList, err := Repo.DB.GetAllNeighborhoodByDistrictId(districtId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, districtList)
}

func AddNewAddress(w http.ResponseWriter, r *http.Request) {

	address := &models.Address{}

	err := utils.ReadJSON(w, r, &address)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	addressId, err := Repo.DB.SaveAddress(address)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	type returnType struct {
		Data      string `json:"data"`
		AddressId int    `json:"addressId"`
	}

	returnObj := &returnType{
		"success",
		addressId,
	}

	utils.WriteJSON(w, http.StatusOK, returnObj)
}

func GetAllAddressesByUserId(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))

	addressList, err := Repo.DB.GetAllAddressesByUserId(userId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, addressList)
}

func MakePrimaryAddress(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	addressId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := Repo.DB.MakePrimaryAddress(addressId, userId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "success")
}

func DeleteAddressUserIdAndId(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	addressId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := Repo.DB.DeleteAddressUserIdAndId(addressId, userId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	//TODO: SUCCESS OR FAIL DECIDE LATER
	utils.WriteJSON(w, http.StatusOK, "success")
}

func UpdateAddressByUserIdAndId(w http.ResponseWriter, r *http.Request) {

	address := &models.Address{}

	err := utils.ReadJSON(w, r, &address)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err = Repo.DB.UpdateAddressByUserIdAndId(address)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "success")
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := Repo.DB.GetAllCategories()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, categories)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	// fmt.Printf("add product hit!")

	//TODO: handle eof or more than one json
	var product interface{}
	err := json.NewDecoder(r.Body).Decode(&product)

	// err := utils.ReadJSON(w, r, &product)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	p := product.(map[string]interface{})
	// fmt.Printf("product : %v ", p)

	productId, err := Repo.DB.SaveProduct(p, userId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	type response struct {
		ProductId int
		UserId    int
	}
	res := &response{
		ProductId: productId,
		UserId:    userId,
	}

	utils.WriteJSON(w, http.StatusOK, res)
}

func GetAllProductsByUserId(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	// fmt.Printf("GetAllProductsByUserId hit!")

	productList, err := Repo.DB.GetAllProductsByUserId(userId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, productList)
}

func UpdateProductByUserIdAndId(w http.ResponseWriter, r *http.Request) {

	//TODO: someone can change other's product just only if i can make api public

	// fmt.Printf("UpdateProductByUserIdAndId hit!")

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))

	var product interface{}
	err := json.NewDecoder(r.Body).Decode(&product)

	// err := utils.ReadJSON(w, r, &product)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	p := product.(map[string]interface{})
	// fmt.Printf("product : %v ", p)

	_, err = Repo.DB.UpdateProductByUserIdAndId(p, id, userId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "success")
}

func DeleteProductUserIdAndId(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := Repo.DB.DeleteProductUserIdAndId(id, userId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "success")
}

func GetUserInfosById(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	// fmt.Printf("GetAllProductsByUserId hit!")

	user, err := Repo.DB.GetUserInfosById(userId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func MakeUserSeller(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))

	// fmt.Printf("userId : %d", userId)

	_, err := Repo.DB.MakeUserSeller(userId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "success")
}

func GetProductsByCategoryName(w http.ResponseWriter, r *http.Request) {

	categoryName := chi.URLParam(r, "categoryName")
	// fmt.Printf("categoryName %s", categoryName)

	productList, err := Repo.DB.GetProductByCategoryName(categoryName)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	// fmt.Printf("productList : %v", productList)

	utils.WriteJSON(w, http.StatusOK, productList)
}

func GetProductsByQuery(w http.ResponseWriter, r *http.Request) {
	categoryName := chi.URLParam(r, "categoryName")

	fmt.Printf("\n categoryName : %s ", categoryName)

	var searchableProperties interface{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&searchableProperties)
	//TODO:JSON READ ERROR

	if err != nil {
		fmt.Printf("\n err : %v ", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	s := searchableProperties.(map[string]interface{})
	fmt.Printf("\n var : %v ", s)

	searchQuery := "select id, product, user_id from product where (product -> 'product'-> 'searchable'-> 'categories')::text ilike " + "'%" + categoryName + "%' and "
	cProperty := ""
	cOptions := "any (array["
	cOptionsArr := ""
	cQuery := "(product -> 'product'-> 'searchable'-> " + cProperty + " )::text ilike " + cOptions

	iteration := 0

	//TODO: make key-property control to prevent sql injection, read searchables
	//from one spesific table and check. It's kinda data sanitization.

	for key, value := range s {
		// fmt.Println("[", value, "] has items:")
		for k, v := range value.(map[string]interface{}) {

			if fmt.Sprintf("%v", v) != "[]" {
				// fmt.Println("\t-->", v, ":", v)

				// fmt.Println("\t-->", "k", ":", k)

				if k != "categories" && k != "price" && k != "name" {
					cProperty = k

					for _, elV := range v.([]interface{}) {
						//any (array['%female%','%male%'])
						cOptionsArr = cOptionsArr + "'%" + fmt.Sprintf("%v", elV) + "%',"
						// fmt.Println("\t-->", elV, ":", elV)

					}
					cOptions = cOptions + strings.TrimSuffix(cOptionsArr, ",") + "])"
					cQuery = " (product -> 'product'-> 'searchable'-> " + "'" + cProperty + "'" + " )::text ilike " + cOptions
					// fmt.Println("\t-->", "cQuery", ":", cQuery)

					cOptions = "any (array["
					cOptionsArr = ""

					fmt.Println("\t-->", "iteration", ":", iteration)

					if iteration == 0 {
						fmt.Println("[", key, "] has items:")
						searchQuery = searchQuery + cQuery
					} else {
						searchQuery = searchQuery + cQuery + " and "
					}
					iteration++
				}

			}

		}

	}

	searchQuery = strings.TrimSuffix(searchQuery, " and ")

	fmt.Println("\t-->", "searchQuery", ":", searchQuery)

	productList, err := Repo.DB.GetProductsByQuery(searchQuery)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	//TODO: return just productList OR?

	utils.WriteJSON(w, http.StatusOK, productList)

}

func GetProductById(w http.ResponseWriter, r *http.Request) {

	productId, _ := strconv.Atoi(chi.URLParam(r, "productId"))

	product, err := Repo.DB.GetProductById(productId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}

func GetProductCommentsByProductId(w http.ResponseWriter, r *http.Request) {

	productId, _ := strconv.Atoi(chi.URLParam(r, "productId"))

	commentList, err := Repo.DB.GetProductCommentsByProductId(productId)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, commentList)
}
