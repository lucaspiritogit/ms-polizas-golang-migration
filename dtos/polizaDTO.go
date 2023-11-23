package dtos

type PolizaDTO struct {
	CodigoCompania      string `json:"codigoCompania" binding:"required"`
	Cuit                int64  `json:"cuit" binding:"required"`
	CodigoRamo          int64  `json:"codigoRamo" binding:"required"`
	CodigoSubRamo       int64  `json:"codigoSubRamo" binding:"required"`
	CodigoSeguimiento   string `json:"codigoSeguimiento"`
	CampoReservadoSSN1  string `json:"campoReservadoSSN1"`
	CampoReservadoSSN2  string `json:"campoReservadoSSN2"`
	NroOrdenEndoso      int64  `json:"nroOrdenEndoso"`
	NroPoliza           string `json:"nroPoliza" binding:"required"`
	NroEndoso           int64  `json:"nroEndoso"`
	TipoEndoso          int64  `json:"tipoEndoso" binding:"required"`
	Coaseguro           string `json:"coaseguro" binding:"required"`
	Piloto              string `json:"piloto"`
	CompaniaPiloto      string `json:"companiaPiloto"`
	NroPolizaPiloto     string `json:"nroPolizaPiloto"`
	Participacion       int64  `json:"participacion"`
	NroPolizaRenovacion string `json:"nroPolizaRenovacion"`
}