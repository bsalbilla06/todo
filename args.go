package main

import (
	"fmt"
	"os"
)

// Safer way to get args
func GetArg(index int) (string, error) {
	for i, v := range os.Args {
		if i < index {
			continue
		}
		return v, nil
	}
	return "", fmt.Errorf("argument index out of range [%v] with length %v", index, len(os.Args))
}