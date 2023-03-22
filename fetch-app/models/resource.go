package models

import "time"

type Resources struct {
	AreaKota     string    `json:"area_kota"`
	AreaProvinsi string    `json:"area_provinsi"`
	Komoditas    string    `json:"komoditas"`
	Price        string    `json:"price"`
	Usd          float64   `json:"usd"`
	Size         string    `json:"size"`
	TglParsed    time.Time `json:"tgl_parsed"`
	Timestamp    string    `json:"timestamp"`
	Uuid         string    `json:"uuid"`
}
