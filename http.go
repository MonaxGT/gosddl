package gosddl

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var permisson Permissons

func getInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello")
}

func decode(w http.ResponseWriter, r *http.Request) {
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
	router.HandleFunc("/sddl", getInfo).Methods("GET")
	router.HandleFunc("/sddl/{sddl}", decode).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
