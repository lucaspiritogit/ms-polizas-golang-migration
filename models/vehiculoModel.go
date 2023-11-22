package models

import (
	"time"
)

func (Vehiculo) TableName() string {
	return "SSNT_GED.NEGO_VEHICULOS"
}

type Vehiculo struct {
	Poliza
	Domicilio
	TipoVehiculo         string    `json:"tipoVehiculo" gorm:"column:TIPO_VEHICULO;type:varchar(2);not null"`
	Patente              string    `json:"patente" gorm:"column:PATENTE;type:varchar(20);not null"`
	UbicacionRiesgo      string    `json:"ubicacionRiesgo" gorm:"column:UBICACION_RIESGO;type:varchar(8);not null"`
	Marca                string    `json:"marca" gorm:"column:MARCA;type:varchar(25);not null"`
	Modelo               string    `json:"modelo" gorm:"column:MODELO;type:varchar(50);not null"`
	AnioFabricacion      int64     `json:"anioFabricacion" gorm:"column:ANIO_FABRICACION;type:bigint;not null"`
	NroMotor             string    `json:"nroMotor" gorm:"column:NRO_MOTOR;type:varchar(25)"`
	NroChasis            string    `json:"nroChasis" gorm:"column:NRO_CHASIS;type:varchar(25)"`
	CodSeguimiento       string    `json:"codSeguimiento" gorm:"column:CODIGO_SEGUIMIENTO;type:varchar(20)"`
	Suma                 float64   `json:"suma" gorm:"column:SUMA;type:numeric(21,2);not null"`
	TipoMoneda           int64     `json:"tipoMoneda" gorm:"column:TIPO_MONEDA;type:bigint;not null"`
	Prenda               string    `json:"prenda" gorm:"column:PRENDA;type:varchar(2);not null"`
	EstadoVehiculoCod    string    `json:"estadoVehiculoCod" gorm:"column:ESTADO_VEHICULO_CODIGO;type:varchar(1)"`
	NrOrdenItem          int64     `json:"nrOrdenItem" gorm:"column:NRO_ORDEN_ITEM;type:bigint"`
	CampoReservadoSSN7   string    `json:"campoReservadoSSN7" gorm:"column:CAMPO_RESERVADO_SSN7;type:varchar(50)"`
	CampoReservadoSSN8   string    `json:"campoReservadoSSN8" gorm:"column:CAMPO_RESERVADO_SSN8;type:varchar(50)"`
	Provincia            int64     `json:"provincia" gorm:"column:PROVINCIA;type:bigint;not null"`
	Alcance              string    `json:"alcance" gorm:"column:ALCANCE;type:varchar(2)"`
	JurisdiccionNacional string    `json:"jurisdiccionNacional" gorm:"column:JURISDICCION_NACIONAL;type:varchar(2)"`
	TipoServicio         string    `json:"tipoServicio" gorm:"column:TIPO_SERVICIO;type:varchar(2)"`
	UsoDestino           string    `json:"usoDestino" gorm:"column:USO_DESTINO;type:varchar(2);not null"`
	PrimaTarifa          float64   `json:"primaTarifa" gorm:"column:PRIMA_TARIFA;type:numeric(21,2);not null"`
	Premio               float64   `json:"premio" gorm:"column:PREMIO;type:numeric(21,2);not null"`
	FechaHoraCreacionSSN time.Time `json:"fechaHoraCreacionSSN" gorm:"column:FECHA_HORA_CREACION_SSN;type:timestamp;not null"`
}
