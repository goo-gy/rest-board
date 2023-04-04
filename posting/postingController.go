package posting

import (
	"encoding/json"
	"net/http"

	"github.com/restBoard/apiError"
)

func HandleCreatePosting(res http.ResponseWriter, req *http.Request) {
	// validation
	postingRequest := PostingRequest{}
	err := json.NewDecoder(req.Body).Decode(&postingRequest)
	if err != nil {
		apiError.HandleError(res, apiError.ErrInvalidInput);
		return
	}
	// create
	postingResponse := CreatePosting(postingRequest)
	createdPostingJson, _ := json.Marshal(postingResponse)
	res.WriteHeader(http.StatusOK)
	res.Write(createdPostingJson)
}