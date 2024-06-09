package main

import "os"

// Checks if the passed in string was passed as an arg
// index checks if the arg is only at that index
// if index < 0, every arg is checked
// Returns the index the string was found in os.Args
// Returns -1 if the string was not found in os.Args
func CheckArg(arg string, index int) int {
	if index < 0 {
		args := os.Args
		for i, v := range args {
			if v == arg {
				return i
			}
		}
		return -1
	} else {
		if arg == os.Args[index] {
			return index
		} else {
			return -1
		}
	}
}
