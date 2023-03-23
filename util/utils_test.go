package util

import (
	"fmt"
	"testing"
)

func TestDivideWords(t *testing.T) {
	str := "Hello world 啊哈"

	str = DividePreprocessedWords(str)

	fmt.Println(str)
}

func TestGenerateRedemptionCode(t *testing.T) {
	code := GenerateRedemptionCode(8, "")
	fmt.Println(code)
}
