package main

import "log"

func main() {
	log.Println("start")
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

func fibFuncGen() func(int) int {
	// cache
	var fibMap = make(map[int]int)

	// initialize it first so that
	// we can call it from the inside
	var fibFunc func(v int) int

	// this is to set the value of fibFunc
	fibFunc = func(v int) int {
		if v <= 2 {
			return 1
		}

		// if using recursion
		// return fibFunc(v-1) + fibFunc(v-2)

		// if using cache
		// check if exist in map first
		p, q := fibMap[v-1], fibMap[v-2]
		if p == 0 || q == 0 {
			p = fibFunc(v - 1)
			q = fibFunc(v - 2)
		}
		fibMap[v] = p + q
		return fibMap[v]
	}

	return fibFunc
}
