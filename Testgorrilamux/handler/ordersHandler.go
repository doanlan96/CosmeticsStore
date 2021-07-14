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

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var orders []models.Order
	result := repository.GetOrders()
	json.NewEncoder(w).Encode(result)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var order models.Order
	json.Unmarshal(body, &order)
	errr := repository.CreateOrder(order)
	if errr != nil {
		fmt.Fprintln(w, "Cannot create Order!")
	}
	fmt.Fprintf(w, "New Order was created")
	json.NewEncoder(w).Encode(order)
}

func CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var orderItem models.OrderItem
	json.Unmarshal(body, &orderItem)
	errr := repository.CreateOrderItem(orderItem)
	if errr != nil {
		fmt.Fprintln(w, "Cannot create Order!")
	}
	fmt.Fprintf(w, "New OrderItem was created")
	json.NewEncoder(w).Encode(orderItem)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := repository.GetOrder(id)
	json.NewEncoder(w).Encode(result)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var order models.Order
	json.Unmarshal(body, &order)
	errr := repository.UpdateOrder(order, id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot update order!")
	}
	json.NewEncoder(w).Encode(order)
	fmt.Fprintf(w, "order with ID = %s was updated", params["id"])
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	errr := repository.DeleteOrder(id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete order!")
	}
	fmt.Fprintf(w, "order with ID = %s was deleted", params["id"])
}

func DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	errr := repository.DeleteOrderItem(id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete order!")
	}
	fmt.Fprintf(w, "order item with ID = %s was deleted", params["id"])
}
