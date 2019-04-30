package main

import (
	"flag"

	"fmt"
	"github.com/MonaxGT/gosddl"
)

func main() {
	apiPtr := flag.Bool("api", false, "Use API mode")
	apiPortPtr := flag.String("port", ":8000", "Default port 8000")
	fileSIDs := flag.String("f", "", "File with users's SIDs")
	flag.Parse()
	var app gosddl.ACLProcessor
	app.File = *fileSIDs
	if *apiPtr {
		fmt.Println("API Interface started on port", *apiPortPtr)
		app.HTTPHandler(*apiPortPtr)
	} else if flag.Args() != nil {
		err := app.Processor(flag.Args()[0])
		if err != nil {
			panic(err)
		}
	}
	panic("You should give me SDDL string or use API mode")
}
