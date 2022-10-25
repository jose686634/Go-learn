package main

import (
	"fmt"

	"github.com/Go-learn/types/StringHelper"
)

func main() {

	// basic test for getting the length of a string
	var str string = "hi my name is Jose"
	fmt.Println(StringHelper.GetStringlen(str))

	/*
	Experiment on suffix array standard library in golang
	Given a list of word, find the word that have following substring in it/
	*/
	words := []string{
		"banana",
		"apple",
		"pear",
		"tangerine",
		"orange",
		"lemon",
		"peach",
		"persimmon",
	}
	fmt.Println(StringHelper.RecommendMatch(&words, "ang")) // returns ["tangerine","orange"]

}
