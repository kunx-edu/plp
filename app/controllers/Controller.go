package controllers

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
}

// json响应
func (c *Controller) json(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
