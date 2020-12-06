package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strconv"
)

func calc(){
	fmt.Println("to calculate something choose an operand. ex (*, /, -, +)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	o := scanner.Text()
	fmt.Printf("enter first number: ")
	scanner.Scan()
	fn, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("enter second number: ")
	scanner.Scan()
	sn, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	switch {
	case o == "+":
		fmt.Printf("%v + %v = %v \n", fn, sn, fn+sn)
	case o == "-":
		fmt.Printf("%v - %v = %v \n", fn, sn, fn-sn)
	case o == "/":
		fmt.Printf("%v / %v = %v \n", fn, sn, fn/sn)
	case o == "*":
		fmt.Printf("%v * %v = %v \n", fn, sn, fn*sn)
	}
}

func main(){
	calc()
}
