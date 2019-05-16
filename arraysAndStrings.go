package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("-Print Evens-")
	printEvens()
	fmt.Println("-Get century from year-")
	fmt.Printf("%d \n", centuryFromYear(1905))
	fmt.Println("-Reverse String-")
	fmt.Printf("%s \n", reverse("Hola"))
	fmt.Println("-Find first not repeating char-")
	fmt.Printf("%s \n", firstNotRepeatingCharacterLucas("abcdefghijklmnopqrstuvwxyziflskecznslkjfabe"))
	fmt.Println("-Rotate Matrix-")

	rotateImage(a)
	rotateImage(b)
	rotateImage(c)
	rotateImage(d)

	fmt.Println("- First Duplicate -")
	fmt.Printf("%d \n", firstDuplicate(smallArray))
}


func printEvens() {
	for i := 0; i < 10 ; i++  {
		if i % 2 == 0 {
			//s := fmt.Sprintf("%d is even", i)
			//fmt.Println(s)
			fmt.Printf("%d is even\n", i)
		}
	}
}

func centuryFromYear(year int) int {
		var century float64 = (float64(year)/ 100)
		return int(math.Ceil(century))
}

func reverse(s string) (result string) {
	for _,v := range s {
		fmt.Print("\n", string(v))
		result = string(v) + result
	}
	return
}

func checkPalindrome(inputString string) bool {
	return inputString == reverse(inputString)
}

func firstDuplicate(a []int) int {
	m := make(map[int]int)
	for _, value := range a {
		m[value] += 1
		if m[value] == 2 {
			return  value

		}
	}
	return -1
}

func firstNotRepeatingCharacterLucas(s string) string {
	m := make (map[rune]int)
	a := []rune{}

	for _,v := range s  {
		if _, ok := m[v]; !ok {
			a = append(a, v) //appends rune v to array a
		}
		m[v] += 1
	}

	for _ ,value := range a {
		if m[value] == 1 { return string(value) } //needs to be casted to string because doing range on an string returns runes
	}

	return "_"
}

func rotateImage(a [][]int) [][]int {
	fmt.Printf("Old array %d \n", a)
	fmt.Printf("length %d \n", len(a))
	if len(a) == 1 { return a }

	lenA := len(a)
	var rotatedArray = make([][]int, lenA)
	for i := range rotatedArray {
		rotatedArray[i] = make([]int, lenA)
	}

	for index, subArray := range a {
		fmt.Printf("Rotated array %d \n", rotatedArray)
		fmt.Printf("Array length in index %v is %d and content %d \n", index, len(a), subArray)
		for subIndex, elementValue := range subArray {
			fmt.Printf("Index: %v Subindex: %v Value: %d \n", index, subIndex, elementValue)
			fmt.Printf("?? %d", len(a) - (index + 1))
			rotatedArray[subIndex][len(a) - (index + 1)] = a[index][subIndex]
		}
	}
	fmt.Printf("Old array %d \n", a)
	fmt.Printf("Rotated array %d \n", rotatedArray)
	return  rotatedArray
}


func rotateArray(a [][]int) [][]int {
	if len(a) == 1 { return a }

	lenA := len(a)
	var rotatedArray = make([][]int, lenA)
	for i := range rotatedArray {
		rotatedArray[i] = make([]int, lenA)
	}

	for index, subArray := range a {
		for subIndex, _ := range subArray {
			rotatedArray[subIndex][len(a) - (index + 1)] = a[index][subIndex]
		}
	}
	return  rotatedArray
}