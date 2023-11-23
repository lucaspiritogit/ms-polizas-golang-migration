package models

import (
	"math/big"
	"time"
)

func (Domicilio) TableName() string {
	return "SSNT_GED.NEGO_DOMICILIO"
}

type Domicilio struct {
	IDDomicilio          int64     `json:"idDomicilio" gorm:"column:ID_DOMICILIO;type:bigint;autoIncrement;primaryKey"`
	IDPoliza             string    `json:"idPoliza" gorm:"column:ID_POLIZA;type:varchar(50);not null"`
	CalleRuta            string    `json:"calleRuta" gorm:"column:CALLE_RUTA;type:varchar(50);not null"`
	NumeroKM             string    `json:"numeroKM" gorm:"column:NUMERO_KM;type:varchar(50)"`
	Interseccion1        string    `json:"interseccion1" gorm:"column:INTERSECCION1;type:varchar(50)"`
	Interseccion2        string    `json:"interseccion2" gorm:"column:INTERSECCION2;type:varchar(50)"`
	Piso                 string    `json:"piso" gorm:"column:PISO;type:varchar(4)"`
	Depto                string    `json:"depto" gorm:"column:DEPTO;type:varchar(5)"`
	Barrio               string    `json:"barrio" gorm:"column:BARRIO;type:varchar(50)"`
	DepartamentoPartido  string    `json:"departamentoPartido" gorm:"column:DEPARTAMENTO_PARTIDO;type:varchar(50)"`
	Localidad            string    `json:"localidad" gorm:"column:LOCALIDAD;type:varchar(50)"`
	CodigoPostal         string    `json:"codigoPostal" gorm:"column:CODIGO_POSTAL;type:varchar(8);not null"`
	Provincia            int64     `json:"provincia" gorm:"column:PROVINCIA;type:bigint;not null"`
	Geolocalizacion      string    `json:"geolocalizacion" gorm:"column:GEOLOCALIZACION;type:varchar(50)"`
	Telefono             *big.Int  `json:"telefono" gorm:"column:TELEFONO;type:numeric(12)"`
	CorreoElectronico    string    `json:"correoElectronico" gorm:"column:CORREO_ELECTRONICO;type:varchar(50)"`
	CampoReservadoSSN9   string    `json:"campoReservadoSSN9" gorm:"column:CAMPO_RESERVADO_SSN9;type:varchar(30)"`
	FechaHoraCreacionSSN time.Time `json:"fechaHoraCreacionSSN" gorm:"column:FECHA_HORA_CREACION_SSN;type:timestamp;not null"`
}
