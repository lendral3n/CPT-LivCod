package main

import (
	"fmt"
)

func encodeDecode(word string) string {
	// your code here

	shift := 2 // shift digunakan untuk menentukan apakah kita akan menggeser ke depan atau ke belakang

	// Jika kata diawali dengan "<decode>", kita akan menggeser ke belakang
	if word[:8] == "<decode>" {
		shift = -2
		word = word[8:]
	} else if word[:8] == "<encode>" { // Jika kata diawali dengan "<encode>", kita akan menggeser ke depan
		word = word[8:]
	}

	result := "" // Variabel result digunakan untuk menyimpan hasil akhir

	 // Looping setiap karakter dalam kata
	for _, char := range word { // Jika karakter adalah huruf kecil
		if char >= 'a' && char <= 'z' {   // Geser karakter
			char = rune((int(char-'a')+shift+26)%26+'a')
		}
		result += string(char) // Tambahkan karakter yang telah digeser ke result
	}
	return result // Kembalikan hasil akhir
}

func main() {
	// encode
	fmt.Println(encodeDecode("<encode>abcd")) // cdef
	fmt.Println(encodeDecode("<encode>xyz"))  // zab
	// decode
	fmt.Println(encodeDecode("<decode>cdef")) // abcd
	fmt.Println(encodeDecode("<decode>zab"))  // xyz
}
