package main

// generating function that spawns out fib

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
