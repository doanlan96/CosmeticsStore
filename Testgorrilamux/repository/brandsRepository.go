package repository

import (
	"Testgorillamux/database"
	"Testgorillamux/models"
)

func GetBrands() (brands []models.Brand) {
	query, err := database.DB.Query("SELECT * FROM brands")
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		var brand models.Brand
		err := query.Scan(&brand.Id, &brand.BrandName)
		if err != nil {
			err.Error()
		}
		brands = append(brands, brand)
	}
	return brands
}

func GetBrand(id int) (brand models.Brand) {
	query, err := database.DB.Query("SELECT * FROM brands WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	for query.Next() {
		err := query.Scan(&brand.Id, &brand.BrandName)
		if err != nil {
			err.Error()
		}
	}
	return brand
}

func CreateBrand(brand models.Brand) error {
	stmt, err := database.DB.Prepare("INSERT INTO brands(brand_name) VALUES (?)")
	if err != nil {
		err.Error()
	}
	_, err = stmt.Exec(brand.BrandName)
	if err != nil {
		err.Error()
	}
	return err
}

func UpdateBrand(brand models.Brand, id int) error {
	stmt, err := database.DB.Prepare("UPDATE brands SET brand_name=? WHERE id = ?;")
	if err != nil {
		err.Error()
	}
	defer stmt.Close()
	_, err = stmt.Exec(brand.BrandName, id)
	if err != nil {
		err.Error()
	}
	return err
}

func DeleteBrand(id int) error {
	query, err := database.DB.Query("DELETE FROM brands WHERE id = ?", id)
	if err != nil {
		err.Error()
	}
	defer query.Close()
	return err
}
