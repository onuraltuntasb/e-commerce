package repository

import "github.com/onuraltuntasb/e-commerce/internal/models"

type DatabaseRepo interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	SaveUser(user *models.User) (int, error)
	SaveRefreshToken(token string, userId int) (int, error)
	GetAllCities() ([]models.City, error)
	GetAllDistrictsByCityId(cityId int) ([]models.District, error)
	GetAllNeighborhoodByDistrictId(districtId int) ([]models.Neighborhood, error)
	SaveAddress(address *models.Address) (int, error)
	GetAllAddressesByUserId(userId int) ([]models.Address, error)
	MakePrimaryAddress(addressId, userId int) (int, error)
	DeleteAddressUserIdAndId(addressId int, userId int) (int, error)
	UpdateAddressByUserIdAndId(address *models.Address) (int, error)
	GetAllCategories() (interface{}, error)
	SaveProduct(product map[string]interface{}, userId int) (int, error)
	GetAllProductsByUserId(userId int) ([]models.Product, error)
	UpdateProductByUserIdAndId(product map[string]interface{}, id int, userId int) (int, error)
	DeleteProductUserIdAndId(id int, userId int) (int, error)
	GetUserInfosById(id int) (*models.User, error)
	MakeUserSeller(id int) (int, error)
	GetProductByCategoryName(categoryName string) ([]models.Product, error)
	GetProductsByQuery(searchQuery string) ([]models.Product, error)
	GetProductById(productId int) (*models.Product, error)
	GetProductCommentsByProductId(productId int) (*[]models.ProductComment, error)
}
