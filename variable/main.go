//variables
// string
// int
// float
// bool true or false

package main

import "fmt"

func main() {
	// cara 1 deklarasi variable
	// var nama string = "Seriusman Waruwu"
	// var umur int = 20

	//cara 2 deklarasi variable
	// var nama string
	// var umur = 20

	//cara 3 deklarasi variable
	// nama := "Seriusman Waruwu"
	// umur := 20

	// fmt.Println(nama)
	// fmt.Println(umur)

	// cara 4 deklarasi variable
	namaDepan, namaBelakang := "Seriusman", "Waruwu"
	umurSekarang, umurTahunDepan := 20, 21

	fmt.Println(namaDepan, namaBelakang)
	fmt.Println(umurSekarang, umurTahunDepan)

}
