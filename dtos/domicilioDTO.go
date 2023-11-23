package dtos

import "math/big"

type DomicilioDTO struct {
	CalleRuta           string   `json:"calleRuta" binding:"required"`
	NumeroKM            string   `json:"numeroKM"`
	Interseccion1       string   `json:"interseccion1"`
	Interseccion2       string   `json:"interseccion2"`
	Piso                string   `json:"piso"`
	Depto               string   `json:"depto"`
	Barrio              string   `json:"barrio"`
	DepartamentoPartido string   `json:"departamentoPartido"`
	Localidad           string   `json:"localidad"`
	CodigoPostal        string   `json:"codigoPostal" binding:"required"`
	Provincia           int64    `json:"provincia" binding:"required"`
	Geolocalizacion     string   `json:"geolocalizacion"`
	Telefono            *big.Int `json:"telefono"`
	CorreoElectronico   string   `json:"correoElectronico"`
	CampoReservadoSSN9  string   `json:"campoReservadoSSN9"`
}
