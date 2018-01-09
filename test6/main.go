package main

import (
  "fmt"
  "strconv"
)

func main() {
  	val, err := strconv.Atoi("-1")
	if err == nil && val < 0 {
		//err := fmt.Errorf("negative value")
	}
	fmt.Println(err)
}