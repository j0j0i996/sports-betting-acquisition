package acquisition

import (
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type FixtureItem struct {
	Date     time.Time
	Div      string
	HomeTeam string
	AwayTeam string
	B365H    float64
	B365A    float64
	B365D    float64
}

// FetchData requests csv from URL and saves / returns it as string
// example url: https://www.football-data.co.uk/mmz4281/1718/N1.csv
func fetchHistoricData(year_code string, league string) ([][]string, error) {

	url := BASE_URL + "/" + HIST_ODD_ENDPOINT + "/" + year_code + "/" + LeagueMap[league] + ".csv"
	data, err := ReadCSVFromUrl(url)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// stringInSlice checks if a single string is in a list of strings
func isStringInSlice(str_to_check string, slice []string) bool {
	for _, b := range slice {

		if b == str_to_check {
			return true
		}
	}
	return false
}

// filterColumnsByName filters an array of array by only keeping the columns
// in which the name in the first row matches one of the col_names_to_keep input
// It retuns the data as map, in which the key is the column name
func modelData(data [][]string) ([]FixtureItem, error) {

	// Get index of columns
	idx_map := make(map[string]int)

	// Get column names to keep
	cols_to_keep := getFieldNamesOfFixtureItem(FixtureItem{})

	// Find column number of cols_to_keep in returned csv
	for index, element := range data[0] {

		if isStringInSlice(element, cols_to_keep) {
			idx_map[element] = index
		}
	}
	if len(cols_to_keep) != len(idx_map) {
		return []FixtureItem{}, errors.New("One of the excpected column names is not part in csv header")
	}

	// Model data
	var modelled_data []FixtureItem
	for i := 1; i < len(data); i++ {
		fixture := FixtureItem{}
		writeFixture := true
		for fieldName, index := range idx_map {
			v := reflect.ValueOf(&fixture).Elem()
			f := v.FieldByName(fieldName)

			if f.IsValid() {
				if f.CanSet() {
					if f.Type().String() == "float64" {
						// handle if no odds are available for a game
						if data[i][index] == "" {
							writeFixture = false
						} else {
							float_number, err := strconv.ParseFloat(data[i][index], 32)
							if err != nil {
								return []FixtureItem{}, errors.Wrap(
									err,
									"Could not convert string to float: "+data[i][index]+" for row: "+strconv.Itoa(i)+" and field name "+fieldName)
							}
							f.SetFloat(math.Floor(float_number*10000) / 10000)
						}
					} else if f.Type().String() == "string" {
						f.SetString(data[i][index])
					} else if f.Type().String() == "time.Time" {
						var parsed_time time.Time
						var err error
						if len(data[i][index]) == 10 {
							parsed_time, err = time.Parse("02/01/2006", data[i][index])
						} else if len(data[i][index]) == 8 {
							parsed_time, err = time.Parse("02/01/06", data[i][index])
						} else {
							err = errors.New("unknown time format")
						}
						if err != nil {
							panic(err)
						}
						f.Set(reflect.ValueOf(parsed_time))
					} else {
						return []FixtureItem{}, errors.New("Unknown Field type " + f.Type().String())
					}
				} else {
					return []FixtureItem{}, errors.New("Field is not settable")
				}
			} else {
				return []FixtureItem{}, errors.New("Field not found")
			}
		}
		if writeFixture {
			modelled_data = append(modelled_data, fixture)
		}
	}

	return modelled_data, nil
}

func GetHistoricData(year_code string, league string) ([]FixtureItem, error) {
	// Get Data from URL as slcie of slices, with first array being the header
	data, err := fetchHistoricData(year_code, league)
	if err != nil {
		return []FixtureItem{}, err
	}

	modelledData, err := modelData(data)

	return modelledData, err
}
