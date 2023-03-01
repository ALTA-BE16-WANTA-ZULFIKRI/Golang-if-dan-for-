package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	var j int 
	for j < 10 {
		fmt.Println(j)
		j++
	}

	var limit = 500 
	var uang = 0 
	var x = 1 
	for uang < limit {
		fmt.Printf("perulangan ke - %d")
		if x%2 == 0{
			uang -= 25
		} else {
			uang += 50
		}
		x++

    }


	var utang = 2000
	var bayar = 100 
	var tenor = 0

	for utang >= bayar {
		tenor ++ 
		utang -= bayar 
		fmt.Printf("pembayaran ke - %d, sisa utang - %d\n", tenor, utang)
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i)
		if 10/i == 5 {
			break
		}
	}

	for j := 1; j < 10; j++ {
		if j%2 == 0 && j%4 == 0 {
			fmt.Println("GENAP, ", j)
		} else if j%2 == 0{
			continue
		} else {
			fmt.Println("YANG LAIN", j)
		}
	}

	fmt.Println("=============")

	for i := 0; i < 10; i ++ {
		for j := 9; j >= i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}