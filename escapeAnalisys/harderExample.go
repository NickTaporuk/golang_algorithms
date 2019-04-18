package main

type SS struct {}

func main() {
	var x SS
	y := &x
	_ = *identityHard(y)
}

func identityHard(z *SS) *SS {
	return z
}
