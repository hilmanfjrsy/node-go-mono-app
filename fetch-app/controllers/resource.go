package controllers

import (
	"encoding/json"
	"fetch-app/config"
	"fetch-app/models"
	"fetch-app/utils"
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetResource(c *gin.Context) {
	res, err := utils.GetResource()
	if err != nil {
		log.Default().Println("Err get resource:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	var resources []models.Resources
	if err := json.Unmarshal(res, &resources); err != nil {
		log.Default().Println("Err unmarshal resource:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, resources)
}

func GetResourceWithCurrency(c *gin.Context) {
	var resources, newResources []models.Resources
	reply, err := config.GetCache("resource-currency")
	if err == nil {
		fmt.Println("Resource currency from cache")
		if err := json.Unmarshal(reply, &newResources); err != nil {
			log.Default().Println("Err unmarshal resource currency:", err.Error())
			utils.ResponseError(c, http.StatusBadRequest, err.Error())
		}
		utils.ResponseSuccess(c, http.StatusOK, newResources)
		return
	}

	res, err := utils.GetResource()
	if err != nil {
		log.Default().Println("Err get resource currency:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.Unmarshal(res, &resources); err != nil {
		log.Default().Println("Err unmarshal resource currency:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	rate, err := utils.GetExchange()
	if err != nil {
		log.Default().Println("Err get exchange currency", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	var rateCurrency models.Currency
	if err := json.Unmarshal(rate, &rateCurrency); err != nil {
		log.Default().Println("Err unmarshal exchange currency:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	for _, r := range resources {
		if !reflect.DeepEqual(r, models.Resources{}) {
			usdRate := rateCurrency.Rates["USD"].(float64)
			idr, err := strconv.ParseFloat(r.Price, 64)
			if err != nil {
				utils.ResponseError(c, http.StatusBadRequest, err.Error())
				return
			}
			ratio := math.Pow(10, float64(2))
			r.Usd = math.Round((usdRate*idr)*ratio) / ratio
			newResources = append(newResources, r)
		}
	}
	b, err := json.Marshal(newResources)
	if err != nil {
		log.Default().Println("Err marshal new resource currency:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	err = config.SetCache("resource-currency", b, time.Hour)
	if err != nil {
		log.Default().Println("Err set resource-currency cache:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, newResources)
}

func GetResourceAggregate(c *gin.Context) {
	var resources []models.Resources
	var newResources []utils.ResourceByProvinsi
	replyAgg, err := config.GetCache("resource-agg")
	if err == nil {
		fmt.Println("Resource Agg from cache")
		if err := json.Unmarshal(replyAgg, &newResources); err != nil {
			log.Default().Println("Err unmarshal resource:", err.Error())
			utils.ResponseError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.ResponseSuccess(c, http.StatusOK, newResources)
		return
	}
	reply, err := config.GetCache("resource-currency")
	if err != nil {
		res, err := utils.GetResource()
		if err != nil {
			log.Default().Println("Err get resource:", err.Error())
			utils.ResponseError(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := json.Unmarshal(res, &resources); err != nil {
			log.Default().Println("Err unmarshal resource:", err.Error())
			utils.ResponseError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := json.Unmarshal(reply, &resources); err != nil {
			log.Default().Println("Err unmarshal resource:", err.Error())
			utils.ResponseError(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	newResources = utils.Aggregates(&resources)
	b, err := json.Marshal(newResources)
	if err != nil {
		log.Default().Println("Err marshal new resource:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	err = config.SetCache("resource-agg", b, time.Hour)
	if err != nil {
		log.Default().Println("Err set resource-currency cache:", err.Error())
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, newResources)
}
