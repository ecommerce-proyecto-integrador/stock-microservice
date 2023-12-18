package models

type Tiendas struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

// Definición del struct para la tabla de stock
type Stocks struct {
	ID         int `json:"id"`
	ProductoID int `json:"producto_id"`
	TiendaID   int `json:"tienda_id"`
	Cantidad   int `json:"cantidad"`
}
