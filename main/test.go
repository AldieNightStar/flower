package main

import (
	"fmt"

	"github.com/AldieNightStar/flower"
)

func main() {
	env := flower.NewEnv(nil)
	result, err := flower.EvalFile(env, "app.flower")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("DONE: ", result)
}
