package main

import (
	"fmt"
	"strconv"
	"strings"
	"flag"
	"time"
)

func main() {
	var color string
	var toDis string
	timer := flag.Int("timer", 0, "an int")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		toDis = strings.Join(args, " ")
	} else {
		toDis = "yes"
	}
	for i := 0; i <= 256; i++ {
		color = "\x1b[38;5;" + strconv.Itoa(i) + "m"
		time.Sleep(time.Duration(*timer) * time.Millisecond)
		fmt.Print(color + toDis + "\x1b[0m")
          if i == 256{
               i = -1
          }
	}
}
