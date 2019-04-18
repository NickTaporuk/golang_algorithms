package main

type SSS struct {}

func main() {
	var x SSS
	_ = *ref(x)
}

func ref(z SSS) *SSS {
	return &z
}
