package main

import "fmt"

type karyawan struct {
	nama    string
	usia    int
	jabatan string
	gaji    int
}

func main() {
	var mana [3]string
	mana[0] = "hai"
	mana[1] = "coy"
	mana[2] = "fay"
	// var x string
	// var y string
	// x = "Hello"
	// y = " World!"
	x := "Hello"
	x += " Euy!"

	var karyawan1 = karyawan{nama: "irwan", gaji: 120000000000}

	fmt.Println(x)
	// fmt.Println(y)
	fmt.Println(mana)
	fmt.Println(karyawan1)
}
