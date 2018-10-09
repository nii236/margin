package main

import (
	"fmt"

	"github.com/nii236/margin/pkg/client"
)

func main() {
	c := client.New("http", "localhost", "8080", "v1")
	positions, err := c.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, pos := range positions.Positions {
		fmt.Println(pos.User, pos.Pair)

	}
}
