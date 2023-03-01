package main 

import "fmt"

func main () {
	fmt.Println("Hello world")
	var angka int 
	fmt.Scanln(&angka)
	fmt.Println("angka yang saya input", angka)

	// percabang 1 tingkat
	if angka >= 50 {
		fmt.Println("Inputnya angka yang besar")
	} else if angka >= 10 && angka <= 20 {
		fmt.Println("masih berkembang")
	} else {
		fmt.Println("Inputnya angka yang kecil")
	}
	 
	// percabangan dalam percabangan

	if angka >= 50 {
		fmt.Println("Inputnya angka yang besar")
	} else {
		if angka >= 10 && angka <= 20 {
			fmt.Println("masih berkembang")
		}
		fmt.Println("Inputnya angka yang kecil")
	}


	if angka >= 50 {
		fmt.Println("angka sudah besar")
	} else if  angka >= 20 {
		fmt.Println("sudah dewasa ")

	}

	var umur, tinggi int 
	var kerja string
	fmt.Print("masukan umur")
	fmt.Scanln(&umur)

    
	// 1x cek (jika kondisi pertama tidak terpenuhi, cek pada kondisi selanjutnya)
	if umur >= 20 {
		fmt.Println("Tergolong dewasa")
		fmt.Print("masukan pekerjaan")
		fmt.Scanln(&kerja)
		if kerja == "TNI" {
		fmt.Println("Gratis!")
		} else {
			fmt.Println("Bayar!")
		}
	} else if (umur > 5) {
		fmt.Println("Tergolong anak yang bayar")
	} else {
		fmt.Print("masukan tinggi badan")
		fmt.Scanln(&tinggi)
		if tinggi > 120 {
			fmt.Println("Tergolong anak yang bayar")
		} else {
			fmt.Println("GRATIS")
		} 
	}
    // 2x cek 
	// jika konsidi pertama terpenuhi 
     if umur >= 20 {
		fmt.Println("Tergolong dewasa")
	}
	if (umur > 5){
		fmt.Println("Tergolong anak yang bayar")
	}
}