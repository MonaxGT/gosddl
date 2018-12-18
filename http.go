package gosddl

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var permisson Permissons

func GetInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello")
}

func Decode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["sddl"] != "" {
		sddl := params["sddl"]
		permisson.FindGroupIndex(sddl)
		json.NewEncoder(w).Encode(permisson)
		return
	}
}

func HttpHandler(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/sddl", GetInfo).Methods("GET")
	router.HandleFunc("/sddl/{sddl}", Decode).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
