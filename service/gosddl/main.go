package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/MonaxGT/gosddl"
)

func main() {
	var permisson gosddl.Permissons
	apiPtr := flag.Bool("api", false, "a bool")
	apiPortPtr := flag.String("port", ":8000", "Default port 8000")
	flag.Parse()
	if *apiPtr {
		fmt.Println("API Interface started on port", *apiPortPtr)
		gosddl.HttpHandler(*apiPortPtr)
	} else if flag.Args() != nil {
		permisson.FindGroupIndex(flag.Args()[0])
		b, err := json.Marshal(permisson)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
}
