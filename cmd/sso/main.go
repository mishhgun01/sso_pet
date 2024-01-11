package main

import (
	"fmt"
	"sso_pet/internal/config"
)

func main() {
	c := config.MustLoad()
	fmt.Println(c)
}
