package utils

import (
	"os"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

const resourceString = `{
	"data":[{
		"area_kota": "LAMPUNG TIMUR",
		"area_provinsi": "LAMPUNG",
		"komoditas": "BANDENG MALANG",
		"price": "99800",
		"size": "200",
		"tgl_parsed": "2023-03-20T21:19:41+07:00",
		"timestamp": "1679321981",
		"uuid": "91b685ef-ff94-4645-a0d9-25c92ffab34f"
	}]
}`

const currencyString = `{
	"success": true,
	"timestamp": 1679575083,
	"base": "IDR",
	"date": "2023-03-23",
	"rates": {
			"USD": 6.6189222e-05
	}
}`

func TestRequestResource(t *testing.T) {
	defer os.RemoveAll("cache")
	defer gock.Off()
	gock.New("https://stein.efishery.com").
		Get("/v1/storages/5e1edf521073e315924ceab4/list").
		Reply(200).
		JSON(resourceString)

	b, err := GetResource()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.JSONEq(t, resourceString, string(b))
}

func TestRequestCurrency(t *testing.T) {
	defer os.RemoveAll("cache")
	defer gock.Off()
	gock.New("https://api.apilayer.com").
		Get("/exchangerates_data/latest").
		MatchParam("symbols", "USD").
		MatchParam("base", "IDR").
		MatchHeader("apikey", "").
		MatchHeader("Accept", "application/json").
		Reply(200).
		JSON(currencyString)

	b, err := GetExchange()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.JSONEq(t, currencyString, string(b))
}
