package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
)

func main() {
    var filterValue  string
    // Open the CSV file for reading
    file, err := os.Open("work.csv")
    if err != nil {
	panic(err)
        return
    }
    defer file.Close()

    // Parse the CSV file
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
	panic(err)
        return
    }

    // Define the criteria for filtering (for example, filter rows where column 2 equals "filter_value")
    filterColumn := 4      // Adjust the column index (0-based) you want to filter on
    filterValue = "Global Business Mobility"
    // Create a new CSV file for writing the filtered data
    outFile, err := os.Create("filtered_output.csv")
    if err != nil {
	panic(err)
        return
    }
    defer outFile.Close()

    // Create a CSV writer for the output file
    writer := csv.NewWriter(outFile)
    defer writer.Flush()

    // Iterate through the records, filter, and write to the output file
    for _, record := range records {
        if len(record) > filterColumn && strings.Contains(record[filterColumn],filterValue)==true {
		err := writer.Write(record)
		if err!=nil{
			panic(err)
		}
        }
    }

    fmt.Println("Filtered data has been written to filtered_output.csv")
}
