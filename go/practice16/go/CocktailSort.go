package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	sizes := []int{100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000}
	times := [][]float64{}

	for k := 0; k < 5; k++ {
		times = append(times, []float64{})
		fmt.Println(":::::::::", k)
		for i := 0; i < len(sizes); i++ {
			list := []int{}
			fileName := "data" + strconv.Itoa(sizes[i]) + ".csv"

			list = Read(fileName)

			start := time.Now()
			cocktailSort(list)

			times[k] = append(times[k], time.Since(start).Seconds())

			fmt.Println("finished: ", sizes[i])
		}
	}
	times = append(times, []float64{})
	Write(times)
}

func Read(filePath string) []int {
	list := []int{}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
	}

	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			number, err := strconv.Atoi(data[i][j])
			if err == nil {
				list = append(list, number)
			}
		}
	}

	return list
}

func Write(data [][]float64) {
	records := [][]string{{}, {}}

	file, err := os.Create("cocktailSortTimesGo.csv")
	defer file.Close()

	if err != nil {
		fmt.Println("failed to open file", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < len(data[0]); i++ {
		j := 0.0
		for k := 0; k < 5; k++ {
			j += data[k][i]
		}
		data[5] = append(data[5], j/5.0)
	}

	for i := 0; i < 6; i++ {
		records = append(records, []string{})
		for k := 0; k < len(data[i]); k++ {
			j := data[i][k]
			records[i] = append(records[i], strconv.FormatFloat(j, 'E', -1, 64))
		}
	}

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			fmt.Println("error writing record to file", err)
		}
	}
}

func cocktailSort(list []int) {
	last := len(list) - 1
	for {
		swapped := false
		for i := 0; i < last; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
		swapped = false
		for i := last - 1; i >= 0; i-- {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
