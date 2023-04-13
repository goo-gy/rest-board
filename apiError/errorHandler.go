package apiError

import (
	"encoding/json"
	"net/http"
)

func HandleError(res http.ResponseWriter, errorBody ErrorBody) {
	errorJson, _ := json.Marshal(errorBody)
	res.WriteHeader(http.StatusBadRequest)
	res.Write(errorJson);
}