package transform

import (
	"testing"
)

type InStruct struct {
	Name   string
	Weight float64
	Height int
}

func TestMapToStruct(t *testing.T) {
	testCases := []struct {
		Name     string
		InStruct *InStruct
		InMap    map[string]interface{}
	}{
		{
			Name: "not all fields",
			InMap: map[string]interface{}{
				"Name":   "Dasha",
				"Weight": 68.5,
			},
			InStruct: &InStruct{},
		},
		{
			Name: "extra fields",
			InMap: map[string]interface{}{
				"Name":   "Dasha",
				"Weight": 68.5,
				"Height": 160,
				"extra":  6,
			},
			InStruct: &InStruct{},
		},
		{
			Name: "nil",
			InMap: map[string]interface{}{
				"Name":   nil,
				"Weight": 68.5,
				"Height": 160,
			},
			InStruct: &InStruct{},
		},
		{
			Name: "types do not match",
			InMap: map[string]interface{}{
				"Name":   "Dasha",
				"Weight": "68.5",
				"Height": 160,
			},
			InStruct: &InStruct{},
		},
	}

	t.Run(testCases[0].Name, func(t *testing.T) {
		if out_error := MapToStruct(testCases[0].InStruct, testCases[0].InMap); out_error != nil {
			t.Fatalf("test %s - fail", testCases[0].Name)
		}
	})

	t.Run(testCases[1].Name, func(t *testing.T) {
		if out_error := MapToStruct(testCases[1].InStruct, testCases[1].InMap); out_error != nil {
			t.Fatalf("test %s - fail", testCases[1].Name)
		}
	})

	t.Run(testCases[2].Name, func(t *testing.T) {
		if out_error := MapToStruct(testCases[2].InStruct, testCases[2].InMap); out_error == nil {
			t.Fatalf("test %s - fail", testCases[2].Name)
		}
	})

	t.Run(testCases[3].Name, func(t *testing.T) {
		if out_error := MapToStruct(testCases[3].InStruct, testCases[3].InMap); out_error == nil {
			t.Fatalf("test %s - fail", testCases[3].Name)
		}
	})
}
