// This module implements a console command
// that can be used to convert observation
// from CSV to ascii WRF format.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	ws2wrf "github.com/meteocima/magda_ws2wrf"
	"github.com/meteocima/magda_ws2wrf/types"
)

func main() {
	outfile := flag.String("outfile", "./out", "where to save converted file")
	domainS := flag.String("domain", "", "domain to filter stations to convert [MinLat,MaxLat,MinLon,MaxLon]")
	dateS := flag.String("date", "", "date and hour to filter stations data to convert [YYYYMMDDHH]")

	flag.Parse()
	var date time.Time
	var domain *types.Domain
	var err error

	if *dateS != "" {
		if date, err = time.Parse("2006010215", *dateS); err != nil {
			fmt.Fprintf(os.Stderr, "Invalid date option: %s\n", err.Error())
			flag.Usage()
			os.Exit(1)
		}
	}

	if domain, err = types.DomainFromS(*domainS); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid domain option: %s\n", err.Error())
		flag.Usage()
		os.Exit(1)
	}

	inputFiles := flag.Args()

	if len(inputFiles) == 0 {
		fmt.Fprintf(os.Stderr, "No input files specified\n")
		flag.Usage()
		os.Exit(1)
	}

	err = ws2wrf.Convert(inputFiles, *domain, date, *outfile)

	if err != nil {
		log.Fatal(err)
	}
}
