package StringHelper

import (
	"index/suffixarray"
	"strings"
)

func RecommendMatch(strSlice *[]string, searchWord string) []string {
	/*
		Description:
			Given a list of strings strSlice, find the words that have a substring as searchWord
		Parameter:
			strSlice: A slice of string that need to be searched to find a match
			searchWord: A string which need to be searched
		Return:
			A slice of string having a match

		Working:
			A single byte slice is created from joining all the strings seperated by /x00
			a suffix array is created for that.searchWord
			suffix array is lookedup by the pattern and corresponding starting index are return by it
			These index are further traversed in both direction to obtain a /x00
			the string  is then appended to a string slice and returned
	*/
	data := []byte("\x00" + strings.Join(*strSlice, "\x00") + "\x00")
	sa := suffixarray.New(data)                  // creates a typed/struct named Index and returns a pointer of it.
	indices := sa.Lookup([]byte(searchWord), -1) // Lookup is a mthod of Index type/struct that return a []int of indexes

	var strResult []string
	for _, ind := range indices {
		var start, end int
		for i := ind - 1; i >= 0; i-- {
			if data[i] == 0 {
				start = i + 1
				break
			}
		}
		for i := ind + 1; i <= len(data); i++ {
			if data[i] == 0 {
				end = i
				break
			}
		}
		strResult = append(strResult, string(data[start:end]))
	}
	return strResult
}
