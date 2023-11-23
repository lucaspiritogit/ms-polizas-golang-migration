package models

func (Poliza) TableName() string {
	return "SSNT_GED.NEGO_POLIZAS"
}

type Poliza struct {
	IDPoliza            string `gorm:"column:ID_POLIZA;type:varchar(50);not null"`
	NroExpediente       string `json:"nroExpediente" gorm:"column:NRO_EXPEDIENTE;type:varchar(50);not null"`
	CodigoCompania      string `json:"codigoCompania" gorm:"column:CODIGO_COMPANIA;type:varchar(4);not null;primaryKey"`
	Cuit                int64  `json:"cuit" gorm:"column:CUIT;type:bigint;not null"`
	CodigoRamo          int64  `json:"codigoRamo" gorm:"column:CODIGO_RAMO;type:bigint;not null;primaryKey"`
	CodigoSubRamo       int64  `json:"codigoSubRamo" gorm:"column:CODIGO_SUBRAMO;type:bigint;not null"`
	CodigoSeguimiento   string `json:"codigoSeguimiento" gorm:"column:CODIGO_SEGUIMIENTO;type:varchar(50)"`
	CampoReservadoSSN1  string `json:"campoReservadoSSN1" gorm:"column:CAMPO_RESERVADO_SSN1;type:varchar(100)"`
	CampoReservadoSSN2  string `json:"campoReservadoSSN2" gorm:"column:CAMPO_RESERVADO_SSN2;type:varchar(50)"`
	NroOrdenEndoso      int64  `json:"nroOrdenEndoso" gorm:"column:NRO_ORDEN_ENDOSO;type:bigint"`
	NroPoliza           string `json:"nroPoliza" gorm:"column:NRO_POLIZA;type:varchar(24);not null;primaryKey"`
	NroEndoso           int64  `json:"nroEndoso" gorm:"column:NRO_ENDOSO;type:bigint;not null;primaryKey"`
	TipoEndoso          int64  `json:"tipoEndoso" gorm:"column:TIPO_ENDOSO;type:bigint;not null"`
	Coaseguro           string `json:"coaseguro" gorm:"column:COASEGURO;type:varchar(2);not null"`
	Piloto              string `json:"piloto" gorm:"column:PILOTO;type:varchar(2)"`
	CompaniaPiloto      string `json:"companiaPiloto" gorm:"column:COMPANIA_PILOTO;type:varchar(4)"`
	NroPolizaPiloto     string `json:"nroPolizaPiloto" gorm:"column:NRO_POLIZA_PILOTO;type:varchar(24)"`
	Participacion       int64  `json:"participacion" gorm:"column:PARTICIPACION;type:bigint"`
	NroPolizaRenovacion string `json:"nroPolizaRenovacion" gorm:"column:NRO_POLIZA_RENOVACION;type:varchar(24)"`
}
