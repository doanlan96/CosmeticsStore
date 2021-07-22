package repository

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
)

func GetOrderItems() (orderItems []models.OrderItem) {
	query, err := database.DB.Query("SELECT * FROM order_items")
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var orderItem models.OrderItem
		err := query.Scan(&orderItem.Id, &orderItem.UserId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.IsPaid)
		orderItem.Product = GetProduct(int(orderItem.ProductId))
		orderItem.User = GetUser(int(orderItem.UserId))
		if err != nil {
			err.Error()
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems
}

func GetOrderItem(id int) (orderItem models.OrderItem) {
	query, err := database.DB.Query("SELECT * FROM `order_items` WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		err := query.Scan(&orderItem.Id, &orderItem.UserId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.IsPaid)
		if err != nil {
			err.Error()
		}
	}
	return orderItem
}

func GetOrderItemsByUserId(id int) (orderItems []models.OrderItem) {
	query, err := database.DB.Query("SELECT * FROM `order_items` WHERE user_id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var orderItem models.OrderItem
		err := query.Scan(&orderItem.Id, &orderItem.UserId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.IsPaid)
		if err != nil {
			err.Error()
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems
}

func CreateOrderItem(orderItem models.OrderItem) error {
	stmt, err := database.DB.Prepare("INSERT INTO order_items(user_id, product_id, quantity, is_paid) VALUES (?,?,?,?)")
	if err != nil {
		err.Error()
	}
	_, err = stmt.Exec(orderItem.UserId, orderItem.ProductId, orderItem.Quantity, orderItem.IsPaid)
	if err != nil {
		err.Error()
	}
	return err
}

func UpdateOrderItem(orderItem models.OrderItem, id int) error {
	stmt, err := database.DB.Prepare("UPDATE order_items SET product_id=?, quantity=?, is_paid=? WHERE id = ?;")
	if err != nil {
		err.Error()
	}
	defer stmt.Close()
	_, err = stmt.Exec(orderItem.ProductId, orderItem.Quantity, orderItem.IsPaid, id)
	if err != nil {
		err.Error()
	}
	return err
}

func DeleteOrderItem(id int) error {
	query, err := database.DB.Query("DELETE FROM order_items WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	return err
}
