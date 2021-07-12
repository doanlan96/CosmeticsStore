package main

import (
	"Testgorillamux/database"
	"Testgorillamux/routes"
	"fmt"
	"github.com/rs/cors"
	"log"

	"net/http"

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


	routes.Setup(router)
	handler := cors.Default().Handler(router)

	fmt.Println("Server running at localhost:8000")

	http.ListenAndServe(":8000", handler)

}


