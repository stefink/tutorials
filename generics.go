package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func Square[V Number](value V) V {
	return value * value
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  312.22,
		"second": 22.2,
	}
	fmt.Println(SumIntsOrFloats[string, int64](ints))
	fmt.Println(SumIntsOrFloats[string, float64](floats))
	fmt.Println(Square[float64](9.6))
	fmt.Println(Square[int64](6))
}
