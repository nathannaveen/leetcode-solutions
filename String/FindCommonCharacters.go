package main

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"headedfdhhbbdghdhbge", "ebecdbgggcbhgecbbdeb", "cahagfhbdedcbgdefcdb", "bahacadbfacgfcfgeehg", "edddaeehfaceecabbdfh", "gaheghhfhfbbfghcedea", "fefcbbdaagadeghggccg", "bceddhfhcchhfghbedda", "ccgbgeafbhhgfcegfagg", "cdabfaahdacefafdfeac", "bfdbhcaaddghbgbfgghg", "cdaachdffcgcgggfdfbc", "bbeceagaedbdchdebfdc", "bcdbafaffhbhedhdhbeb", "gfgadacchggbdehecgeb", "gecbabaahcgefebcabgc", "dedfeeagcgdbeahcahaa", "fdcdbfddgegedcffahda", "eacfgadacfhhhbdghcac", "echccdgccdaechbhafdg"}))
}

func commonChars(A []string) []string {
	res := []string{}
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	//              a  b  c  d  e   f   g   h   i   j   k   l   m   n   o   p   q   r   s   t   u   v   w   x   y   z
	g := int64(1)

	for _, s := range A[0] {
		g *= int64(primes[s-'a'])
	}

	for i := 1; i < len(A)-1; i++ {
		h := int64(1)
		for _, i2 := range A[i] {
			if g%int64(primes[i2-'a']) == 0 {
				h *= int64(primes[i2-'a'])
				g /= int64(primes[i2-'a'])
			}
		}
		g = h
	}

	for _, i := range A[len(A)-1] {
		if g%int64(primes[i-'a']) == 0 {
			res = append(res, string(i))
			g /= int64(primes[i-'a'])
		}
	}

	return res
}
