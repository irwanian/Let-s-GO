package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var x [3]int
	// x[1] = 52

	// fmt.Println(x)

	// ABOVE IS A ARRAY, GO DON'T REALLY RELY ON ARRAY, USE SLICES INSTEAD

	// SLICES WILL BE SHOWN BELOW

	// y := type{values} composite literal, slice allow us to group values of same type
	y := []int{1, 3, 53, 5, 21}

	fmt.Println(y)

	x := []string{"34", "21", "23", "18", "22", "54"}

	for i, v := range x {
		output := ""
		output += strconv.Itoa(i)
		output += "--"
		output += v
		if i < len(x) {
			fmt.Println(output)
		}
	}

	// y = append(y, 21, 22, 45, 3, 31)

	z := []int{211, 222, 54, 33, 11}

	y = append(y, z...)
	fmt.Println(y)

	f := []string{"santuy", "yeah"}
	x = append(x, f...)
	fmt.Println(x)

	aa := []int{1, 32, 12, 52, 21}
	bb := []int{212, 231, 4243, 453}

	// SPREADING THE SLICE
	cc := append(aa, bb...)
	// MULTIDIMENSIONAL SLICE
	dd := [][]int{aa, bb}

	fmt.Println(cc)
	fmt.Println(dd)

	// MAPPING IN GO

	m := map[string]int{
		"Gaji": 70000000,
		"Usia": 33,
	}

	fmt.Println(m["Usia"])

	m["Bonus"] = 300000000

	fmt.Println(m)
	practiceMapping()

}

func practiceMapping() {
	myLife := map[string]int{
		"usia":          33,
		"gaji":          55000000,
		"bonus_tahunan": 300000000,
	}
	fmt.Println("gaji gw Rp.", myLife["gaji"])
	if v, ok := myLife["bonus_tahunan"]; ok {
		fmt.Println("value =", v, myLife)
		delete(myLife, "bonus_tahunan")
		fmt.Println(myLife)
	} else {
		fmt.Println("no such key exists")
	}
}
