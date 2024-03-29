package handlers

import (
	"demo/db"
	"demo/dtos"
	"demo/models"
	"demo/validations"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func AltaPolizaVehiculo(c *gin.Context) {
	var vehiculoPolizaDTO dtos.VehiculoPolizaDTO

	if err := c.ShouldBindJSON(&vehiculoPolizaDTO); err != nil {
		validations.ValidarBindingDelJSON(c, err)
		return
	}

	tx := db.GetGormDB().Begin()

	poliza, err := altaDePoliza(vehiculoPolizaDTO, tx, c)
	if err != nil {
		handleError(tx, c, err)
		return
	}

	if err := altaDeDomicilios(vehiculoPolizaDTO, poliza, tx, c); err != nil {
		handleError(tx, c, err)
		return
	}

	if err := altaDeVehiculos(vehiculoPolizaDTO, poliza, tx, c); err != nil {
		handleError(tx, c, err)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Poliza insertada correctamente"})
}

func handleError(tx *gorm.DB, c *gin.Context, err error) {
	log.Printf("Error: %v", err)
	tx.Rollback()
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar la Poliza"})
}
func convertToDate(fecha string) time.Time {
	date, err := time.Parse("02012006", fecha)
	if err != nil {
		return time.Time{}
	}

	return date
}
func altaDePoliza(vehiculoPolizaDTO dtos.VehiculoPolizaDTO, tx *gorm.DB, c *gin.Context) (models.Poliza, error) {
	nroExpediente := generarNroExpediente(vehiculoPolizaDTO.CodigoCompania)

	strCodigoRamo := strconv.FormatInt(vehiculoPolizaDTO.CodigoRamo, 10)
	strNroEndoso := strconv.FormatInt(vehiculoPolizaDTO.NroEndoso, 10)

	fechaEmisionDate := convertToDate(vehiculoPolizaDTO.FechaEmision)

	poliza := models.Poliza{
		NroExpediente:             nroExpediente,
		NroPoliza:                 vehiculoPolizaDTO.NroPoliza,
		IDPoliza:                  vehiculoPolizaDTO.CodigoCompania + "-" + strCodigoRamo + "-" + vehiculoPolizaDTO.NroPoliza + "-" + strNroEndoso,
		CodigoCompania:            vehiculoPolizaDTO.CodigoCompania,
		Cuit:                      vehiculoPolizaDTO.Cuit,
		CodigoRamo:                vehiculoPolizaDTO.CodigoRamo,
		CodigoSubRamo:             vehiculoPolizaDTO.CodigoSubRamo,
		CodigoSeguimiento:         vehiculoPolizaDTO.CodigoSeguimiento,
		CampoReservadoSSN1:        vehiculoPolizaDTO.CampoReservadoSSN1,
		CampoReservadoSSN2:        vehiculoPolizaDTO.CampoReservadoSSN2,
		NroOrdenEndoso:            vehiculoPolizaDTO.NroOrdenEndoso,
		NroEndoso:                 vehiculoPolizaDTO.NroEndoso,
		TipoEndoso:                vehiculoPolizaDTO.TipoEndoso,
		Coaseguro:                 vehiculoPolizaDTO.Coaseguro,
		Piloto:                    vehiculoPolizaDTO.Piloto,
		CompaniaPiloto:            vehiculoPolizaDTO.CompaniaPiloto,
		NroPolizaPiloto:           vehiculoPolizaDTO.NroPolizaPiloto,
		Participacion:             vehiculoPolizaDTO.Participacion,
		NroPolizaRenovacion:       vehiculoPolizaDTO.NroPolizaRenovacion,
		FechaEmision:              fechaEmisionDate,
		FechaDesde:                time.Time{},
		FechaHasta:                time.Time{},
		TomadorAsegurado:          vehiculoPolizaDTO.TomadorAsegurado,
		RazonSocialNomApell:       vehiculoPolizaDTO.RazonSocialNomApell,
		TipoDocumento:             vehiculoPolizaDTO.TipoDocumento,
		NroDocumento:              vehiculoPolizaDTO.NroDocumento,
		ProvinciaTomador:          vehiculoPolizaDTO.ProvinciaTomador,
		TipoActoAdministrativo:    vehiculoPolizaDTO.TipoActoAdministrativo,
		EsDirecto:                 vehiculoPolizaDTO.EsDirecto,
		EsFlota:                   vehiculoPolizaDTO.EsFlota,
		NroActoAdministrativo:     vehiculoPolizaDTO.NroActoAdministrativo,
		CantidadRiesgo:            vehiculoPolizaDTO.CantidadRiesgo,
		TipoMoneda:                vehiculoPolizaDTO.TipoMoneda,
		MatriculaProdAgentIntTipo: vehiculoPolizaDTO.MatriculaProdAgentIntTipo,
		MatriculaProdAgentInt:     vehiculoPolizaDTO.MatriculaProdAgentInt,
		OrganizadorTipo:           vehiculoPolizaDTO.OrganizadorTipo,
		Organizador:               vehiculoPolizaDTO.Organizador,
		PrimaPura:                 vehiculoPolizaDTO.PrimaPura,
		GastosProduccion:          vehiculoPolizaDTO.GastosProduccion,
		GastosExplotacion:         vehiculoPolizaDTO.GastosExplotacion,
		PrimaTarifa:               vehiculoPolizaDTO.PrimaTarifa,
		RecargoFinanciero:         vehiculoPolizaDTO.RecargoFinanciero,
		IVA:                       vehiculoPolizaDTO.IVA,
		IngresosBrutos:            vehiculoPolizaDTO.IngresosBrutos,
		Sellados:                  vehiculoPolizaDTO.Sellados,
		TasaSSN:                   vehiculoPolizaDTO.TasaSSN,
		CuotaSocial:               vehiculoPolizaDTO.CuotaSocial,
		OtrosImpuestos:            vehiculoPolizaDTO.OtrosImpuestos,
		Bonificacion:              vehiculoPolizaDTO.Bonificacion,
		Premio:                    vehiculoPolizaDTO.Premio,
		FormaPago:                 vehiculoPolizaDTO.FormaPago,
		CantidadDias:              vehiculoPolizaDTO.CantidadDias,
		CampoReservadoSSN3:        vehiculoPolizaDTO.CampoReservadoSSN3,
		ProvinciaAsegurado:        vehiculoPolizaDTO.ProvinciaAsegurado,
		SumaAsegurada:             vehiculoPolizaDTO.SumaAsegurada,
		TipoProducto:              vehiculoPolizaDTO.TipoProducto,
	}

	if err := tx.Create(&poliza).Error; err != nil {
		return models.Poliza{}, err
	}

	return poliza, nil
}

func altaDeVehiculos(vehiculoPolizaDTO dtos.VehiculoPolizaDTO, poliza models.Poliza, tx *gorm.DB, c *gin.Context) error {
	var vehiculos []models.VehiculosModel

	var coberturasList [][]dtos.CoberturaDTO
	for _, vehiculoDTO := range vehiculoPolizaDTO.VehiculoDTO {
		coberturasList = append(coberturasList, vehiculoDTO.CoberturaDTO)
	}

	for id, vehiculoDTO := range vehiculoPolizaDTO.VehiculoDTO {
		vehiculo := models.VehiculosModel{
			IDPoliza:             poliza.IDPoliza,
			IDVehiculo:           id,
			TipoVehiculo:         vehiculoDTO.TipoVehiculo,
			FechaHoraCreacionSSN: time.Now(),
		}
		vehiculos = append(vehiculos, vehiculo)

		if err := altaDeCoberturas(coberturasList[id], poliza, vehiculo, id, tx, c); err != nil {
			handleError(tx, c, err)
			return err
		}
	}

	if err := tx.Create(&vehiculos).Error; err != nil {
		return err
	}

	return nil
}

func altaDeCoberturas(coberturasDTO []dtos.CoberturaDTO, poliza models.Poliza, vehiculo models.VehiculosModel, idAuto int, tx *gorm.DB, c *gin.Context) error {
	var coberturas []models.Cobertura

	for id, coberturaDTO := range coberturasDTO {
		cobertura := models.Cobertura{
			IDVehiculo:              idAuto,
			IDPoliza:                poliza.IDPoliza,
			IDCobertura:             id,
			TipoCobertura:           coberturaDTO.TipoCobertura,
			CoberturaFranquicia:     coberturaDTO.CoberturaFranquicia,
			MontoFranquicia:         coberturaDTO.MontoFranquicia,
			LimiteMaxAcontecimiento: coberturaDTO.LimiteMaxAcontecimiento,
		}
		coberturas = append(coberturas, cobertura)
	}

	if err := tx.Create(&coberturas).Error; err != nil {
		return err
	}

	return nil
}

func altaDeDomicilios(vehiculoPolizaDTO dtos.VehiculoPolizaDTO, poliza models.Poliza, tx *gorm.DB, c *gin.Context) error {
	var domicilios []models.Domicilio
	for _, domicilioDTO := range vehiculoPolizaDTO.DomicilioDTO {
		domicilio := models.Domicilio{
			IDPoliza:             poliza.IDPoliza,
			CalleRuta:            domicilioDTO.CalleRuta,
			FechaHoraCreacionSSN: time.Now(),
		}
		domicilios = append(domicilios, domicilio)
	}

	if err := tx.Create(&domicilios).Error; err != nil {
		return err
	}

	return nil
}

func generarNroExpediente(codigoCia string) string {
	anio := time.Now().Year()
	currentTimeInNano := time.Now().UnixNano()
	hexValue := fmt.Sprintf("%x", currentTimeInNano)
	upperHexValue := strings.ToUpper(hexValue)

	return fmt.Sprintf("WS-%v-%s-%s-EMI", anio, upperHexValue, codigoCia)
}
