package dbrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/onuraltuntasb/e-commerce/internal/models"
)

const dbTimeout = time.Second * 3

func (m *postgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from users where email = $1`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)

	//db column order is important !!!
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Email,
		&user.Status,
		&user.IsSeller,
		&user.AuthType,
		&user.Address,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *postgresDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	query := `select * from users where id = $1`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, id)

	//db column order is important !!!
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Email,
		&user.Status,
		&user.IsSeller,
		&user.AuthType,
		&user.Address,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *postgresDBRepo) SaveUser(user *models.User) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `insert into users (first_name, last_name, password, email, status, is_seller, auth_type, address_id, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id`

	var newID int

	//db column order is important !!!
	err := m.DB.QueryRowContext(ctx, stmt,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Email,
		user.Status,
		user.IsSeller,
		user.AuthType,
		nil,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *postgresDBRepo) SaveRefreshToken(token string, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `insert into refresh_token (refresh_token, user_id, created_at, updated_at)
		values ($1, $2, $3, $4) returning id`

	var newID int

	//db column order is important !!!
	err := m.DB.QueryRowContext(ctx, stmt,
		token,
		userId,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *postgresDBRepo) GetAllCities() ([]models.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	query := `select * from city`

	var city models.City
	var cityList []models.City
	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&city.ID,
			&city.CityCode,
			&city.Name,
		)
		if err != nil {
			return nil, err
		}

		cityList = append(cityList, city)
	}

	return cityList, nil
}

func (m *postgresDBRepo) GetAllDistrictsByCityId(cityId int) ([]models.District, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	query := `select * from district where city_id = $1`

	var district models.District
	var districtList []models.District
	rows, err := m.DB.QueryContext(ctx, query, cityId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&district.ID,
			&district.CityId,
			&district.Name,
		)
		if err != nil {
			return nil, err
		}

		districtList = append(districtList, district)
	}

	return districtList, nil
}

func (m *postgresDBRepo) GetAllNeighborhoodByDistrictId(districtId int) ([]models.Neighborhood, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	query := `select * from neighborhood where district_id = $1`

	var neighborhood models.Neighborhood
	var neighborhoodList []models.Neighborhood
	rows, err := m.DB.QueryContext(ctx, query, districtId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&neighborhood.ID,
			&neighborhood.DistrictId,
			&neighborhood.PostalCode,
			&neighborhood.Name,
		)
		if err != nil {
			return nil, err
		}

		neighborhoodList = append(neighborhoodList, neighborhood)
	}

	return neighborhoodList, nil
}

func (m *postgresDBRepo) GetNeighborhoodById(neighborhoodId int) (*models.Neighborhood, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	query := `select * from neighborhood where id = $1`

	var neighborhood models.Neighborhood
	row := m.DB.QueryRowContext(ctx, query, neighborhoodId)

	err := row.Scan(
		&neighborhood.ID,
		&neighborhood.DistrictId,
		&neighborhood.PostalCode,
		&neighborhood.Name,
	)
	if err != nil {
		return nil, err
	}

	return &neighborhood, nil
}

func (m *postgresDBRepo) GetDistrictById(districtId int) (*models.District, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	query := `select * from district where id = $1`

	var district models.District
	row := m.DB.QueryRowContext(ctx, query, districtId)

	err := row.Scan(
		&district.ID,
		&district.CityId,
		&district.Name,
	)
	if err != nil {
		return nil, err
	}

	return &district, nil
}

func (m *postgresDBRepo) GetCityById(cityId int) (*models.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	query := `select * from city where id = $1`

	var city models.City
	row := m.DB.QueryRowContext(ctx, query, cityId)

	err := row.Scan(
		&city.ID,
		&city.CityCode,
		&city.Name,
	)
	if err != nil {
		return nil, err
	}

	return &city, nil
}

func (m *postgresDBRepo) SaveAddress(address *models.Address) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `insert into address (user_id, header, name, country, city_id, district_id, neighborhood_id, description, phone, bill_type, ssn, is_primary, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) returning id`

	var newID int

	//db column order is important !!!
	err := m.DB.QueryRowContext(ctx, stmt,
		address.UserId,
		address.Header,
		address.Name,
		address.Country,
		address.CityId,
		address.DistrictId,
		address.NeighborhoodId,
		address.Description,
		address.Phone,
		address.BillType,
		address.SSN,
		address.IsPrimary,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// TODO: GET CITY NAME DISTRICT NAME NEIGHBORHOOD NAME
func (m *postgresDBRepo) GetAllAddressesByUserId(userId int) ([]models.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	query := `select * from address where user_id = $1 order by id`

	var address models.Address
	var addressList []models.Address
	rows, err := m.DB.QueryContext(ctx, query, userId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&address.ID,
			&address.UserId,
			&address.Header,
			&address.Name,
			&address.Country,
			&address.CityId,
			&address.DistrictId,
			&address.NeighborhoodId,
			&address.Description,
			&address.Phone,
			&address.BillType,
			&address.SSN,
			&address.IsPrimary,
			&address.CreatedAt,
			&address.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// fmt.Printf("NeighborhoodId id : %v", address.NeighborhoodId)
		// fmt.Printf("DistrictId id : %v", address.DistrictId)
		// fmt.Printf("CityId id : %v", address.CityId)

		neighborhood, err := m.GetNeighborhoodById(address.NeighborhoodId)
		if err != nil {
			return nil, err
		}
		district, err := m.GetDistrictById(address.DistrictId)
		if err != nil {
			return nil, err
		}
		city, err := m.GetCityById(address.CityId)
		if err != nil {
			return nil, err
		}

		address.Neighborhood = neighborhood.Name
		address.District = district.Name
		address.City = city.Name

		addressList = append(addressList, address)
	}

	return addressList, nil
}

func (m *postgresDBRepo) MakePrimaryAddress(addressId, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `update address set is_primary = $1 where id = $2 returning id`

	var newID int

	//db column order is important !!!
	err := m.DB.QueryRowContext(ctx, stmt,
		true,
		addressId,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	addressIds, err := m.GetAddressIdsByUserId(userId)
	if err != nil {
		return 0, err
	}

	for _, v := range addressIds {
		if v.ID != addressId {
			_, err = m.MakePrimaryFalseAddress(v.ID, userId)
			if err != nil {
				return 0, err
			}
		}
	}

	return newID, nil
}

func (m *postgresDBRepo) MakePrimaryFalseAddress(addressId, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `update address set is_primary = $1 where id = $2 returning id`

	var newID int

	//db column order is important !!!
	err := m.DB.QueryRowContext(ctx, stmt,
		false,
		addressId,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *postgresDBRepo) GetAddressIdsByUserId(userId int) ([]models.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `select id from address where user_id = $1`

	var address models.Address
	var addressList []models.Address
	rows, err := m.DB.QueryContext(ctx, stmt, userId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&address.ID,
		)
		if err != nil {
			return nil, err
		}

		addressList = append(addressList, address)

	}

	return addressList, nil

}

func (m *postgresDBRepo) DeleteAddressUserIdAndId(addressId int, userId int) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `delete from address where id = $1 and user_id = $2`

	//db column order is important !!!
	_, err := m.DB.QueryContext(ctx, stmt,
		addressId, userId,
	)

	if err != nil {
		return -1, err
	}

	return addressId, nil
}

func (m *postgresDBRepo) UpdateAddressByUserIdAndId(address *models.Address) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `update address set user_id = $1, header = $2, name = $3, country = $4, city_id = $5, district_id = $6, neighborhood_id = $7, description = $8, phone = $9, bill_type = $10, ssn = $11,  updated_at = $12
	 where id = $13 returning id`

	var newID int

	//db column order is important !!!
	err := m.DB.QueryRowContext(ctx, stmt,
		address.UserId,
		address.Header,
		address.Name,
		address.Country,
		address.CityId,
		address.DistrictId,
		address.NeighborhoodId,
		address.Description,
		address.Phone,
		address.BillType,
		address.SSN,
		time.Now(),
		address.ID,
	).Scan(&newID)

	if err != nil {
		return -1, err
	}

	return newID, nil
}

func (m *postgresDBRepo) GetAllCategories() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	//db column order is important !!!
	query := `select categories from category where id = $1;`

	row := m.DB.QueryRowContext(ctx, query, 1)

	var jsonData []byte
	err := row.Scan(
		&jsonData,
	)

	if err != nil {
		return nil, err
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(jsonData, &dat); err != nil {
		return nil, err
	}

	return dat, nil
}

func (m *postgresDBRepo) SaveProduct(product map[string]interface{}, userId int) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// fmt.Printf(" \n save product : %v \n", product)

	//db column order is important !!!
	stmt := `insert into product (product, user_id, created_at, updated_at)
		values ($1,$2,$3,$4) returning id`

	var productId int

	//db column order is important !!!
	err := m.DB.QueryRowContext(ctx, stmt,
		product,
		userId,
		time.Now(),
		time.Now(),
	).Scan(&productId)

	if err != nil {
		return 0, err
	}

	return productId, nil
}

func (m *postgresDBRepo) GetAllProductsByUserId(userId int) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	//db column order is important !!!
	query := `select id, product, user_id from product where user_id = $1;`

	var productList []models.Product
	rows, err := m.DB.QueryContext(ctx, query, userId)

	if err != nil {
		return nil, err
	}

	var pro = &models.Product{}
	var jsonData []byte
	for rows.Next() {
		err := rows.Scan(
			&pro.ID,
			&jsonData,
			&pro.UserId,
		)
		if err != nil {
			return nil, err
		}
		var dat map[string]interface{}
		if err := json.Unmarshal(jsonData, &dat); err != nil {
			return nil, err
		}

		pro.Product = dat
		productList = append(productList, *pro)
	}

	// fmt.Printf("productList : %v", productList)

	return productList, nil
}

func (m *postgresDBRepo) UpdateProductByUserIdAndId(product map[string]interface{}, id int, userId int) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `update product set product = $1, updated_at = $2
		where id = $3 and user_id = $4 returning id`

	var newID int

	//db column order is important !!!
	err := m.DB.QueryRowContext(ctx, stmt,
		product,
		time.Now(),
		id,
		userId,
	).Scan(&newID)

	if err != nil {
		return -1, err
	}

	return newID, nil
}

func (m *postgresDBRepo) DeleteProductUserIdAndId(id int, userId int) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	stmt := `delete from product where id = $1 and user_id = $2`

	//db column order is important !!!
	_, err := m.DB.QueryContext(ctx, stmt,
		id, userId,
	)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *postgresDBRepo) GetUserInfosById(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select first_name, last_name, email, is_seller from users where id = $1`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, id)

	//db column order is important !!!
	err := row.Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.IsSeller,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *postgresDBRepo) MakeUserSeller(id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	fmt.Printf(" user id %d", id)

	//db column order is important !!!
	stmt := `update users set is_seller = $1 where id = $2 returning id`

	var newID int

	//db column order is important !!!
	err := m.DB.QueryRowContext(ctx, stmt,
		true,
		id,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *postgresDBRepo) GetProductByCategoryName(categoryName string) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	//db column order is important !!!
	query := `select id, product, user_id from product where (product -> 'product'-> 'searchable'-> 'categories')::text ilike '%' || $1 || '%';`

	var productList []models.Product
	rows, err := m.DB.QueryContext(ctx, query, categoryName)

	if err != nil {
		return nil, err
	}

	var pro = &models.Product{}
	var jsonData []byte
	for rows.Next() {
		err := rows.Scan(
			&pro.ID,
			&jsonData,
			&pro.UserId,
		)
		if err != nil {
			return nil, err
		}
		var dat map[string]interface{}
		if err := json.Unmarshal(jsonData, &dat); err != nil {
			return nil, err
		}

		pro.Product = dat
		productList = append(productList, *pro)
	}

	// fmt.Printf("productList : %v", productList)

	return productList, nil
}

func (m *postgresDBRepo) GetProductsByQuery(searchQuery string) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	//db column order is important !!!

	var productList []models.Product
	rows, err := m.DB.QueryContext(ctx, searchQuery)

	if err != nil {
		return nil, err
	}

	var pro = &models.Product{}
	var jsonData []byte
	for rows.Next() {
		err := rows.Scan(
			&pro.ID,
			&jsonData,
			&pro.UserId,
		)
		if err != nil {
			return nil, err
		}
		var dat map[string]interface{}
		if err := json.Unmarshal(jsonData, &dat); err != nil {
			return nil, err
		}

		pro.Product = dat
		productList = append(productList, *pro)
	}

	// fmt.Printf("productList : %v", productList)

	return productList, nil
}

func (m *postgresDBRepo) GetProductById(productId int) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	//db column order is important !!!
	query := `select id, product, user_id from product where id = $1`

	var pro models.Product
	row := m.DB.QueryRowContext(ctx, query, productId)

	//db column order is important !!!
	var jsonData []byte
	err := row.Scan(
		&pro.ID,
		&jsonData,
		&pro.UserId,
	)

	var dat map[string]interface{}
	if err := json.Unmarshal(jsonData, &dat); err != nil {
		return nil, err
	}

	pro.Product = dat

	if err != nil {
		return nil, err
	}

	return &pro, nil
}

func (m *postgresDBRepo) GetProductCommentsByProductId(productId int) (*[]models.ProductComment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//db column order is important !!!
	query := `select * from product_comment where product_id = $1`

	var productComment models.ProductComment
	var productCommentList []models.ProductComment

	rows, err := m.DB.QueryContext(ctx, query, productId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(
			&productComment.ID,
			&productComment.Description,
			&productComment.Star,
			&productComment.UserId,
			&productComment.ProductId,
			&productComment.Description,
			&productComment.CreatedAt,
			&productComment.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		productCommentList = append(productCommentList, productComment)
	}
	return &productCommentList, nil
}
