package repository

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
)

func GetOrders() (orders []models.Order) {
	query, err := database.DB.Query("SELECT O.id, O.user_id, OI.product_id, OI.quantity, O.payment, O.total FROM `orders` O JOIN order_items OI ON OI.order_id = O.id")
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var order models.Order
		err := query.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.Payment, &order.Total)
		if err != nil {
			err.Error()
		}
		orders = append(orders, order)
	}
	return orders
}

func GetOrder(id int) (orders []models.Order) {
	query, err := database.DB.Query("SELECT O.id, O.user_id, OI.product_id, OI.quantity, O.payment, O.total FROM `orders` O JOIN order_items OI ON OI.order_id = O.id WHERE O.id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var order models.Order
		err := query.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.Payment, &order.Total)
		if err != nil {
			err.Error()
		}
		orders = append(orders, order)
	}
	return orders
}

func CreateOrder(order models.Order) error {
	stmt, err := database.DB.Prepare("INSERT INTO Orders(user_id, payment, total) VALUES (?,?,?)")
	if err != nil {
		err.Error()
	}
	_, err = stmt.Exec(order.UserId, order.Payment, order.Total)
	if err != nil {
		err.Error()
	}
	return err
}

func CreateOrderItem(orderItem models.OrderItem) error {
	stmt, err := database.DB.Prepare("INSERT INTO order_items(order_id, product_id, quantity) VALUES (?,?,?)")
	if err != nil {
		err.Error()
	}
	_, err = stmt.Exec(orderItem.OrderId, orderItem.ProductId, orderItem.Quantity)
	if err != nil {
		err.Error()
	}
	return err
}

func UpdateOrder(order models.Order, id int) error {
	stmt, err := database.DB.Prepare("UPDATE orders SET user_id=?, payment=?, total=? WHERE id = ?;")
	if err != nil {
		err.Error()
	}
	defer stmt.Close()
	_, err = stmt.Exec(order.UserId, order.Payment, order.Total, id)
	if err != nil {
		err.Error()
	}
	return err
}

func DeleteOrder(id int) error {
	query, err := database.DB.Query("DELETE FROM orders WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
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
