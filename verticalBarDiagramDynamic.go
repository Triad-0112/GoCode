package main

import (
	"fmt"
)

func main() {
	fmt.Print("Masukkan Jumlah Data yang akan dimasukkan: ")
	var jumlahdata int
	fmt.Scanln(&jumlahdata)
	var DataDiagram = make([]int, jumlahdata)
	fmt.Print("Masukkan data secara berurutan dengan spasi sebagai pemisah antar angka: ")
	for i := 0; i < jumlahdata; i++ {
		fmt.Scanf("%d", &DataDiagram[i])
	}
	fmt.Print("\n")
	Sorting(DataDiagram)
}
func Grafik(jumlahdata int, max int, DataDiagram []int) { //fungsi membuat histogram
	for i := max; i >= 1; i-- {
		for j := 0; j < jumlahdata; j++ {
			if DataDiagram[j] >= i {
				fmt.Print(" | ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print(" ")
	for i := 0; i < jumlahdata; i++ {
		fmt.Print(DataDiagram[i], "  ")
	}
	fmt.Print("\n")
}
func Sorting(data []int) { //Proses Sorting
	var max int = data[0]
	for _, value := range data { // Menemukan nilai maximum
		if value > max {
			max = value
		}
	}
	ditukar := true
	for ditukar {
		Grafik(len(data), max, data)
		fmt.Print("\n")
		ditukar = false
		for i := 0; i < len(data)-1; i++ {
			if data[i] > data[i+1] {
				mem := data[i]
				data[i] = data[i+1]
				data[i+1] = mem
				ditukar = true
			}
		}
	}
}
