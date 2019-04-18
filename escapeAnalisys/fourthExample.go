package main

type SSSS struct {
	M *int
}

func main() {
	var i int
	refStruct(i)
}

func refStruct(y int) (z SSSS) {
	z.M = &y
	return z
}
