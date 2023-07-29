package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/Script-MongoDB-sliceString/controllers"
)

func main() {

	err := controllers.UpdateUserData()
	if err != nil {
		fmt.Errorf("error")
	}
}
