package main

import "log"

func main() {
	log.Println("start")

	// fibnormalrectest()
	// fibfuncgenTest()
}

func fibnormalrectest() {
	for i := 0; i < 50; i++ {
		log.Println(i, fibNormalRecursion(i))
	}
}

func fibfuncgenTest() {
	fibA := fibFuncGen()
	fibB := fibFuncGen()

	count := 70
	for i := 50; i <= count; i++ {
		log.Println(fibA(i))
	}
	log.Println(fibA(50), fibB(50))
	log.Println(fibA(51), fibB(51))
	log.Println(fibA(52), fibB(52))
	log.Println(fibA(53), fibB(53))
}
