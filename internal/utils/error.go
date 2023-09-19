package util

import "fmt"

func HasError(err error) bool {
	if err != nil {
		LogInput(fmt.Sprint("has error:", err))
	}

	return err != nil
}

func LogInput(str string) {
	fmt.Printf("\n\nerror:\n %s \n\n", str)
}
