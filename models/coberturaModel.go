package models

func (Cobertura) TableName() string {
	return "SSNT_GED.NEGO_COBERTURAS"
}

type Cobertura struct {
	IDVehiculo              int      `gorm:"column:ID_VEHICULO;type:integer;not null"`
	IDPoliza                string   `gorm:"column:ID_POLIZA;type:varchar(50);not null"`
	IDCobertura             int      `gorm:"column:ID_COBERTURA;type:integer;not null"`
	TipoCobertura           int64    `gorm:"column:TIPO_COBERTURA;not null"`
	CoberturaFranquicia     string   `gorm:"column:COBERTURA_FRANQUICIA;not null"`
	MontoFranquicia         *float64 `gorm:"column:MONTO_FRANQUICIA"`
	LimiteMaxAcontecimiento float64  `gorm:"column:LIMITE_MAX_ACONTECIMIENTO;not null"`
}
