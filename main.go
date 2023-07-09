package main

import (
	"fmt"

	"github.com/coderx31/puppy"
)

type ByteSize int

func main() {
	fmt.Println("learn how to code go - udemy course")

	// 13. variables, zero values and blank identifier

	/*
		var h int = 44
		fmt.Println(h)

		a := 42
		fmt.Println(a)

		b, c, d, _, f := 0, 1, 2, 3, "happiness"

		fmt.Println(b, c, d, f)

		var g int
		fmt.Println(g)
		g = 23
		fmt.Println(g)
	*/

	// 14. Using printf for decimal and hexadecimal values

	/*
		adams := 42

		fmt.Printf("42 as binary %b \n", adams)
		fmt.Printf("42 as hexadecimal %x \n", adams)

		// print these values as both binary and hexadecimal
		a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

		fmt.Printf("a as binary %b \n", a)
		fmt.Printf("a as hexadecimal %x \n", a)
		fmt.Printf("b as binary %b \n", b)
		fmt.Printf("b as hexadecimal %x \n", b)
		fmt.Printf("c as binary %b \n", c)
		fmt.Printf("c as hexadecimal %x \n", c)
		fmt.Printf("d as binary %b \n", d)
		fmt.Printf("d as hexadecimal %x \n", d)
		fmt.Printf("e as binary %b \n", e)
		fmt.Printf("e as hexadecimal %x \n", e)
		fmt.Printf("f as binary %b \n", f)
		fmt.Printf("f as hexadecimal %x \n", f)

	*/

	// 16. values, types, conversions, scope and housekeeping

	// 25. hands on exercise
	//printMeasures()
	//printMeasureSolution()

	// dependency management
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

}

func printMeasures() {
	const (
		_ = iota
		kb
		mb
		gb
		tb
		pb
		eb
	)

	fmt.Println("KB ->", 1024<<kb)
	fmt.Println("MB ->", 1024<<mb)
	fmt.Println("GB ->", 1024<<gb)
	fmt.Println("TB ->", 1024<<tb)
	fmt.Println("pb ->", 1024<<pb)
	fmt.Println("eb ->", 1024<<eb)
}

func printMeasureSolution() {
	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	fmt.Printf("%d \t\t\t %b \n", KB, KB)
	fmt.Printf("%d \t\t\t %b \n", MB, MB)
	fmt.Printf("%d \t\t\t %b \n", GB, GB)
	fmt.Printf("%d \t\t\t %b \n", TB, TB)
	fmt.Printf("%d \t\t\t %b \n", PB, PB)
	fmt.Printf("%d \t\t\t %b \n", EB, EB)
}
