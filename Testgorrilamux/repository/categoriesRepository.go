package repository

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
)

func GetCategories() (categories []models.Category) {
	query, err := database.DB.Query("SELECT * FROM categories")
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var category models.Category
		err := query.Scan(&category.Id, &category.CategoryName)
		if err != nil {
			err.Error()
		}
		categories = append(categories, category)
	}
	return categories
}

func GetCategory(id int) (category models.Category) {
	query, err := database.DB.Query("SELECT * FROM categories WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		err := query.Scan(&category.Id, &category.CategoryName)
		if err != nil {
			err.Error()
		}
	}
	return category
}

func CreateCategory(category models.Category) error {
	stmt, err := database.DB.Prepare("INSERT INTO categories(category_name) VALUES (?)")
	if err != nil {
		err.Error()
	}
	_, err = stmt.Exec(category.CategoryName)
	if err != nil {
		err.Error()
	}
	return err
}

func UpdateCategory(category models.Category, id int) error {
	stmt, err := database.DB.Prepare("UPDATE categories SET category_name=? WHERE id = ?;")
	if err != nil {
		err.Error()
	}
	defer stmt.Close()
	_, err = stmt.Exec(category.CategoryName, id)
	if err != nil {
		err.Error()
	}
	return err
}

func DeleteCategory(id int) error {
	query, err := database.DB.Query("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	return err
}
