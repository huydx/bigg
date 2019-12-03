package main

import (
	"fmt"
	"strconv"
	"strings"
)

// opcode
const (
	add  = 1
	mult = 2
	halt = 99
)

func main() {
	program := make([]int, 0)
	for _, s := range strings.Split(input, ",") {
		if s == "" {
			continue
		}
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		program = append(program, int(i))
	}

	i := 0
	l := len(program)
	fmt.Println(l)
	fmt.Println(program)
Loop:
	for {
		assert(l, i)
		opcode := program[i]
		if opcode == add {
			assert(l, i+3)
			v1 := program[i+1]
			v2 := program[i+2]
			o1 := program[i+3]
			assert(l, v1, v2, o1)

			i1 := program[v1]
			i2 := program[v2]
			program[o1] = i1 + i2
			i += 4
		} else if opcode == mult {
			assert(l, i+3)
			v1 := program[i+1]
			v2 := program[i+2]
			o1 := program[i+3]
			assert(l, v1, v2, o1)

			i1 := program[v1]
			i2 := program[v2]
			program[o1] = i1 * i2
			i += 4
		} else if opcode == halt {
			break Loop
		}
	}

	fmt.Println(program[0])
}

func assert(len int, indices ...int) {
	for _, index := range indices {
		if index >= len {
			panic(fmt.Sprintf("overflow len %v indices %v", len, indices))
		}
	}
}

var input = `1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,10,19,23,1,6,23,27,1,5,27,31,1,10,31,35,2,10,35,39,1,39,5,43,2,43,6,47,2,9,47,51,1,51,5,55,1,5,55,59,2,10,59,63,1,5,63,67,1,67,10,71,2,6,71,75,2,6,75,79,1,5,79,83,2,6,83,87,2,13,87,91,1,91,6,95,2,13,95,99,1,99,5,103,2,103,10,107,1,9,107,111,1,111,6,115,1,115,2,119,1,119,10,0,99,2,14,0,0`
