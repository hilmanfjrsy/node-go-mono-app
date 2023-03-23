package utils

import (
	"fetch-app/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func GetExchange() ([]byte, error) {
	reply, err := config.GetCache("USDtoIDR")
	if err == nil {
		fmt.Println("Exchange from cache")
		return reply, nil
	}
	req, err := http.NewRequest("GET", "https://api.apilayer.com/exchangerates_data/latest?symbols=USD&base=IDR", nil)
	if err != nil {
		log.Default().Println("Err new request exchange:", err.Error())
		return []byte{}, err
	}

	req.Header.Set("apikey", os.Getenv("API_KEY_ENCHAGE_RATES"))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Default().Println("Err client do exchange:", err.Error())
		return []byte{}, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	err = config.SetCache("USDtoIDR", res, time.Hour)
	if err != nil {
		log.Default().Println("Err set USDtoIDR cache:", err.Error())
		return []byte{}, err
	}
	return res, err
}

func GetResource() ([]byte, error) {
	reply, err := config.GetCache("resource-list")
	if err == nil {
		fmt.Println("Resource from cache")
		return reply, nil
	}
	req, err := http.NewRequest("GET", "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list", nil)
	if err != nil {
		log.Default().Println("Err new request resource:", err.Error())
		return []byte{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Default().Println("Err client do resource:", err.Error())
		return []byte{}, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	err = config.SetCache("resource-list", res, time.Hour)
	if err != nil {
		log.Default().Println("Err set resource-list cache:", err.Error())
		return []byte{}, err
	}
	return res, err
}
