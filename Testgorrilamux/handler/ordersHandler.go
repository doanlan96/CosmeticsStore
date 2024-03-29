package handler

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
	"Testgorillamux/repository"
	"Testgorillamux/sendGrid"
	"Testgorillamux/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type resOrder struct {
	OrderItems []models.OrderItem
	Count      int
}

func GetOrderItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var orders []models.Order
	limit := 10
	params := r.URL.Query()
	offset, _ := strconv.Atoi(params.Get("offset"))
	result, count := repository.GetOrderItems(limit, offset)
	json.NewEncoder(w).Encode(resOrder{result, count})
}

func GetOrderItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println(1)
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
	orders, _ := database.DB.Query("SELECT * FROM order_items where user_id = ? AND is_paid = 0", id)
	var orderItems []models.OrderItem
	for orders.Next() {
		var order models.OrderItem
		orders.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.IsPaid, &order.Ship)
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

func UpSertOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err.Error()
	}
	var data map[string]int
	json.Unmarshal(body, &data)
	var quantity uint
	order := database.DB.QueryRow("SELECT quantity FROM order_items WHERE user_id = ? AND product_id = ? AND is_paid = 0", data["user_id"], data["product_id"])
	err = order.Scan(&quantity)
	if err == sql.ErrNoRows {
		database.DB.Exec("INSERT INTO order_items(user_id, product_id, quantity, is_paid) values (?, ?, ?, ?)", data["user_id"], data["product_id"], data["quantity"], data["is_paid"])
		return
	}
	_, err = database.DB.Exec("UPDATE order_items SET quantity = ? WHERE user_id = ? AND product_id = ? AND is_paid = 0", int(quantity)+data["quantity"], data["user_id"], data["product_id"])
	return
}

func DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var idd string
	var userId int
	cookie, _ := r.Cookie("userId")
	if cookie == nil {
		userId = 0
	} else {
		idd, _ = util.ParseJwt(cookie.Value)
		userId, _ = strconv.Atoi(idd)
	}
	errr := repository.DeleteOrderItem(id, userId)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete order!")
	}
	fmt.Fprintf(w, "order item with ID = %s was deleted", params["id"])
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	errr := repository.DeleteOrder(id)
	if errr != nil {
		fmt.Fprintln(w, "Cannot delete order!")
	}
	fmt.Fprintf(w, "order item with ID = %s was deleted", params["id"])
}




func ChangeStatusOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var value models.Data
	json.Unmarshal(body, &value)
	fmt.Println(value.Orders)
	for _, item := range value.Products {
		_, err := database.DB.Exec("update order_items set is_paid = 1 where user_id = ? and product_id = ? and is_paid = 0", value.UserId, item.Id)
		fmt.Println(err)
	}
	database.DB.Exec("INSERT INTO address_shipping(user_id, address, email, phone, full_name) values (?, ?, ?, ?, ?)",
		value.AddressShipping.UserId, value.AddressShipping.Address,value.AddressShipping.Email, value.AddressShipping.Phone, value.AddressShipping.FullName)
	sendGrid.SendEmailBySendGrid(value)
}