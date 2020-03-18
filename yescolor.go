package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var color string
	var toDis string
	args := os.Args[1:]
	if len(args) > 0 {
		toDis = strings.Join(args, " ")
	} else {
		toDis = "yes"
	}
	for i := 0; i <= 256; i++ {
		color = "\x1b[38;5;" + strconv.Itoa(i) + "m"
		fmt.Print(color + toDis + "\x1b[0m")
          if i == 256{
               i = -1
          }
	}
}
