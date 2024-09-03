package main

import (
	"fmt"
	"math/big"
	"src/practice"
)

// globle variable
var d int = 234139;

var float float32;

func main() {


	bigNumber := new(big.Int)
	bigNumber.SetString("349562384598888844444444777777777777777777",10)

	//local veriable
	var l string= "10"
	//local varible with dynamic type
	x := 1034;
	y := "go";

	fmt.Println("hello world!!",l);
	fmt.Println(y);
	fmt.Printf("%T\n", bigNumber);
	fmt.Printf("data type --> %T\n",x);
	fmt.Println(" big number ----->>",bigNumber);
	fmt.Println(float);

	fmt.Println("Max number : ", practice.MaxValue(10,20));

	a, b := practice.Swap("go", "java");
	fmt.Println("after swap : ", a,b);

	practice.Day2();
}



