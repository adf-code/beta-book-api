package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  string      `json:"status"`         // "success" or "failed"
	Entity  string      `json:"entity"`         // e.g. "books"
	State   string      `json:"state"`          // e.g. "getAllBooks"
	Message string      `json:"message"`        // e.g. "Success Get All Books"
	Data    interface{} `json:"data,omitempty"` // actual payload
}

func Success(w http.ResponseWriter, code int, entity string, state string, message string, data interface{}) {
	JSON(w, code, entity, state, message, data, true)
}

func Failed(w http.ResponseWriter, code int, entity string, state string, message string) {
	JSON(w, code, entity, state, message, nil, false)
}

func JSON(w http.ResponseWriter, code int, entity, state, message string, data interface{}, success bool) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	status := "failed"
	if success {
		status = "success"
	}

	json.NewEncoder(w).Encode(APIResponse{
		Status:  status,
		Entity:  entity,
		State:   state,
		Message: message,
		Data:    toSafeData(data),
	})
}
