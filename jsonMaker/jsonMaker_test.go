package main

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestConvertCSVRowsToJSON(t *testing.T) {

	testCases := []string{"test1.csv"}

	for _, filename := range testCases {
		t.Run(filename, func(t *testing.T) {
			runTest(filename, t)
		})
	}
	// get result

}

// readExpectedOutput reads expected JSON from the file
func readExpectedOutput(filename string) ([]byte, error) {
	expectedFilename := strings.Replace(filename, ".csv", "_expected.json", 1)
	data, err := os.ReadFile(expectedFilename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func runTest(filename string, t *testing.T) {
	file := GetFile(filename)
	if file == nil {
		return
	}
	records := ReadCSV(file)
	file.Close()
	jsons := GetJson(records)

	// get expected result
	expected, err := readExpectedOutput(filename)
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}

	// Unmarshal both expected and actual data for comparison
	var expectedJSON, actualJSON []map[string]string
	if err := json.Unmarshal(expected, &expectedJSON); err != nil {
		t.Fatalf("Error unmarshaling expected JSON: %v", err)
	}
	if err := json.Unmarshal(jsons, &actualJSON); err != nil {
		t.Fatalf("Error unmarshaling actual JSON: %v", err)
	}

	// Compare the two slices of maps
	if !reflect.DeepEqual(expectedJSON, actualJSON) {
		t.Errorf("Expected:\n%v\nGot:\n%v", expectedJSON, actualJSON)
	}

}
