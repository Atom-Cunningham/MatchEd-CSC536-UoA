package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	filename := GetFilename()
	file := GetFile(filename)
	if file == nil {
		return
	}
	records := ReadCSV(file)
	file.Close()
	jsons := GetJson(records)
	fmt.Println(string(jsons))

	//writeExpectedOutput(filename, jsons)
}

// grabs the second run command arg as filename string
func GetFilename() string {
	// Ensure a filename argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Missing argument\nUsage: go run jsonMaker.go <filename>")
		os.Exit(1)
	}
	return os.Args[1]
}

// opens file in directory with given name
func GetFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s, \n%s\n", filename, err)
		return nil
	}
	return file
}

// converts csv file into string matrix
func ReadCSV(file *os.File) [][]string {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading csv: %s\n", err)
		return nil
	}
	return records
}

// converts records to json format
func GetJson(records [][]string) []byte {
	keys := records[0]
	var jsonObjects []map[string]string

	for _, row := range records[1:] {
		newMap := make(map[string]string)
		for i, value := range row {
			if i < len(keys) {
				newMap[keys[i]] = value
			}
		}
		jsonObjects = append(jsonObjects, newMap)
	}

	jsonData, err := json.MarshalIndent(jsonObjects, "", "	")
	if err != nil {
		fmt.Printf("Error converting to json: %s\n", err)
		return nil
	}
	return jsonData
}

// writeExpectedOutput writes the actual output to the expected file for comparison
func writeExpectedOutput(filename string, data []byte) error {
	expectedFilename := strings.Replace(filename, ".csv", "_expected.json", 1)
	return os.WriteFile(expectedFilename, data, 0644)
}
