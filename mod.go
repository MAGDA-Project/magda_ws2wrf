// This module is the forefront of a framework
// of modules that can be used to convert weather
// stations observations in various format into ascii
// WRF format.
// Look at InputFormat for the various input format
// supported.
// Conversion can be done using Convert function.
package dewetra2wrf

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/meteocima/magda_ws2wrf/conversion"
	"github.com/meteocima/magda_ws2wrf/obsreader"
	"github.com/meteocima/magda_ws2wrf/types"
)

// NewReader returns a obsreader.ObsReader that
// read observations stored in this format.
func NewReader() obsreader.ObsReader {
	return obsreader.CsvReader{}
}

// Convert converts a set of observations, saved in
// format, contained in inputpath directory or file,
// reading only data for stations contained in geographicval area
// defined by domain arg, and skipping observation not occurred
// within 15 minutes from date.
// Converted file is saved to outputpath, replacing existing file
// if any, and using os.FileMode(0644) if the file has to be created.
func Convert(inputpath string, domainS string, date time.Time, outputpath string) error {
	domainP, err := types.DomainFromS(domainS)
	if err != nil {
		panic(err)
	}
	domain := *domainP

	sensorsObservations, err := NewReader().ReadAll(inputpath, domain, date)
	if err != nil {
		return err
	}

	results := make([]string, len(sensorsObservations))
	for i, result := range sensorsObservations {
		results[i] = conversion.ToWRFASCII(result)
	}

	resultsS := strings.Join(results, "\n")

	header := fmt.Sprintf(headerFormat, len(results), len(results))

	return os.WriteFile(outputpath, []byte(header+resultsS), os.FileMode(0644))

}

var headerFormat = "TOTAL = %6d, MISS. =-888888.,\n" +
	"SYNOP = %6d, METAR =      0, SHIP  =      0, BUOY  =      0, BOGUS =      0, TEMP  =      0,\n" +
	"AMDAR =      0, AIREP =      0, TAMDAR=      0, PILOT =      0, SATEM =      0, SATOB =      0,\n" +
	"GPSPW =      0, GPSZD =      0, GPSRF =      0, GPSEP =      0, SSMT1 =      0, SSMT2 =      0,\n" +
	"TOVS  =      0, QSCAT =      0, PROFL =      0, AIRSR =      0, OTHER =      0,\n" +
	"PHIC  =  40.00, XLONC = -95.00, TRUE1 =  30.00, TRUE2 =  60.00, XIM11 =   1.00, XJM11 =   1.00,\n" +
	"base_temp= 290.00, base_lapse=  50.00, PTOP  =  5000., base_pres=100000., base_tropo_pres= 20000., base_strat_temp=   215.,\n" +
	"IXC   =     60, JXC   =     90, IPROJ =      1, IDD   =      1, MAXNES=      1,\n" +
	"NESTIX=     60,\n" +
	"NESTJX=     90,\n" +
	"NUMC  =      1,\n" +
	"DIS   =  60.00,\n" +
	"NESTI =      1,\n" +
	"NESTJ =      1,\n" +
	"INFO  = PLATFORM, DATE, NAME, LEVELS, LATITUDE, LONGITUDE, ELEVATION, ID.\n" +
	"SRFC  = SLP, PW (DATA,QC,ERROR).\n" +
	"EACH  = PRES, SPEED, DIR, HEIGHT, TEMP, DEW PT, HUMID (DATA,QC,ERROR)*LEVELS.\n" +
	"INFO_FMT = (A12,1X,A19,1X,A40,1X,I6,3(F12.3,11X),6X,A40)\n" +
	"SRFC_FMT = (F12.3,I4,F7.2,F12.3,I4,F7.3)\n" +
	"EACH_FMT = (3(F12.3,I4,F7.2),11X,3(F12.3,I4,F7.2),11X,3(F12.3,I4,F7.2))\n" +
	"#------------------------------------------------------------------------------#\n"
