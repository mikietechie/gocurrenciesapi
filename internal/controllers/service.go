/*
Date Created		1 May 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			A Fintech Data Service
Inspired by			https://freecurrencyapi.com
*/

package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikietechie/gocurrenciesapi/internal/middleware"
	"github.com/mikietechie/gocurrenciesapi/internal/responses"
	"github.com/mikietechie/gocurrenciesapi/internal/services"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
)

// Account			godoc
// @Summary      	Account
// @Description  	Gets Account
// @Tags         	Service
// @Produce      	json
// @Success      	200   {object}  models.Client
// @Router       	/api/v1/service/account [get]
// @Security 		ApiKeyAuth
func GetClientInfo(c *gin.Context) {
	data := middleware.GetClientFromC(c)
	responses.JSON200(c, data)
}

// Currencies State	godoc
// @Summary      	Currencies State
// @Description  	Gets Currencies State
// @Tags         	Service
// @Produce      	json
// @Success      	200   {object}  structs.BeaconResponse
// @Router       	/api/v1/service/live [get]
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
// @Tags         Service
// @Produce      json
// @Success      200   {object}  []string
// @Router       /api/v1/service/currencies [get]
// @Security 	ApiKeyAuth
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
// @Tags         Service
// @Produce      json
// @Param        toCurrency path string true "To Currency"
// @Param        fromCurrency path string true "From Currency"
// @Param        amount path float64 true "Amount"
// @Success      200   {object}  float64
// @Router       /api/v1/service/convert/{toCurrency}/{fromCurrency}/{amount} [get]
// @Security 	ApiKeyAuth
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

// Rate at Datetime Endpoint        godoc
// @Summary      Rate at Datetime Endpoint
// @Description  Gets Rate at Datetime Endpoint
// @Tags         Service
// @Produce      json
// @Param        currency    query    string  true  "Currency Code"
// @Param        timestamp   query    string  true  "Time Stamp"
// @Success      200   {object}  models.Rate
// @Router       /api/v1/service/prevailing [get]
// @Security 	ApiKeyAuth
func GetRateAt(c *gin.Context) {
	var body structs.RateAtDateBody
	c.BindQuery(&body)
	rate, err := services.GetRateAt(body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, rate)
}

// Rate in Period Endpoint        godoc
// @Summary      Rate in Period Endpoint
// @Description  Gets Rate in Period Endpoint
// @Tags         Service
// @Produce      json
// @Param        currencies  query    []string  true  "Currency Code"
// @Param        start       query    string    true  "Start"
// @Param        end         query    string    true  "Start"
// @Success      200   {object}  []models.Rate
// @Router       /api/v1/service/historical [get]
// @Security 	ApiKeyAuth
func GetRatesInPeriod(c *gin.Context) {
	var body structs.RatesInPeriodBody
	err := c.BindQuery(&body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	data, err := services.GetRatesBetween(body)
	if err != nil {
		responses.JSON400(c, err.Error())
		return
	}
	responses.JSON200(c, data)
}

func ServiceRouter(r gin.RouterGroup) {
	r.Use(middleware.WithClient())
	r.GET("/live", GetExchangeRates)
	r.GET("/currencies", GetCurrencies)
	r.GET("/convert/:toCurrency/:fromCurrency/:amount", GetConversion)
	r.GET("/prevailing", GetRateAt)
	r.GET("/historical", GetRatesInPeriod)
}
