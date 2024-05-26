package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/mikietechie/gocurrenciesapi/internal/cache"
	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"github.com/mikietechie/gocurrenciesapi/internal/drivers"
	"github.com/mikietechie/gocurrenciesapi/internal/models"
	"github.com/mikietechie/gocurrenciesapi/internal/structs"
)

const RATES_KEY = "ExchangeRates"

func FetchExchangeRates() (structs.BeaconResponse, error) {
	log.Println("Fetching Exchange Rates")
	var obj structs.BeaconResponse
	var body []byte
	url := fmt.Sprintf("%s/latest/?base=%s&api_key=%s", config.BEACON_URL, config.BEACON_BASE_CURRENCY, config.BEACON_KEY)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return obj, err
	}
	if res.StatusCode != 200 {
		err = errors.New("server returned something else")
		log.Println(err.Error())
		return obj, err
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return obj, err
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		log.Println(err.Error())
		return obj, err
	}
	_, err = cache.RDB.Set(config.CTX, RATES_KEY, string(body), time.Minute*15).Result()
	if err != nil {
		log.Println("failed store in cache\n", err)
	}
	go StoreRates(GetRates(obj))
	log.Println("Fetched Exchange Rates")
	return obj, nil
}

func GetCachedExchangeRates() (structs.BeaconResponse, error) {
	var obj structs.BeaconResponse
	data, err := cache.RDB.Get(config.CTX, RATES_KEY).Result()
	if data != "" && err == nil {
		log.Println("Fetched Exchange Rates from cache")
		json.Unmarshal([]byte(data), &obj)
		// log.Println("data\n", data)
	}
	return obj, err
}

func GetExchangeRates() (structs.BeaconResponse, error) {
	obj, err := GetCachedExchangeRates()
	if err == nil {
		return obj, err
	}
	obj, err = FetchExchangeRates()
	if err != nil {
		return obj, err
	}
	return obj, nil
}

func GetCurrencies() ([]string, error) {
	obj, err := GetExchangeRates()
	if err != nil {
		return nil, err
	}
	var currencies []string
	for k := range obj.Data.Rates {
		currencies = append(currencies, k)
	}
	return currencies, nil
}

func GetConversion(toCurrency, fromCurrency string, amount float64) (float64, error) {
	obj, err := GetExchangeRates()
	if err != nil {
		return 0, err
	}
	var toRate, fromRate float64 = 1, 1
	if toCurrency != config.BEACON_BASE_CURRENCY {
		toRate = obj.Data.Rates[toCurrency]
		if toRate == 0 {
			return 0, nil
		}
	}
	if fromCurrency != config.BEACON_BASE_CURRENCY {
		fromRate = obj.Data.Rates[fromCurrency]
		if fromRate == 0 {
			return 0, nil
		}
	}
	return amount * toRate / fromRate, nil
}

func GetRates(obj structs.BeaconResponse) []interface{} {
	var rates []interface{}
	for currency, value := range obj.Data.Rates {
		rates = append(rates, models.Rate{
			Currency:  currency,
			Value:     value,
			Timestamp: obj.Data.Date,
		})
	}
	return rates
}

func StoreRates(rates []interface{}) error {
	_, err := drivers.Mongod.Collection("rates").InsertMany(
		config.CTX,
		rates,
	)
	if err != nil {
		log.Println("Error Store Rates\n", err)
		return err
	}
	return nil
}
