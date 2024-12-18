package util

import "testing"

func Test(t *testing.T) {
	tokenizer := NewStringTokenizer3("this is a test", " ", false)

	for tokenizer.HasMoreTokens() {
		token, err := tokenizer.NextToken()
		if err != nil {
			panic(err)
		}
		println(token)
	}

	println("Remaining tokens:", tokenizer.CountTokens())
}
