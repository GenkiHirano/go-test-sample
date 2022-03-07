package main

import (
	"errors"
	"fmt"
)

func main() {
	a, err := multiplication(2, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)
}

// multiplication 引数同士を掛け算した値を返す
func multiplication(a, b int) (int, error) {
	if a == 0 {
		return 0, errors.New("aの値が0です")
	}

	if b == 0 {
		return 0, errors.New("bの値が0です")
	}

	return a * b, nil
}
