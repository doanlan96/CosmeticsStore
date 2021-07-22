package routes

import (
	"Testgorillamux/handler"
	"Testgorillamux/middlewares"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {

	//products
	router.HandleFunc("/api/products", handler.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/search", handler.SearchProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", handler.GetProduct).Methods("GET")
	router.HandleFunc("/api/products/cate/{category_name}", handler.GetProductByCategoryName).Methods("GET")
	router.HandleFunc("/api/products/brand/{brand_name}", handler.GetProductByBrandName).Methods("GET")
	// s := router.PathPrefix("").Subrouter()
	// s.Use(middlewares.JwtVerify)
	// router.Use(middlewares.JwtVerify)
	router.HandleFunc("/api/products", middlewares.JwtVerify(handler.CreateProduct)).Methods("POST")
	router.HandleFunc("/api/products/{id}", middlewares.JwtVerify(handler.UpdateProduct)).Methods("PUT")
	router.HandleFunc("/api/products/{id}", middlewares.JwtVerify(handler.DeleteProduct)).Methods("DELETE")
	router.HandleFunc("/api/users", middlewares.JwtVerify(handler.GetUsers)).Methods("GET")
	router.HandleFunc("/api/users/{id}", middlewares.JwtVerify(handler.GetUser)).Methods("GET")
	router.HandleFunc("/api/users/{id}", middlewares.JwtVerify(handler.DeleteUser)).Methods("DELETE")

	//authentication
	router.HandleFunc("/api/user", middlewares.LoginVerify(handler.User)).Methods("GET")
	router.HandleFunc("/api/user", middlewares.LoginVerify(handler.UserProfile)).Methods("PUT")
	router.HandleFunc("/api/user/password", middlewares.LoginVerify(handler.UserPassword)).Methods("PUT")
	router.HandleFunc("/api/register", handler.Register).Methods("POST")
	router.HandleFunc("/api/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/logout", handler.Logout).Methods("POST")

	//brands
	router.HandleFunc("/api/brands", handler.GetBrands).Methods("GET")
	router.HandleFunc("/api/brands", handler.CreateBrand).Methods("POST")
	router.HandleFunc("/api/brands/{id}", handler.GetBrand).Methods("GET")
	router.HandleFunc("/api/brands/{id}", handler.UpdateBrand).Methods("PUT")
	router.HandleFunc("/api/brands/{id}", handler.DeleteBrand).Methods("DELETE")

	//categories
	router.HandleFunc("/api/categories", handler.GetCategories).Methods("GET")
	router.HandleFunc("/api/categories", handler.CreateCategory).Methods("POST")
	router.HandleFunc("/api/categories/{id}", handler.GetCategory).Methods("GET")
	router.HandleFunc("/api/categories/{id}", handler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/api/categories/{id}", handler.DeleteCategory).Methods("DELETE")

	//orders
	router.HandleFunc("/api/orders", middlewares.JwtVerify(handler.GetOrderItems)).Methods("GET")
	router.HandleFunc("/api/orders", handler.CreateOrderItem).Methods("POST")
	router.HandleFunc("/api/orders/{id}", handler.GetOrderItem).Methods("GET")
	router.HandleFunc("/api/orders/user/{id}", handler.GetOrderItemsByUserId).Methods("GET")
	router.HandleFunc("/api/orders/{id}", handler.UpdateOrderItem).Methods("PUT")
	router.HandleFunc("/api/orders/{id}", middlewares.JwtVerify(handler.DeleteOrderItem)).Methods("DELETE")
	router.HandleFunc("/api/order/upsert", handler.UpSertOrder).Methods("POST")
}
