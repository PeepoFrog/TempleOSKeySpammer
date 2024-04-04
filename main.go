package main

import (
	"fmt"
	"os/exec"

	"github.com/google/shlex"
)

const virtualBoxMachineName string = "templeOS"

func main() {

	randomWordKey, err := shlex.Split(fmt.Sprintf("controlvm %v keyboardputscancode 41 c1", virtualBoxMachineName))
	if err != nil {
		panic(err)
	}
	enterKey, err := shlex.Split(fmt.Sprintf("controlvm %v keyboardputscancode 1c 9c", virtualBoxMachineName))
	if err != nil {
		panic(err)
	}

	for range 100 {
		for range 7 {
			_, err = exec.Command("VBoxManage", randomWordKey...).Output()
			if err != nil {
				panic(err)
			}
		}

		_, err = exec.Command("VBoxManage", enterKey...).Output()
		if err != nil {
			panic(err)

		}
	}

}
