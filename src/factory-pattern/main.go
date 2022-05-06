package main

import (
	"fmt"
	"go-DesignPattern/src/factory-pattern/factory"
	"go-DesignPattern/src/factory-pattern/general-ipml"
)

func main() {
	test := general_ipml.NewAwardRes()
	test.TestPrizeController()
	fmt.Println("-------------------------------")
	factory.Test()
}
