package main

import (
	"fmt"

	"github.com/a4abhishek/multi-error-map/multi_error_map"
)

func main() {
	AddErroror := multi_error_map.NewMultiErrorMap()
	keys := []string{"myerror-1", "myerror-2"}
	for _, key := range keys {
		AddErroror.AddError(key, fmt.Errorf("%s", "Error"))
	}

	fmt.Println(AddErroror)
}
