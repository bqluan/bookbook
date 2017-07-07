package httpx

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, v interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	return json.NewEncoder(w).Encode(v)
}
