package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите арифметическое выражение")
	arifmetic, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(arifmetic)
	num1, num2 := getDigit(arifmetic)
	fmt.Println(getSign(arifmetic))
	fmt.Println(calculator(num1, num2, getSign(arifmetic)))

}
