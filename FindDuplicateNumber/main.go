package main

import (
	"fmt"


)

func findDuplicationNumber(numbers []int) []int {
    duplicateMap := make(map[int]int) // map untuk melacak jumlah kemunculan setiap angka
    duplicates := []int{} // slice untuk menyimpan angka yang terduplikasi

    // loop untuk menghitung jumlah kemunculan setiap angka dalam slice numbers
    for _, number := range numbers {
        duplicateMap[number]++ // menamvahkan jumlah angka ke dalam map
    }

    // loop untuk mengecek map dan menemukan angka yang lebih kemunculan nya dari satu
    for number, count := range duplicateMap {
        if count > 1 { // jika jumlah kemunculan nya lebih dari satu
            duplicates = append(duplicates, number) // tambahkan angka tersebut ke dalam slice duplicates
        }
    }

    // kembalikan slice duplicates
    return duplicates
}

func main() {
	fmt.Println(findDuplicationNumber([]int{1, 1}))                        // [1]
	fmt.Println(findDuplicationNumber([]int{1, 2, 3, 4}))                   // []
	fmt.Println(findDuplicationNumber([]int{1, 2, 1, 2, 2, 3, 4, 5, 5, 5, 5})) // [1, 2, 5]
}
