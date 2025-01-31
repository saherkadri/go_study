package rest_api

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var csvData [][]string

func init() {
	fmt.Println("my test env", os.Getenv("TEST_ENV"))
	if os.Getenv("TEST_ENV") != "" && os.Getenv("TEST_ENV") == "true" {
		fmt.Println("Test environment detected, skipping CSV loading.")
		return // Skipping to load CSV file in test environment
	}
	var err error
	csvData, err = loadCSV("matrix.csv")
	if err != nil {
		fmt.Printf("Error loading CSV: %v\n", err)
		os.Exit(1) // Exiting the program if the CSV file fails to load
	}
}

// This function will be called by init() and it will load and save csv file data into memory.
// So that we load file only once and access the data throughout endpoints.
func loadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error while reading file: %v", err)
	}
	fmt.Println("Matrix loaded successfully")
	return records, nil
}

// Helper function to set csvData for testing purposes
func SetCSVData(data [][]string) {
	csvData = data
}

func Echo(w http.ResponseWriter, r *http.Request) {
	//records, _ := loadCSV(w, r) -> Tried loading csv in all endpoints then optimised the code to load once. Hence commented this.
	var response string
	for _, row := range csvData {
		fmt.Println("rows:", row)
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	response = strings.TrimSuffix(response, "\n")
	fmt.Fprint(w, response)
}

func Invert(w http.ResponseWriter, r *http.Request) {
	//records, _ := loadCSV(w, r)
	records := csvData
	if len(records) == 0 {
		return
	}

	// Creating new [][]matrix
	transposed := make([][]string, len(records[0]))
	for i := range transposed {
		transposed[i] = make([]string, len(records))
	}
	for i := range records {
		for j := range records[i] {
			transposed[j][i] = records[i][j]
		}
	}
	var response string
	for _, row := range transposed {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	response = strings.TrimSuffix(response, "\n")
	fmt.Fprint(w, response)
}

func Flatten(w http.ResponseWriter, r *http.Request) {
	//records, _ := loadCSV(w, r)
	var response string
	for _, row := range csvData {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}
	response = strings.TrimSuffix(response, ",")
	fmt.Fprint(w, response)
}

func Sum(w http.ResponseWriter, r *http.Request) {
	//records, _ := loadCSV(w, r)
	var response int
	for _, row := range csvData {
		fmt.Println("myrow:", row)
		for _, ele := range row {
			intValue, err := strconv.Atoi(ele)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			response += intValue
		}
	}
	fmt.Fprint(w, response)
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	//records, _ := loadCSV(w, r)
	response := 1
	for _, row := range csvData {
		fmt.Println("myrow:", row)
		for _, ele := range row {
			intValue, err := strconv.Atoi(ele)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			response *= intValue
		}
	}
	fmt.Fprint(w, response)
}
