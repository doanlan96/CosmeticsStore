package repository

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
	"fmt"
)

func GetProducts(sortType string, prop string, limit int, offset int) (products []models.Product) {
	query, err := database.DB.Query("")
	if sortType == "ASC" && prop == "name" {
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id ORDER BY P.product_name ASC LIMIT ? OFFSET ?", limit, offset)
	} else if sortType == "DESC" && prop == "name" {
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id ORDER BY P.product_name DESC LIMIT ? OFFSET ?", limit, offset)
	} else if sortType == "ASC" && prop == "price" {
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id ORDER BY P.price ASC LIMIT ? OFFSET ?", limit, offset)
	} else if sortType == "DESC" && prop == "price" {
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id ORDER BY P.price DESC LIMIT ? OFFSET ?", limit, offset)
	} else {
		query, err = database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id LIMIT ? OFFSET ?", 20, 1)

	}
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var product models.Product
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			err.Error()
		}
		products = append(products, product)
	}
	return products
}

func SearchProducts(name string) (products []models.Product) {
	query, err := database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id WHERE P.product_name LIKE ?", "%"+name+"%")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	for query.Next() {
		var product models.Product
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	return products
}

func GetProductByCategoryName(name string) (products []models.Product) {
	query, err := database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id WHERE C.category_name = ?", name)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var product models.Product
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			err.Error()
		}
		products = append(products, product)
		fmt.Println(1)
		fmt.Println(products)
	}
	return products
}

func GetProductByBrandName(name string) (products []models.Product) {
	query, err := database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id WHERE C.category_name = ?", name)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var product models.Product
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			err.Error()
		}
		products = append(products, product)
	}
	return products
}

func GetProduct(id int) (product models.Product) {
	query, err := database.DB.Query("SELECT P.*, C.category_name, B.brand_name FROM products P JOIN categories C ON C.id = P.category_id JOIN brands B ON B.id = P.brand_id WHERE P.id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		err := query.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Image, &product.IsHot, &product.CategoryId, &product.BrandId, &product.NumberAvailable, &product.CategoryName, &product.BrandName)
		if err != nil {
			err.Error()
		}
	}
	return product
}

func CreateProduct(product models.Product) error {
	stmt, err := database.DB.Prepare("INSERT INTO products(product_name, description, price, image, is_hot, category_id, brand_id, number_available) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		err.Error()
	}
	_, err = stmt.Exec(product.ProductName, product.Description, product.Price, product.Image, product.IsHot, product.CategoryId, product.BrandId, product.NumberAvailable)
	if err != nil {
		err.Error()
	}
	return err
}

func UpdateProduct(product models.Product, id int) error {
	stmt, err := database.DB.Prepare("UPDATE products SET product_name=?,description=?,price=?,image=?,is_hot=?,category_id=?,brand_id=?,number_available=? WHERE id = ?;")
	if err != nil {
		err.Error()
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ProductName, product.Description, product.Price, product.Image, product.IsHot, product.CategoryId, product.BrandId, product.NumberAvailable, id)
	if err != nil {
		err.Error()
	}
	return err
}

func DeleteProduct(id int) error {
	query, err := database.DB.Query("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	return err
}
