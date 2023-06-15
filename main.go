package main

import (
	"fmt"
	"learn_unit_test_v2/helper"
)

func main() {
	a := helper.Person{
		Name: "Ridho Galih Pambudi",
		Nick: "Ridho",
		Age:  17,
	}

	fmt.Println(a)
}
