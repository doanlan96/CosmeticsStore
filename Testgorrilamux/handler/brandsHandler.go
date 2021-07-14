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

func GetBrands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var Brands []models.Brand
	result := repository.GetBrands()
	json.NewEncoder(w).Encode(result)
}

func CreateBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var brand models.Brand
	json.Unmarshal(body, &brand)
	errr := repository.CreateBrand(brand)
	if errr != nil {
		fmt.Fprintln(w, "Cannot create Brand!")
	}
	fmt.Fprintf(w, "New Brand was created")
	json.NewEncoder(w).Encode(brand)
}

func GetBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := repository.GetBrand(id)
	json.NewEncoder(w).Encode(result)
}

func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var brand models.Brand
	json.Unmarshal(body, &brand)
	errr := repository.UpdateBrand(brand, id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot update Brand!")
	}
	json.NewEncoder(w).Encode(brand)
	fmt.Fprintf(w, "Brand with ID = %s was updated", params["id"])
}

func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	errr := repository.DeleteBrand(id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete Brand!")
	}
	fmt.Fprintf(w, "Brand with ID = %s was deleted", params["id"])
}
