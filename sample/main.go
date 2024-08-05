package main

import (
	"fmt"
	"strconv"

	"github.com/randyardiansyah25/gerr"
)

func main() {
	fmt.Println(FirstLevel())
}

func FirstLevel() error {
	er := SecondLevel()
	return gerr.AddNew(er, "error from first level")
}

func SecondLevel() error {
	er := ThirdLevel()
	return gerr.AddNew(er, "error from second level")
}

func ThirdLevel() error {
	a := "a" //
	_, er := strconv.Atoi(a)
	return gerr.AddNew(er, fmt.Sprintf("error convert %s to integer", a))
}
