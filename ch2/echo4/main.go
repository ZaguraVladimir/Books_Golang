package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	n   = flag.Bool("n", false, "Пропуск символа новой строки")
	sep = flag.String("s", " ", "Разделитель")
)

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
