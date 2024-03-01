package main
import "fmt"

func main(){
	const pi = 3.14
	var radius, tinggi, x float64
	
	fmt.Print("Masukkan Radius Tabung : ")
	fmt.Scan(&radius)

	fmt.Print("Masukkan Tinggi Tabung : ")
	fmt.Scan(&tinggi)

	x = pi * (radius*radius) * tinggi
	fmt.Printf("Hasil Volume Tabung : %v", x)
}