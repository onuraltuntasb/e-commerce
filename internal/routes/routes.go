package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/onuraltuntasb/e-commerce/internal/handlers"
	im "github.com/onuraltuntasb/e-commerce/internal/middleware"
)

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(im.EnableCors)
	//mux.Use(im.NoSurf)

	mux.Post("/signup", handlers.Signup)
	mux.Post("/login", handlers.Authenticate)
	mux.Get("/categories", handlers.GetAllCategories)

	// mux.Get("/products-get/category/{categoryName}", handlers.GetProductsByCategoryName)

	mux.Post("/products-get/category/{categoryName}", handlers.GetProductsByQuery)

	mux.Get("/product-get/{productId}", handlers.GetProductById)
	mux.Get("/product-comments-get/{productId}", handlers.GetProductCommentsByProductId)

	//NOTE: don't forget to add "/auth path" before main path
	mux.Route("/auth", func(mux chi.Router) {
		mux.Use(im.AuthRequired)

		// mux.Get("/auth-check", handlers.AuthCheck)

		mux.Get("/renew-access-token", handlers.RenewAccessTokenByRefreshToken)
		mux.Get("/logout", handlers.Logout)

		mux.Get("/city", handlers.GetAllCities)
		mux.Get("/district/{id}", handlers.GetAllDistrictsByCityId)
		mux.Get("/neighborhood/{id}", handlers.GetAllNeighborhoodByDistrictId)

		mux.Post("/address-add", handlers.AddNewAddress)
		mux.Get("/addresses-get/{userId}", handlers.GetAllAddressesByUserId)
		mux.Post("/make-address-primary/{userId}/{id}", handlers.MakePrimaryAddress)
		mux.Delete("/address-delete/{userId}/{id}", handlers.DeleteAddressUserIdAndId)
		mux.Put("/address-update", handlers.UpdateAddressByUserIdAndId)

		mux.Post("/product-add/{userId}", handlers.AddProduct)
		mux.Get("/products-get-seller/{userId}", handlers.GetAllProductsByUserId)
		mux.Put("/product-update/{id}/{userId}", handlers.UpdateProductByUserIdAndId)
		mux.Delete("/product-delete/{id}/{userId}", handlers.DeleteProductUserIdAndId)

		mux.Get("/account-info-get/{userId}", handlers.GetUserInfosById)
		mux.Post("/make-account-seller/{userId}", handlers.MakeUserSeller)

	})

	return mux
}
