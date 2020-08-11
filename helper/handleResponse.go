package helper

import(
	"net/http"
	"encoding/json"
)

type Response struct {
	Result  int         `json:"result"`
	Status  string      `json:"status"`
	Message string      `json:"pesan"`
	Data    interface{} `json:"data,omitempty"`
}

func HandleResponse(w http.ResponseWriter, result int, data interface{}) {
	status, message := getStatusMessage(result)

	response := Response{
		Result:  result,
		Status:  status,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getStatusMessage(result int) (string, string) {
	var status, message string
	switch result {
	case 200:
		status = "success"
		message = ""
	case 500:
		status = "unknown"
		message = "internal server error"
	case 501:
		status = "failed"
		message = "data tidak ditemukan"
	}
	return status, message
}
