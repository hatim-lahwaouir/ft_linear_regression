package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dataNormalization(m [][]float64) [][]float64 {
	result := make([][]float64, len(m))
	// creating a copy of the matrix
	copy(result, m)
	for i := range m {
		result[i] = make([]float64, len(m[i]))
		copy(result[i], m[i])
	}

	for y := 0; y < len(m); y++ {

		mx := result[y][0]
		mn := result[y][0]

		for x := 0; x < len(result[y])-1; x++ {
			if mn > result[y][x] {
				mn = result[y][x]
			}

			if mx < result[y][x] {
				mx = result[y][x]
			}
		}

		for x := 0; x < len(result[y])-1; x++ {
			result[y][x] = (result[y][x] - mn) / (mx - mn)
		}
	}

	return result
}

func readData(filePath string) ([][]float64, error) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(strings.NewReader(string(data)))
	// skiping headers
	header, err := r.Read()

	if err != nil {
		return nil, err
	}

	m := make([][]float64, len(header))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// let's convert the record and append to the matrix

		for i, s := range record {
			nbr, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
			if err != nil {
				return nil, err
			}
			m[i] = append(m[i], nbr)
		}
	}
	return m, nil
}

// debuging function
func print(m [][]float64) {

	fmt.Println(len(m), len(m[0]))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j], "  ")
		}
		fmt.Println("")
	}
}
func main() {
	fmt.Println(" -- Reading data -- ")
	m, err := readData("data.csv")

	if err != nil {
		fmt.Println("error", err.Error())
	}
	r := dataNormalization(m)

	print(m)
	print(r)
}
