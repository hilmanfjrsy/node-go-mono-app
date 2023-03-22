package utils

import (
	"fetch-app/models"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"time"
)

type ResourceByProvinsi struct {
	AreaProvinsi string           `json:"area_provinsi"`
	Weekly       []WeeklyResource `json:"weekly"`
}

type WeeklyResource struct {
	Week     string            `json:"week"`
	AllPrice []float64         `json:"all_price"`
	AllSize  []float64         `json:"all_size"`
	Price    AggregateResource `json:"price"`
	Size     AggregateResource `json:"size"`
}

type AggregateResource struct {
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
	Median float64 `json:"median"`
	Avg    float64 `json:"avg"`
}

func weekRangeDate(date time.Time) string {
	offset := (int(time.Monday) - int(date.Weekday()) - 7) % 7
	startWeek := date.Add(time.Duration(offset*24) * time.Hour)
	endWeek := startWeek.Add(time.Duration(6*24) * time.Hour)
	return fmt.Sprintf(`%s - %s`, startWeek.Format("02-01-2006"), endWeek.Format("02-01-2006"))
}

func Aggregates(resource *[]models.Resources) []ResourceByProvinsi {
	resourceByProvinsi := []ResourceByProvinsi{}
	for _, res := range *resource {
		if !reflect.DeepEqual(res, models.Resources{}) {
			idxProvinsi := findProvinceIndex(&resourceByProvinsi, res.AreaProvinsi)
			weekly := weekRangeDate(res.TglParsed)
			price, _ := strconv.ParseFloat(res.Price, 64)
			size, _ := strconv.ParseFloat(res.Size, 64)

			if idxProvinsi != -1 {
				currProvinsi := &resourceByProvinsi[idxProvinsi]
				idxWeekly := findWeeklyIndex(&currProvinsi.Weekly, weekly)

				if idxWeekly != -1 {
					currWeekly := &currProvinsi.Weekly[idxWeekly]
					currWeekly.AllPrice = append(currWeekly.AllPrice, price)
					currWeekly.AllSize = append(currWeekly.AllSize, size)
				} else {
					newWeekly := WeeklyResource{
						Week:     weekly,
						AllPrice: []float64{price},
						AllSize:  []float64{size},
					}
					currProvinsi.Weekly = append(currProvinsi.Weekly, newWeekly)
				}
			} else {
				newResourceProvinsi := ResourceByProvinsi{
					AreaProvinsi: res.AreaProvinsi,
					Weekly: []WeeklyResource{{
						Week:     weekly,
						AllPrice: []float64{price},
						AllSize:  []float64{size},
					}},
				}
				resourceByProvinsi = append(resourceByProvinsi, newResourceProvinsi)
			}
		}
	}

	for _, res := range resourceByProvinsi {
		for i, week := range res.Weekly {
			aggregates(week.AllPrice, &res.Weekly[i].Price)
			aggregates(week.AllSize, &res.Weekly[i].Size)
		}
	}
	return resourceByProvinsi
}

func findProvinceIndex(res *[]ResourceByProvinsi, areaProvinsi string) int {
	for i, r := range *res {
		if r.AreaProvinsi == areaProvinsi {
			return i
		}
	}
	return -1
}

func findWeeklyIndex(res *[]WeeklyResource, week string) int {
	for i, r := range *res {
		if r.Week == week {
			return i
		}
	}
	return -1
}

func aggregates(dt []float64, agg *AggregateResource) {
	l := len(dt)
	agg.Min = math.NaN()
	agg.Max = math.NaN()
	agg.Avg = 0
	sort.Float64s(dt)
	for _, d := range dt {
		if d < agg.Min || (math.IsNaN(agg.Min) && !math.IsNaN(d)) {
			agg.Min = d
		}
		if d < agg.Max || (math.IsNaN(agg.Max) && !math.IsNaN(d)) {
			agg.Max = d
		}
		agg.Avg += d
	}
	if l%2 == 0 {
		agg.Median = (dt[l/2-1] + dt[l/2]) / 2
	} else {
		agg.Median = dt[l/2]
	}
	ratio := math.Pow(10, float64(2))
	agg.Avg = math.Round((agg.Avg/float64(l))*ratio) / ratio
}
