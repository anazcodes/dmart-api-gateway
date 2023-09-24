package util

import (
	"fmt"
)

func HasError(err error) bool {
	if err != nil {
		Logger("has error:", err)

	}
	return err != nil
}

// Logger can print any type of data
func Logger(any ...any) {
	fmt.Println("\n" + fmt.Sprint(any...) + "\n")
}
