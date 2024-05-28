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
	"go.mongodb.org/mongo-driver/bson"
)

const RATES_KEY = "ExchangeRates"

func FetchExchangeRates() (structs.BeaconResponse, error) {
	log.Println("Fetching Exchange Rates")
	var beacon structs.BeaconResponse
	var body []byte
	url := fmt.Sprintf("%s/latest/?base=%s&api_key=%s", config.BEACON_URL, config.BEACON_BASE_CURRENCY, config.BEACON_KEY)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return beacon, err
	}
	if res.StatusCode != 200 {
		err = errors.New("server returned something else")
		log.Println(err.Error())
		return beacon, err
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return beacon, err
	}
	err = json.Unmarshal(body, &beacon)
	if err != nil {
		log.Println(err.Error())
		return beacon, err
	}
	_, err = cache.RDB.Set(config.CTX, RATES_KEY, string(body), time.Minute*time.Duration(config.RATES_LIFETIME)).Result()
	if err != nil {
		log.Println("failed store in cache\n", err)
	}
	go StoreRatesToDB(GetRatesFromBeacon(beacon))
	log.Println("Fetched Exchange Rates")
	return beacon, nil
}

func GetCachedExchangeRates() (structs.BeaconResponse, error) {
	var beacon structs.BeaconResponse
	data, err := cache.RDB.Get(config.CTX, RATES_KEY).Result()
	if data != "" && err == nil {
		log.Println("Fetched Exchange Rates from cache")
		json.Unmarshal([]byte(data), &beacon)
		// log.Println("data\n", data)
	}
	return beacon, err
}

func GetExchangeRates() (structs.BeaconResponse, error) {
	beacon, err := GetCachedExchangeRates()
	if err == nil {
		return beacon, err
	}
	beacon, err = FetchExchangeRates()
	if err != nil {
		return beacon, err
	}
	return beacon, nil
}

func GetCurrencies() ([]string, error) {
	beacon, err := GetExchangeRates()
	if err != nil {
		return nil, err
	}
	var currencies []string
	for k := range beacon.Data.Rates {
		currencies = append(currencies, k)
	}
	return currencies, nil
}

func GetConversion(toCurrency, fromCurrency string, amount float64) (float64, error) {
	beacon, err := GetExchangeRates()
	if err != nil {
		return 0, err
	}
	var toRate, fromRate float64 = 1, 1
	if toCurrency != config.BEACON_BASE_CURRENCY {
		toRate = beacon.Data.Rates[toCurrency]
		if toRate == 0 {
			return 0, nil
		}
	}
	if fromCurrency != config.BEACON_BASE_CURRENCY {
		fromRate = beacon.Data.Rates[fromCurrency]
		if fromRate == 0 {
			return 0, nil
		}
	}
	return amount * toRate / fromRate, nil
}

func GetRatesFromBeacon(beacon structs.BeaconResponse) []interface{} {
	var rates []interface{}
	for currency, value := range beacon.Data.Rates {
		rates = append(rates, models.Rate{
			Currency:  currency,
			Value:     value,
			Timestamp: beacon.Data.Date,
		})
	}
	return rates
}

func StoreRatesToDB(rates []interface{}) error {
	_, err := drivers.Mongod.Collection("rates").InsertMany(
		config.CTX,
		rates,
	)
	if err != nil {
		log.Println("Error Store Rates\n", err)
		return err
	}
	log.Println("Stored rates", len(rates))
	return nil
}

func GetRateAt(params structs.RateAtDateBody) (models.Rate, error) {
	var rate models.Rate
	filter := bson.M{
		"timestamp": bson.M{
			"$gte": params.Timestamp,
			"$lte": params.Timestamp.Add(time.Minute * time.Duration(config.RATES_LIFETIME)),
		},
		"currency": params.Currency,
	}
	log.Println(filter)
	result := drivers.Mongod.Collection("rates").FindOne(config.CTX, filter)
	if result.Err() != nil {
		return rate, result.Err()
	}
	err := result.Decode(&rate)
	return rate, err
}

func GetRatesBetween(params structs.RatesInPeriodBody) ([]models.Rate, error) {
	var rates []models.Rate
	filter := bson.M{
		"timestamp": bson.M{
			"$gte": params.Start,
			"$lte": params.End,
		},
		"currency": bson.M{"$in": params.Currencies},
	}
	cursor, err := drivers.Mongod.Collection("rates").Find(config.CTX, filter)
	if err != nil {
		return rates, err
	}
	err = cursor.All(config.CTX, &rates)
	return rates, err
}
