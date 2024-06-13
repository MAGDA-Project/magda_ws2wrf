package obsreader

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/meteocima/magda_ws2wrf/elevations"
	"github.com/meteocima/magda_ws2wrf/types"
)

// CsvReader reads observations from CSV files
type CsvReader struct{}

// ReadAll implements ObsReader for CsvReader
func (r CsvReader) ReadAll(dataPath string, domain types.Domain, date time.Time) ([]types.Observation, error) {
	observations := []types.Observation{}

	obsF, err := os.Open(dataPath)
	if err != nil {
		return nil, err
	}
	defer obsF.Close()
	csvReader := csv.NewReader(obsF)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	const ID = 0
	const DATE = 1
	const TEMP = 2
	const NAME = 3
	const LON = 4
	const LAT = 5
  const HEIGHT = 6

	for _, row := range data[1:] {
		var obs types.Observation
		var err error
		if row[LAT] == "NA" {
			continue
		}
		obs.Lat, err = strconv.ParseFloat(row[LAT], 64)
		if err != nil {
			return nil, err
		}
		obs.Lon, err = strconv.ParseFloat(row[LON], 64)
		if err != nil {
			return nil, err
		}

		obs.ObsTimeUtc, err = time.Parse("2006-01-02 15:04:05", row[DATE])
		if err != nil {
			return nil, err
		}

		if obs.Lat <= domain.MaxLat && obs.Lat >= domain.MinLat &&
			obs.Lon <= domain.MaxLon && obs.Lon >= domain.MinLon &&
			(obs.ObsTimeUtc.Sub(date).Abs().Minutes() <= 15 || date.IsZero()) {

			obs.StationID = row[ID]
      if len(row) == 7 {
          elev, err := strconv.ParseFloat(row[HEIGHT], 64)
          if err != nil {
            return nil, err
          }
          obs.Elevation = elev
      } else {
					obs.Elevation = elevations.GetFromCoord(obs.Lat, obs.Lon)
      } 
			obs.StationName = row[NAME]
			temp, err := strconv.ParseFloat(row[TEMP], 64)
			if err != nil {
				return nil, err
			}
			obs.Metric.TempAvg = types.Value(temp)

			//obs.Metric.Pressure = types.Value((obs.Metric.PressureMax + obs.Metric.PressureMin) / 2)

			// convert temperatures from °celsius to °kelvin
			obs.Metric.TempAvg += 273.15

			observations = append(observations, obs)
		}
	}

	return observations, nil
}
