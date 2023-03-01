package main

import (
	"fmt"
	"math/rand"
)

func main() {
  var angka1, angka2 float64
  var hasil float64
  var kalkulator int
  var Menu float64
  var state string

  random := rand.NewSource(time.Now().UnixMilli())
  rnd := rand.New(random) 
  randomInt := rnd.Intn(100)
  fmt.Println(randomInt)

  
  fmt.Println("Pilih menu :")
  fmt.Println("1. pertambahan")
  fmt.Println("2. pengurangan")
  fmt.Println("3. perkalian")
  fmt.Println("4. pembagian")
  fmt.Scanln(&kalkulator)
  // if kalkulator >= 1 && kalkulator <= 4 {

  // }
 //if kalkulator >= 1 && kalkulator <= 4
 if kalkulator == 1 {
    fmt.Println("Masukkan angka pertama")
    fmt.Scanln(&angka1)
    fmt.Println("Masukkan angka kedua")
    fmt.Scanln(&angka2)
    hasil = angka1 + angka2
    fmt.Println("Hasilnya :", hasil)
  } else if kalkulator == 2 {
    fmt.Println("Masukkan angka pertama")
    fmt.Scanln(&angka1)
    fmt.Println("Masukkan angka kedua")
    fmt.Scanln(&angka2)
    hasil = angka1 + angka2
    fmt.Println("Hasilnya :", hasil)
  } else if kalkulator == 3 {
    fmt.Println("Masukkan angka pertama")
    fmt.Scanln(&angka1)
    fmt.Println("Masukkan angka kedua")
    fmt.Scanln(&angka2)
    hasil = angka1 + angka2
    fmt.Println("Hasilnya :", hasil)
  } else if kalkulator == 4 {
    fmt.Println("Masukkan angka pertama")
    fmt.Scanln(&angka1)
    fmt.Println("Masukkan angka kedua")
    fmt.Scanln(&angka2)
    hasil = angka1 + angka2
    fmt.Println("Hasilnya :", hasil)
  } else {
    fmt.Println("Input salah!")
    return
  } 

  // lebih singkat praktis 
for isRunning {
  fmt.Println("Menu\n1.pejumlahan\n2.pengurangan\n3.perkalian\n4.pembagian")
  fmt.Print("Masukan Menu")
  fmt.Scanln(&Menu)

  if Menu > 0 && Menu <= 4 {
	fmt.Print("Masukkan Angka ke 1")
	fmt.Scanln(&angka1)
	fmt.Print("Masukkan Angka ke 2")
	fmt.Scanln(&angka2)
	if Menu == 1 {
		state = "pemjumlaha"
		hasil = angka1 + angka2
	} else if Menu == 2 {
		state = "pengurangan"
		hasil = angka1 + angka2 
	} else if Menu == 3 {
	    state = "perkalian"
		hasil = angka1 * angka2 
	} else if Menu == 4 {
		state = "pembagian"
		hasil = angka1 / angka2
    }

	fmt.Printf("Hasil%s adalah %d", hasil,state)

  } 

  switch Menu {
  case 1 :
	state = "penjumlahan"
	hasil = angka1 + angka2 
  case 2 :
	state = "perkalian"
	hasil = angka1 + angka2
  case 3 :
	state = "pengurahan"
	hasil = angka1 + angka2
  case 4 :
	state = "pembagian"
	hasil = angka1 / angka2
} 
  
fmt.Printf("Hasil %s adalah %d", hasil, state)

}

