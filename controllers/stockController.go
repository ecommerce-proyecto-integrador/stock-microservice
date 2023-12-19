package controllers

import (
	"errors"
	"log"

	db "github.com/ecommerce-proyecto-integrador/stock-microservice/mod/config"
	"github.com/ecommerce-proyecto-integrador/stock-microservice/mod/models"
	"gorm.io/gorm"
)

func GetStockByProductId(Productid string) ([]models.StockResponse, error) {
	var allTiendas []models.Tiendas
	err := db.DB.
		Table("tiendas").
		Find(&allTiendas).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var stock []models.StockResponse
	for _, tienda := range allTiendas {
		var stockEntry models.StockResponse
		err := db.DB.
			Table("stocks").
			Select("stocks.id, tiendas.nombre as tienda_nombre, COALESCE(stocks.cantidad, 0) as cantidad").
			Joins("LEFT JOIN tiendas ON stocks.tienda_id = tiendas.id").
			Where("stocks.producto_id = ? AND stocks.tienda_id = ?", Productid, tienda.ID).
			First(&stockEntry).Error

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the error
			log.Println(err)
			return nil, err
		}
		stockEntry.Id = "0"
		stockEntry.Tienda_nombre = tienda.Nombre
		stock = append(stock, stockEntry)
	}
	log.Println(stock)
	return stock, err
}

/*func GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := db.DB.Preload("Images").Select("id, name, price, brand, size_available, reviews, category_name").Order("name ASC").Find(&products).Error

	return products, err
}*/
