package main

import (
	"fmt"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func isValid(str string) bool {
	var internalCounter, firstSize int = 0,0
	asArray := strings.Split(str, "")

	if len(asArray) == 0 || len(asArray) > 10000 {
		return false
	}

	if len(asArray) % 2 != 0 {
		return false
	}

	opposite := func(direct string, last int) string {
		if direct == "[" {
			return "]"
		}

		if direct == "]" {
				return "["
			}

		if direct == "(" {
			return ")"
		}

		if direct == ")" {
				return "("
			}

		if direct == "{" {
			return "}"
		}

		if direct == "}" {
				return "{"
			}

		return "nil"
	}

	backSymbol := func(symbol string) bool {
		switch symbol {
		case ")":
			return false
		case "}":
			return false
		case "]":
			return false
		default:
			return true
		}
	}

	//checkSymbol := func(symbols []string) bool {
	//	for k := 0; k < len(symbols); k++ {
	//		if symbols[k] != "[" {
	//			return false
	//		} else {
	//			if symbols[k] != "]" {
	//				return false
	//			} else {
	//				if symbols[k] != "(" {
	//					return false
	//				} else {
	//					if symbols[k] != ")" {
	//						return false
	//					} else {
	//						if symbols[k] != "{" {
	//							return false
	//						} else {
	//							if symbols[k] != "}" {
	//								return false
	//							}
	//						}
	//					}
	//				}
	//			}
	//		}
	//	}
	//
	//	return true
	//}
	flag := 1
	elements := asArray

	//if !checkSymbol(elements) {
	//	return false
	//}

	firstSize = len(elements)
	for flag > 0 {
		for i := 0; i < len(elements); i++ {
			if internalCounter > firstSize {
				return false
			}
			if len(elements) == 2 {
				if i+1 == len(elements) && elements[i] == opposite(elements[i-1], 1) {
					return true
				}
				internalCounter++
				if  len(elements) == 0 {
					flag = 0
					return true
				}
			}
			if i+1 == len(elements) {
				if elements[i-1] == opposite(elements[i], 1) {
					elements = elements[:len(elements)-1]
					elements = elements[:len(elements)-1]
				}
				internalCounter++
				if  len(elements) == 0 {
					flag = 0
					return true
				}

				if len(elements) > 0 {
					i = -1
				}
			} else {
				oppositeValue := opposite(elements[i+1], 0)
				if backSymbol(elements[i]) && elements[i] == oppositeValue {
					elements = remove(elements, i)
					elements = remove(elements, i)
					if len(elements) == 2 && !backSymbol(elements[0]) {
						return false
					}
					internalCounter++
					i = -1

				} else {

					if i-1 > 0 && elements[i] == opposite(elements[i-1], 0) {
						elements = remove(elements, i-1)
						elements = remove(elements, i-1)
						if len(elements) == 2 && !backSymbol(elements[0]) {
							return false
						}
						internalCounter++
						i = -1
					}

				}
			}

			if  len(elements) == 0 {
				flag = 0
			}
		}
	}

	return true
}

func main() {
	var testString = "[]({})[{(())}{}][{()}}{]"
	fmt.Println(testString)
	fmt.Println(isValid(testString))
}
