package handler

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
	"Testgorillamux/repository"
	"Testgorillamux/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetOrderItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var orders []models.Order
	result := repository.GetOrderItems()
	json.NewEncoder(w).Encode(result)
}

func GetOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := repository.GetOrderItem(id)
	json.NewEncoder(w).Encode(result)
}

func GetOrderItemsByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	orders, _ := database.DB.Query("SELECT * FROM order_items where user_id = ?", id)
	var orderItems []models.OrderItem
	for orders.Next() {
		var order models.OrderItem
		orders.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.IsPaid)
		order.Product = getProduct(order.ProductId)
		orderItems = append(orderItems, order)
	}
	json.NewEncoder(w).Encode(orderItems)
}

func getProduct(id uint) models.Product {
	rows, _ := database.DB.Query("SELECT * FROM products where id = ?", id)
	var product models.Product
	for rows.Next() {
		rows.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable)
	}
	return product
}


func CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err.Error()
	}
	var id string
	var userId int
	cookie, _ := r.Cookie("userId")
	if cookie == nil {
		userId = 0
	} else {
		id, _ = util.ParseJwt(cookie.Value)
		userId, _ = strconv.Atoi(id)
	}
	var data map[string]uint
	json.Unmarshal(body, &data)
	orderItem := models.OrderItem{
		UserId:    uint(userId),
		ProductId: data["product_id"],
		Quantity:  data["quantity"],
		IsPaid:    int(data["is_paid"]),
	}
	errr := repository.CreateOrderItem(orderItem)
	if errr != nil {
		fmt.Fprintln(w, "Cannot create Order!")
	}
	fmt.Fprintf(w, "New OrderItem was created")
	json.NewEncoder(w).Encode(orderItem)
}

func UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err.Error()
	}
	var orderItem models.OrderItem
	json.Unmarshal(body, &orderItem)
	errr := repository.UpdateOrderItem(orderItem, id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot update order!")
	}
	json.NewEncoder(w).Encode(orderItem)
	fmt.Fprintf(w, "order with ID = %s was updated", params["id"])
}

func UpSertOrder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err.Error()
	}
	var data map[string]int
	json.Unmarshal(body, &data)
	var quantity uint
	order := database.DB.QueryRow("SELECT quantity FROM order_items WHERE user_id = ? AND product_id = ?", data["user_id"], data["product_id"])
	err = order.Scan(&quantity)
	if err == sql.ErrNoRows {
		if data["replace"] > 0 {
			database.DB.Exec("INSERT INTO order_items(user_id, product_id, quantity, is_paid, total) values (?, ?, ?, ?, ?)",data["user_id"], data["product_id"], data["replace"], data["is_paid"] )
			return
		}
		if data["quantity"] > 0 {
			database.DB.Exec("INSERT INTO order_items(user_id, product_id, quantity, is_paid) values (?, ?, ?, ?)",data["user_id"], data["product_id"], data["quantity"], data["is_paid"] )
			return
		}

	}
	if data["replace"] == 0 && data["quantity"] != 0{
		fmt.Println(data["quantity"])
		if int(quantity) <= data["quantity"] {
			quantity = uint(data["quantity"])
		}
		_, err = database.DB.Exec("UPDATE order_items SET quantity = ? WHERE user_id = ? AND product_id = ?", int(quantity) + data["quantity"], data["user_id"], data["product_id"])
		return
	}
	_, err = database.DB.Exec("UPDATE order_items SET quantity = ? WHERE user_id = ? AND product_id = ?", data["replace"], data["user_id"], data["product_id"])
	return
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
