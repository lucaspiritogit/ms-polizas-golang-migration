package models

import "time"

func (Poliza) TableName() string {
	return "SSNT_GED.NEGO_POLIZAS"
}

type Poliza struct {
	IDPoliza                  string    `gorm:"column:ID_POLIZA;type:varchar(50);not null"`
	NroExpediente             string    `json:"nroExpediente" gorm:"column:NRO_EXPEDIENTE;type:varchar(50);not null"`
	CodigoCompania            string    `json:"codigoCompania" gorm:"column:CODIGO_COMPANIA;type:varchar(4);not null;primaryKey"`
	Cuit                      int64     `json:"cuit" gorm:"column:CUIT;type:bigint;not null"`
	CodigoRamo                int64     `json:"codigoRamo" gorm:"column:CODIGO_RAMO;type:bigint;not null;primaryKey"`
	CodigoSubRamo             int64     `json:"codigoSubRamo" gorm:"column:CODIGO_SUBRAMO;type:bigint;not null"`
	CodigoSeguimiento         string    `json:"codigoSeguimiento" gorm:"column:CODIGO_SEGUIMIENTO;type:varchar(50)"`
	CampoReservadoSSN1        string    `json:"campoReservadoSSN1" gorm:"column:CAMPO_RESERVADO_SSN1;type:varchar(100)"`
	CampoReservadoSSN2        string    `json:"campoReservadoSSN2" gorm:"column:CAMPO_RESERVADO_SSN2;type:varchar(50)"`
	NroOrdenEndoso            int64     `json:"nroOrdenEndoso" gorm:"column:NRO_ORDEN_ENDOSO;type:bigint"`
	NroPoliza                 string    `json:"nroPoliza" gorm:"column:NRO_POLIZA;type:varchar(24);not null;primaryKey"`
	NroEndoso                 int64     `json:"nroEndoso" gorm:"column:NRO_ENDOSO;type:bigint;not null;primaryKey"`
	TipoEndoso                int64     `json:"tipoEndoso" gorm:"column:TIPO_ENDOSO;type:bigint;not null"`
	Coaseguro                 string    `json:"coaseguro" gorm:"column:COASEGURO;type:varchar(2);not null"`
	Piloto                    string    `json:"piloto" gorm:"column:PILOTO;type:varchar(2)"`
	CompaniaPiloto            string    `json:"companiaPiloto" gorm:"column:COMPANIA_PILOTO;type:varchar(4)"`
	NroPolizaPiloto           string    `json:"nroPolizaPiloto" gorm:"column:NRO_POLIZA_PILOTO;type:varchar(24)"`
	Participacion             int64     `json:"participacion" gorm:"column:PARTICIPACION;type:bigint"`
	NroPolizaRenovacion       string    `json:"nroPolizaRenovacion" gorm:"column:NRO_POLIZA_RENOVACION;type:varchar(24)"`
	FechaEmision              time.Time `gorm:"column:FECHA_EMISION;not null"`
	FechaDesde                time.Time `gorm:"column:FECHA_VIGENCIA_DESDE;not null"`
	FechaHasta                time.Time `gorm:"column:FECHA_VIGENCIA_HASTA"`
	TomadorAsegurado          string    `gorm:"column:TOMADOR_ASEGURADO;type:varchar(1)"`
	RazonSocialNomApell       string    `gorm:"column:NOMBREAPELLIDO_RAZONSOCIAL;not null;type:varchar(50)"`
	TipoDocumento             string    `gorm:"column:TIPO_DOCUMENTO;not null;type:varchar(4)"`
	NroDocumento              string    `gorm:"column:NRO_DOCUMENTO;not null;type:varchar(11)"`
	ProvinciaTomador          int64     `gorm:"column:PROVINCIA_TOMADOR;not null"`
	TipoActoAdministrativo    string    `gorm:"column:TIPO_ACTO_ADMINISTRATIVO;not null;type:varchar(50)"`
	NroActoAdministrativo     string    `gorm:"column:NRO_ACTO_ADMINISTRATIVO;not null;type:varchar(24)"`
	EsFlota                   string    `gorm:"column:ES_FLOTA;type:varchar(2)"`
	CantidadRiesgo            int64     `gorm:"column:CANTIDAD_RIESGO;type:bigint"`
	TipoMoneda                int64     `gorm:"column:TIPO_MONEDA;not null;type:bigint"`
	EsDirecto                 string    `gorm:"column:ES_DIRECTO;not null;type:varchar(2)"`
	MatriculaProdAgentIntTipo string    `gorm:"column:MATPRODUCTOR_AGENTINSTITORIO_TIPO;type:varchar(2)"`
	MatriculaProdAgentInt     int64     `gorm:"column:MATPRODUCTOR_AGENTINSTITORIO;type:bigint"`
	OrganizadorTipo           string    `gorm:"column:ORGANIZADOR_TIPO;type:varchar(2)"`
	Organizador               int64     `gorm:"column:ORGANIZADOR;type:bigint"`
	PrimaPura                 float64   `gorm:"column:PRIMA_PURA;not null;type:numeric(21,2)"`
	GastosProduccion          float64   `gorm:"column:GASTOS_PRODUCCION;not null;type:numeric(21,2)"`
	GastosExplotacion         float64   `gorm:"column:GASTOS_EXPLOTACION;not null;type:numeric(21,2)"`
	PrimaTarifa               float64   `gorm:"column:PRIMA_TARIFA;not null;type:numeric(21,2)"`
	RecargoFinanciero         float64   `gorm:"column:RECARGO_FINANCIERO;not null;type:numeric(21,2)"`
	IVA                       float64   `gorm:"column:IVA;not null;type:numeric(21,2)"`
	IngresosBrutos            float64   `gorm:"column:INGRESOS_BRUTOS;not null;type:numeric(21,2)"`
	Sellados                  float64   `gorm:"column:SELLADOS;not null;type:numeric(21,2)"`
	TasaSSN                   float64   `gorm:"column:TASA_SSN;not null;type:numeric(21,2)"`
	CuotaSocial               float64   `gorm:"column:CUOTA_SOCIAL;not null;type:numeric(21,2)"`
	OtrosImpuestos            float64   `gorm:"column:OTROS_IMPUESTOS;not null;type:numeric(21,2)"`
	Bonificacion              *float64  `gorm:"column:BONIFICACION;type:numeric(21,2)"`
	Premio                    float64   `gorm:"column:PREMIO;not null;type:numeric(21,2)"`
	FormaPago                 string    `gorm:"column:FORMA_PAGO;not null;type:varchar(2)"`
	CantidadDias              int64     `gorm:"column:CANTIDAD_DIAS;type:bigint"`
	CampoReservadoSSN3        string    `gorm:"column:CAMPO_RESERVADO_SSN3;type:varchar(50)"`
	ProvinciaAsegurado        int64     `gorm:"column:PROVINCIA_ASEGURADO;type:bigint"`
	SumaAsegurada             *float64  `gorm:"column:SUMA_ASEGURADA;type:numeric(21,2)"`
	TipoProducto              int64     `gorm:"column:TIPO_PRODUCTO;type:bigint"`
}
