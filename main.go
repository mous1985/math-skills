package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Average(data []float64) float64 {
	moyenne := 0.0
	sum := 0.0
	m := len(data)
	for _, i := range data {
		sum += i
	}
	if m != 0 {
		moyenne = sum / float64(m)
		return moyenne
	}
	return moyenne
}

func Mediane(data []float64) float64 {
	mediane := 0.0
	m := len(data)
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	if len(data)%2 == 0 {
		mediane = (data[(m/2)-1] + data[(m)/2]) / 2
	} else {
		mediane = data[(m / 2)]
	}

	return mediane
}

func Variance(data []float64) float64 {
	variance := 0.0

	for _, k := range data {
		s := k - float64(Average(data))
		if len(data) != 0 {
			variance += math.Pow(s, 2) / float64(len(data))
		}
	}
	return variance
}

func standard_deviation(data []float64) float64 {
	ecarT := math.Sqrt(float64(Variance(data)))

	return ecarT
}

func main() {
	var data []float64
	content, err := os.ReadFile("data.txt") // lire le fichier
	if err != nil {
		fmt.Println(err)
		return
	}

	cont := string(content)
	// fmt.Print(cont)
	str := strings.Split(cont, "\n")

	// fmt.Print(str)
	for i := 0; i < len(str); i++ {
		if str[i] == "" {
			continue
		}

		conv, _ := strconv.ParseFloat(str[i], i)

		data = append(data, conv)

	}
	fmt.Println("Average :", math.Round(Average(data)))
	fmt.Println("Mediane :", math.Round(Mediane(data)))
	fmt.Println("Variance :", math.Round(Variance(data)))
	fmt.Println("Standard Deviation :", math.Round(standard_deviation(data)))
}
