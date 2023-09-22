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

func TestFlattenBsonDoc(t *testing.T) {
	type Test struct {
		Level1 struct {
			Level2 struct {
				Level3 string
			}
		}
	}

	test := Test{
		Level1: struct {
			Level2 struct {
				Level3 string
			}
		}{
			Level2: struct {
				Level3 string
			}{
				Level3: "hello",
			},
		},
	}

	doc := StructToBsonDoc(test)
	flattened := FlattenBsonDoc(doc)

	t.Log(flattened)
}

// func TestFlattenBsonD(t *testing.T) {
// 	type Test struct {
// 		Level1 struct {
// 			Level2 struct {
// 				Level3 string
// 			}
// 		}
// 		L1 struct {
// 			L2 struct {
// 				L3 string
// 			}
// 		}
// 	}

// 	test := Test{
// 		Level1: struct {
// 			Level2 struct {
// 				Level3 string
// 			}
// 		}{
// 			Level2: struct {
// 				Level3 string
// 			}{
// 				Level3: "hello",
// 			},
// 		},
// 		L1: struct {
// 			L2 struct {
// 				L3 string
// 			}
// 		}{
// 			L2: struct {
// 				L3 string
// 			}{
// 				L3: "world",
// 			},
// 		},
// 	}

// 	doc := StructToBsonD(test)
// 	flattened := FlattenBsonD(doc)

// 	t.Log(flattened)
// }

func TestStructToDoc(t *testing.T) {
	res := StructToBsonDoc(nil)

	t.Log(res)
}
