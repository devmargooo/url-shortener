package main

import (
	"fmt"

	"github.com/devmargooo/url-shortener/internal/config"
)

func main() {
    cfg := config.MustLoad()
	fmt.Printf("%+v", cfg);
}