package main

import (
	"flag"

	"github.com/MonaxGT/gosddl"
)

func main() {
	apiPtr := flag.Bool("api", false, "Use API mode")
	apiPortPtr := flag.String("port", ":8000", "Default port 8000")
	fileSIDs := flag.String("f", "", "File with users's SIDs")
	flag.Parse()
	gosddl.Processor(*apiPtr, *apiPortPtr, *fileSIDs)
}
