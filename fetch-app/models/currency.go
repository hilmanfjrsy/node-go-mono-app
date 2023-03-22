package models

type Currency struct {
	Success   bool                   `json:"success"`
	Timestamp interface{}            `json:"timestamp"`
	Base      string                 `json:"base"`
	Date      string                 `json:"date"`
	Rates     map[string]interface{} `json:"rates"`
}
