package controllers

import (
	"log"

	db "github.com/ecommerce-proyecto-integrador/stock-microservice/mod/config"
	"github.com/ecommerce-proyecto-integrador/stock-microservice/mod/models"
)

func GetStockByProductId(Productid string) ([]models.Stocks, error) {
	var stock []models.Stocks
	err := db.DB.
		Table("Stocks").
		Select("Stocks.id, Stocks.producto_id, Stocks.tienda_id, Stocks.cantidad, Tiendas.id as tienda_id, Tiendas.nombre as tienda_nombre").
		Joins("LEFT JOIN Tiendas ON Stocks.tienda_id = Tiendas.id").
		Where("Stocks.producto_id = ?", Productid).
		Find(&stock).Error
	log.Println(stock)
	return stock, err
}

/*func GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := db.DB.Preload("Images").Select("id, name, price, brand, size_available, reviews, category_name").Order("name ASC").Find(&products).Error

	return products, err
}*/
