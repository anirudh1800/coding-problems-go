package main

import "fmt"

func main() {
	permutations("abcd", 0)
}

func permutations(n string, start int) {

	if start >= len(n) {
		fmt.Println(n)
	}

	for i:= start; i < len(n); i++ {
		n = swap(n, start, i)
		permutations(n, start + 1)
		n = swap(n, start, i)
	}

}

func swap(n string, i int, j int) string {
	r := []rune(n)
	tmp := r[i]
	r[i] = r[j]
	r[j] = tmp
	return string(r)
}


