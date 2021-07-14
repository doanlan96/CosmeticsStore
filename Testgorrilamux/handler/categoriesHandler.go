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

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := repository.GetCategories()
	json.NewEncoder(w).Encode(result)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var category models.Category
	json.Unmarshal(body, &category)
	errr := repository.CreateCategory(category)
	if errr != nil {
		fmt.Fprintln(w, "Cannot create Category!")
	}
	fmt.Fprintf(w, "New category was created")
	json.NewEncoder(w).Encode(category)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := repository.GetCategory(id)
	json.NewEncoder(w).Encode(result)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var category models.Category
	json.Unmarshal(body, &category)
	errr := repository.UpdateCategory(category, id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot update Category!")
	}
	json.NewEncoder(w).Encode(category)
	fmt.Fprintf(w, "Category with ID = %s was updated", params["id"])
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	errr := repository.DeleteCategory(id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete Category!")
	}
	fmt.Fprintf(w, "Category with ID = %s was deleted", params["id"])
}
