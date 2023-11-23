package dtos

type VehiculoDTO struct {
	CoberturaDTO         []CoberturaDTO `json:"coberturas"`
	TipoVehiculo         string         `json:"tipoVehiculo" binding:"required"`
	Patente              string         `json:"patente" binding:"required"`
	UbicacionRiesgo      string         `json:"ubicacionRiesgo" binding:"required"`
	Marca                string         `json:"marca" binding:"required"`
	Modelo               string         `json:"modelo" binding:"required"`
	AnioFabricacion      int64          `json:"anioFabricacion" binding:"required"`
	NroMotor             string         `json:"nroMotor"`
	NroChasis            string         `json:"nroChasis"`
	CodSeguimiento       string         `json:"codSeguimiento"`
	Suma                 float64        `json:"suma" binding:"required"`
	TipoMoneda           int64          `json:"tipoMoneda" binding:"required"`
	Prenda               string         `json:"prenda" binding:"required"`
	EstadoVehiculoCod    string         `json:"estadoVehiculoCod"`
	NrOrdenItem          int64          `json:"nrOrdenItem"`
	CampoReservadoSSN7   string         `json:"campoReservadoSSN7"`
	CampoReservadoSSN8   string         `json:"campoReservadoSSN8"`
	Provincia            int64          `json:"provincia" binding:"required"`
	Alcance              string         `json:"alcance"`
	JurisdiccionNacional string         `json:"jurisdiccionNacional"`
	TipoServicio         string         `json:"tipoServicio"`
	UsoDestino           string         `json:"usoDestino" binding:"required"`
	PrimaTarifa          float64        `json:"primaTarifa" binding:"required"`
	Premio               float64        `json:"premio" binding:"required"`
}
