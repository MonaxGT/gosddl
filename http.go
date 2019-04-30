package gosddl

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello")
}

func (app *ACLProcessor) decode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["sddl"] != "" {
		sddl := params["sddl"]
		err := app.findGroupIndex(sddl)
		if err != nil {
			log.Println("Wrong SDDL string")
		}
		json.NewEncoder(w).Encode(app.Rights)
		app.Rights = permissions{}
		return
	}
}

// HTTPHandler start http serve
func (app *ACLProcessor) HTTPHandler(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/sddl", getInfo).Methods("GET")
	router.HandleFunc("/sddl/{sddl}", app.decode).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
