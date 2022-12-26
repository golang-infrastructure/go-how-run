package main

import (
	"fmt"
	how_run "github.com/golang-infrastructure/go-how-run"
)

func main() {
	runType, err := how_run.GetSourceCodeRunType()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(runType)
}
