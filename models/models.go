package models

type ResponseBody struct {
	Success bool               `json:"success"`
	Rates   map[string]float64 `json:"rates"`
}
