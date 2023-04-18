package acquisition

import (
	"reflect"
)

// Returns all
func getFieldNamesOfFixtureItem(fixture FixtureItem) []string {
	t := reflect.TypeOf(fixture)
	var fieldNames []string
	for i := 0; i < t.NumField(); i++ {
		fieldNames = append(fieldNames, t.Field(i).Name)
	}
	return fieldNames
}
