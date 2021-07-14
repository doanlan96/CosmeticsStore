package handler

import (
	"Testgorillamux/models"
	"Testgorillamux/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var products []models.Product
	params := r.URL.Query()
	sortType := params.Get("sortType")
	prop := params.Get("prop")
	limit, _ := strconv.Atoi(params.Get("limit"))
	offset, _ := strconv.Atoi(params.Get("offset"))
	result := repository.GetProducts(sortType, prop, limit, offset)
	json.NewEncoder(w).Encode(result)
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var data map[string]string
	json.Unmarshal(body, &data)
	result := repository.SearchProducts(data["key_word"])
	json.NewEncoder(w).Encode(result)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var product models.Product
	json.Unmarshal(body, &product)
	errr := repository.CreateProduct(product)
	if errr != nil {
		fmt.Fprintln(w, "Cannot create product!")
	}
	fmt.Fprintf(w, "New product was created")
	json.NewEncoder(w).Encode(product)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := repository.GetProduct(id)
	json.NewEncoder(w).Encode(result)
}

func GetProductByCategoryName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["category_name"]
	result := repository.GetProductByCategoryName(name)
	json.NewEncoder(w).Encode(result)
}

func GetProductByBrandName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["brand_name"]
	result := repository.GetProductByCategoryName(name)
	json.NewEncoder(w).Encode(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var product models.Product
	json.Unmarshal(body, &product)
	errr := repository.UpdateProduct(product, id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot update product!")
	}
	json.NewEncoder(w).Encode(product)
	fmt.Fprintf(w, "Product with ID = %s was updated", params["id"])
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	errr := repository.DeleteProduct(id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete product!")
	}
	fmt.Fprintf(w, "Product with ID = %s was deleted", params["id"])
}
