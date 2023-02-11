package gobasic

import "fmt"

func LoopingDemo()  {

	names := SliceRun()

	for _, name := range names {
		fmt.Println(name)
	}

}