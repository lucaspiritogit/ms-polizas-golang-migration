package dtos

type PolizaDTO struct {
	CodigoCompania            string   `json:"codigoCompania" binding:"required"`
	Cuit                      int64    `json:"cuit" binding:"required"`
	CodigoRamo                int64    `json:"codigoRamo" binding:"required"`
	CodigoSubRamo             int64    `json:"codigoSubRamo" binding:"required"`
	CodigoSeguimiento         string   `json:"codigoSeguimiento"`
	CampoReservadoSSN1        string   `json:"campoReservadoSSN1"`
	CampoReservadoSSN2        string   `json:"campoReservadoSSN2"`
	NroOrdenEndoso            int64    `json:"nroOrdenEndoso"`
	NroPoliza                 string   `json:"nroPoliza" binding:"required"`
	NroEndoso                 int64    `json:"nroEndoso"`
	TipoEndoso                int64    `json:"tipoEndoso" binding:"required"`
	Coaseguro                 string   `json:"coaseguro" binding:"required"`
	Piloto                    string   `json:"piloto"`
	CompaniaPiloto            string   `json:"companiaPiloto"`
	NroPolizaPiloto           string   `json:"nroPolizaPiloto"`
	Participacion             int64    `json:"participacion"`
	NroPolizaRenovacion       string   `json:"nroPolizaRenovacion"`
	FechaEmision              string   `json:"fechaEmision"`
	FechaDesde                string   `json:"fechaDesde"`
	FechaHasta                string   `json:"fechaHasta"`
	TomadorAsegurado          string   `json:"tomadorAsegurado"`
	RazonSocialNomApell       string   `json:"razonSocialNomApell"`
	TipoDocumento             string   `json:"tipoDocumento"`
	NroDocumento              string   `json:"nroDocumento"`
	ProvinciaTomador          int64    `json:"provinciaTomador"`
	TipoActoAdministrativo    string   `json:"tipoActoAdministrativo"`
	NroActoAdministrativo     string   `json:"nroActoAdministrativo"`
	EsFlota                   string   `json:"esFlota"`
	CantidadRiesgo            int64    `json:"cantidadRiesgo"`
	TipoMoneda                int64    `json:"tipoMoneda"`
	EsDirecto                 string   `json:"esDirecto"`
	MatriculaProdAgentIntTipo string   `json:"matriculaProdAgentIntTipo"`
	MatriculaProdAgentInt     int64    `json:"matriculaProdAgentInt"`
	OrganizadorTipo           string   `json:"organizadorTipo"`
	Organizador               int64    `json:"organizador"`
	PrimaPura                 float64  `json:"primaPura"`
	GastosProduccion          float64  `json:"gastosProduccion"`
	GastosExplotacion         float64  `json:"gastosExplotacion"`
	PrimaTarifa               float64  `json:"primaTarifa"`
	RecargoFinanciero         float64  `json:"recargoFinanciero"`
	IVA                       float64  `json:"iva"`
	IngresosBrutos            float64  `json:"ingresosBrutos"`
	Sellados                  float64  `json:"sellados"`
	TasaSSN                   float64  `json:"tasaSSN"`
	CuotaSocial               float64  `json:"cuotaSocial"`
	OtrosImpuestos            float64  `json:"otrosImpuestos"`
	Bonificacion              *float64 `json:"bonificacion"`
	Premio                    float64  `json:"premio"`
	FormaPago                 string   `json:"formaPago"`
	CantidadDias              int64    `json:"cantidadDias"`
	CampoReservadoSSN3        string   `json:"campoReservadoSSN3"`
	ProvinciaAsegurado        int64    `json:"provinciaAsegurado"`
	SumaAsegurada             *float64 `json:"sumaAsegurada"`
	TipoProducto              int64    `json:"tipoProducto"`
}
