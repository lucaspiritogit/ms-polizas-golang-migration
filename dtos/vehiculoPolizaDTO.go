package dtos

type VehiculoPolizaDTO struct {
	DomicilioDTO []DomicilioDTO `json:"domicilio"`
	VehiculoDTO  []VehiculoDTO  `json:"vehiculos"`
	PolizaDTO
}
