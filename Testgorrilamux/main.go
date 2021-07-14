
package main

import (
	"Testgorillamux/database"
	"Testgorillamux/routes"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.Connect() //Khởi tạo biến conection
	if err != nil {               //Catch error trong quá trình thực thi
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")

	router := mux.NewRouter()

	handleCross := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowCredentials(),
	)

	routes.Setup(router)
	//crawler.CrawlProduct(db)
	fmt.Println("Server running at localhost:8000")

	http.ListenAndServe(":8000", handleCross(router))

}
