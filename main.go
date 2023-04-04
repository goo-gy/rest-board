package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/restBoard/posting"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posting", posting.HandleCreatePosting).Methods("POST")
	fmt.Println("Serve on :8000")
	http.ListenAndServe(":8000", router)
}