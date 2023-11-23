package dtos

type CoberturaDTO struct {
	TipoCobertura           int64    `json:"cobertura" binding:"required"`
	CoberturaFranquicia     string   `json:"coberturaFranquicia" binding:"required"`
	MontoFranquicia         *float64 `json:"montoFranquicia"`
	LimiteMaxAcontecimiento float64  `json:"limiteMaxAcontecimiento" binding:"required"`
}