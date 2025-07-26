package response

import (
	"encoding/json"
	"net/http"
)

type APIResponseWithMeta struct {
	Status  string      `json:"status"`         // "success" or "failed"
	Entity  string      `json:"entity"`         // e.g. "books"
	State   string      `json:"state"`          // e.g. "getAllBooks"
	Message string      `json:"message"`        // e.g. "Success Get All Books"
	Meta    interface{} `json:"meta,omitempty"` // query metadata (search, filter, range, etc.)
	Data    interface{} `json:"data,omitempty"` // actual payload
}

func SuccessWithMeta(w http.ResponseWriter, code int, entity string, state string, message string, meta interface{}, data interface{}) {
	JSONWithMeta(w, code, entity, state, message, meta, data, true)
}

func FailedWithMeta(w http.ResponseWriter, code int, entity string, state string, message string, meta interface{}) {
	JSONWithMeta(w, code, entity, state, message, meta, nil, false)
}

func JSONWithMeta(w http.ResponseWriter, code int, entity, state, message string, meta interface{}, data interface{}, success bool) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	status := "failed"
	if success {
		status = "success"
	}

	json.NewEncoder(w).Encode(APIResponseWithMeta{
		Status:  status,
		Entity:  entity,
		State:   state,
		Message: message,
		Meta:    toSafeData(meta),
		Data:    toSafeData(data),
	})
}
