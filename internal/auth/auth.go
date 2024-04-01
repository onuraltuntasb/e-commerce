package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Auth = &AuthType{}

func NewAuth(a *AuthType) {
	Auth = a
}

type AuthType struct {
	Issuer                 string
	Audience               string
	Secret                 string
	TokenExpiry            time.Duration
	RefreshExpiry          time.Duration
	CookiePath             string
	CookieAccessTokenName  string
	CookieRefreshTokenName string
	CookieDomain           string
}

type JwtUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type TokenPairs struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Claims struct {
	Name string `json:"name,omitempty"`
	jwt.RegisteredClaims
}

func GenerateTokenPair(user *JwtUser) (TokenPairs, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s %s", user.FirstName, user.LastName, user.Email)
	claims["sub"] = fmt.Sprint(user.ID)
	claims["aud"] = Auth.Audience
	claims["iss"] = Auth.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	claims["exp"] = time.Now().UTC().Add(Auth.TokenExpiry).Unix()

	signedAccessToken, err := token.SignedString([]byte(Auth.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["name"] = fmt.Sprintf("%s %s %s", user.FirstName, user.LastName, user.Email)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()
	refreshTokenClaims["aud"] = Auth.Audience
	refreshTokenClaims["iss"] = Auth.Issuer
	refreshTokenClaims["typ"] = "JWT"
	refreshTokenClaims["exp"] = time.Now().UTC().Add(Auth.RefreshExpiry).Unix()

	signedRefreshToken, err := refreshToken.SignedString([]byte(Auth.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	tokenPairs := TokenPairs{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return tokenPairs, nil

}

func GetAccessTokenCookie(token string) *http.Cookie {

	return &http.Cookie{
		Name:     Auth.CookieAccessTokenName,
		Path:     Auth.CookiePath,
		Value:    token,
		Expires:  time.Now().Add(Auth.TokenExpiry),
		MaxAge:   int(Auth.TokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   Auth.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func GetRefreshTokenCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     Auth.CookieRefreshTokenName,
		Path:     Auth.CookiePath,
		Value:    token,
		Expires:  time.Now().Add(Auth.RefreshExpiry),
		MaxAge:   int(Auth.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   Auth.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func GetExpiredTokenCookie(tokenCookieName string) *http.Cookie {
	return &http.Cookie{
		Name:     tokenCookieName,
		Path:     Auth.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   Auth.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func GetTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *Claims, error) {

	w.Header().Add("Vary", "Authorization")

	var at = ""
	var rt = ""
	var isAtExist = false
	var isRtExist = false

	// fmt.Printf("\n GetTokenFromHeaderAndVerify hit!! \n")

	for i := 0; i < len(r.Cookies()); i++ {
		if r.Cookies()[i].Name == Auth.CookieAccessTokenName || r.Cookies()[i].Name == Auth.CookieRefreshTokenName {
			if r.Cookies()[i].Name == Auth.CookieAccessTokenName {
				isAtExist = true
				at = r.Cookies()[i].Value
			}
			if r.Cookies()[i].Name == Auth.CookieRefreshTokenName {
				isRtExist = true
				rt = r.Cookies()[i].Value
			}
			// fmt.Printf("%s : %v \n", r.Cookies()[i].Name, r.Cookies()[i].Value)
		}
	}

	if isAtExist && isRtExist {

		claims := &Claims{}

		_, err := jwt.ParseWithClaims(at, claims, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(Auth.Secret), nil
		})

		if err != nil {
			if strings.HasPrefix(err.Error(), "token is expired by") {
				return "", nil, errors.New("expired token")
			}
			return "", nil, err
		}

		if claims.Issuer != Auth.Issuer {
			return "", nil, errors.New("invalid issuer")
		}

		return at, claims, nil

	} else if !isAtExist && isRtExist {
		rtClaims := &Claims{}

		_, err := jwt.ParseWithClaims(rt, rtClaims, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(Auth.Secret), nil
		})

		if err != nil {
			if strings.HasPrefix(err.Error(), "token is expired by") {
				return "", nil, errors.New("expired refresh token so session is expired")
			}
			return "", nil, err
		}

		if rtClaims.Issuer != Auth.Issuer {
			return "", nil, errors.New("invalid issuer")
		}

		nAt := jwt.New(jwt.SigningMethodHS256)

		nAtClaims := nAt.Claims.(jwt.MapClaims)

		nAtClaims["name"] = rtClaims.Name
		nAtClaims["sub"] = rtClaims.Subject
		nAtClaims["aud"] = rtClaims.Audience
		nAtClaims["iss"] = rtClaims.Issuer
		nAtClaims["iat"] = rtClaims.IssuedAt
		nAtClaims["typ"] = "JWT"
		nAtClaims["exp"] = time.Now().UTC().Add(Auth.TokenExpiry).Unix()

		sNat, err := nAt.SignedString([]byte(Auth.Secret))
		if err != nil {
			return at, rtClaims, err
		}

		accessCookie := GetAccessTokenCookie(sNat)

		http.SetCookie(w, accessCookie)

		return at, rtClaims, nil

	} else if !isRtExist {
		return "", nil, errors.New("there is no refresh token")
	} else {
		return "", nil, errors.New("there are no jwt token")
	}

}
