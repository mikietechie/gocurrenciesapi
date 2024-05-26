package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/middleware"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
)

// Currencies State	godoc
// @Summary      	Currencies State
// @Description  	Gets Currencies State
// @Tags         	Auth
// @Produce      	json
// @Success      	200   {object}  structs.BeaconResponse
// @Router       	/api/v1/currencies/exchange-rates [get]
// @Security 		ApiKeyAuth
func GetExchangeRates(c *gin.Context) {
	data, err := services.GetExchangeRates()
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, data)
}

// Currencies List        godoc
// @Summary      Currencies List
// @Description  Gets Currencies List
// @Tags         Auth
// @Produce      json
// @Success      200   {object}  []string
// @Router       /api/v1/currencies/list [get]
// @Security ApiKeyAuth
func GetCurrencies(c *gin.Context) {
	data, err := services.GetCurrencies()
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, data)
}

// Conversion Endpoint        godoc
// @Summary      Conversion Endpoint
// @Description  Gets Conversion Endpoint
// @Tags         Auth
// @Produce      json
// @Param        toCurrency path string true "To Currency"
// @Param        fromCurrency path string true "From Currency"
// @Param        amount path float64 true "Amount"
// @Success      200   {object}  float64
// @Router       /api/v1/currencies/conversion/{toCurrency}/{fromCurrency}/{amount} [get]
// @Security ApiKeyAuth
func GetConversion(c *gin.Context) {
	toCurrency, _ := c.Params.Get("toCurrency")
	fromCurrency, _ := c.Params.Get("fromCurrency")
	amountStr, _ := c.Params.Get("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	data, err := services.GetConversion(
		toCurrency,
		fromCurrency,
		amount,
	)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, data)
}

func CurrenciesRouter(r gin.RouterGroup) {
	r.Use(middleware.WithClient())
	r.GET("/exchange-rates", GetExchangeRates)
	r.GET("/list", GetCurrencies)
	r.GET("/conversion/:toCurrency/:fromCurrency/:amount", GetConversion)
}
