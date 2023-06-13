package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}
