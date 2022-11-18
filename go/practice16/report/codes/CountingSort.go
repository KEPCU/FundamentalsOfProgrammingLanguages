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
			countingSort(list)

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

	file, err := os.Create("countingSortTimeGo.csv")
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

func countingSort(list []int) {
	minItem := Min(list)
	countList := Zero(Max(list) - minItem + 1)
	outputList := Zero(len(list))

	for i := 0; i < len(list); i++ {
		countList[list[i]-minItem] += 1
	}

	for i := 1; i < len(countList); i++ {
		countList[i] += countList[i-1]
	}

	for i := len(list) - 1; i >= 0; i-- {
		outputList[countList[list[i]-minItem]-1] = list[i]
		countList[list[i]-minItem] -= 1
	}

	for i := 0; i < len(list); i++ {
		list[i] = outputList[i]
	}
}

func Max(list []int) int {
	max := list[0]
	for i := 0; i < len(list); i++ {
		if max < list[i] {
			max = list[i]
		}
	}
	return max
}

func Min(list []int) int {
	min := list[0]
	for i := 0; i < len(list); i++ {
		if min > list[i] {
			min = list[i]
		}
	}
	return min
}

func Zero(length int) []int {
	list := []int{}
	for i := 0; i < length; i++ {
		list = append(list, 0)
	}
	return list
}
