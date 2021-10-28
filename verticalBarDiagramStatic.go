package main

import (
	"fmt"
)

func main() {
	fmt.Print("Masukkan Jumlah Data yang akan dimasukkan: ")
	var jumlahdata int = 5
	var DataDiagram = make([]int, jumlahdata)
	DataDiagram = []int{2, 4, 1, 3, 0}

	Sorting(DataDiagram)
}
func Grafik(jumlahdata int, max int, DataDiagram []int) {
	for i := max; i >= 1; i-- { // membuat Data Diagram
		for j := 0; j < jumlahdata; j++ {
			if DataDiagram[j] >= i {
				fmt.Print(" | ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Print("\n")
	}
	for i := 0; i < jumlahdata; i++ {
		fmt.Print("---")
	}
	fmt.Print("\n")
	fmt.Print(" ")
	for i := 0; i < jumlahdata; i++ {
		fmt.Print(DataDiagram[i], "  ")
	}
	fmt.Print("\n")
}

func Sorting(data []int) {
	max := data[0]
	for _, value := range data { // Menemukan nilai maximum
		if value > max {
			max = value
		}
	}
	ditukar := true
	for ditukar {
		Grafik(len(data), max, data)
		ditukar = false
		for i := 0; i < len(data)-1; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				ditukar = true
			}
		}
	}
}
