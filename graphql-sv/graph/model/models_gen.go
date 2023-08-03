// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cliente struct {
	ID                string `json:"id"`
	Nombre            string `json:"nombre"`
	CorreoElectronico string `json:"correoElectronico"`
}

type NuevoCliente struct {
	Nombre            string `json:"nombre"`
	CorreoElectronico string `json:"correoElectronico"`
}

type NuevoProducto struct {
	Nombre    string  `json:"nombre"`
	Marca     string  `json:"marca"`
	Categoria string  `json:"categoria"`
	Precio    float64 `json:"precio"`
}

type Producto struct {
	ID        string  `json:"id"`
	Nombre    string  `json:"nombre"`
	Marca     string  `json:"marca"`
	Categoria string  `json:"categoria"`
	Precio    float64 `json:"precio"`
}
