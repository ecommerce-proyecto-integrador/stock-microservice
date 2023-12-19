package models

type Tiendas struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

// Definición del struct para la tabla de stock
type Stocks struct {
	ID         int `json:"id"`
	ProductoID int `json:"ProductoID"`
	TiendaID   int `json:"TiendaID"`
	Cantidad   int `json:"Cantidad"`
}

type StockResponse struct {
	Id            string `json:"id"`
	Tienda_nombre string `json:"tienda_nombre"`
	Cantidad      int    `json:"cantidad"`
}
