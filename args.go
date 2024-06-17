package main

import "os"

func GetArg(index int) string {
	for i, v := range os.Args {
		if i < index {
			continue
		}
		return v
	}
	return ""
}