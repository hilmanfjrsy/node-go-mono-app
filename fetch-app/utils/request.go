package utils

import (
	"fetch-app/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

func GetExchange() ([]byte, error) {
	reply, err := redis.Bytes(config.Redis.Do("GET", "USDtoIDR"))
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
	_, err = config.Redis.Do("SET", "USDtoIDR", string(res), "EX", time.Hour.Abs().Seconds())
	if err != nil {
		log.Default().Println("Err set USDtoIDR redis:", err.Error())
		return []byte{}, err
	}
	return res, err
}

func GetResource() ([]byte, error) {
	reply, err := redis.Bytes(config.Redis.Do("GET", "resource-list"))
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
	_, err = config.Redis.Do("SET", "resource-list", string(res), "EX", time.Hour.Abs().Seconds())
	if err != nil {
		log.Default().Println("Err set resource-list redis:", err.Error())
		return []byte{}, err
	}
	return res, err
}
