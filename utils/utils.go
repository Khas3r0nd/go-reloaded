package utils

import (
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Конвертация строки представляющяя собой двоичное число в десятичную систему исчисления
func BinaryToDecimal(s string) string {
	for _, value := range s {
		if value != '1' && value != '0' {
			return ""
		}
	}
	binaryNum, _ := strconv.Atoi(s)
	var remainder int
	index := 0
	decimalNum := 0
	for binaryNum != 0 {
		remainder = binaryNum % 10
		binaryNum = binaryNum / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}

	return strconv.Itoa(decimalNum)
}

// Конвертация строки представляющюю собой шестандцатиричное число в десятичную систему исчисления
func HexToDecimal(s string) string {
	hexChars := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	formatedS := ToLower(s)

	for _, value := range formatedS {
		if !contains(hexChars, string(value)) {
			return ""
		}
	}
	decimalNum := 0
	for _, value := range s {
		var digit int
		if unicode.IsDigit(value) {
			digit = int(value - '0')
		} else {
			digit = int(unicode.ToLower(value) - 'a' + 10)
		}
		decimalNum = decimalNum*16 + digit
	}

	return strconv.Itoa(decimalNum)
}

// Функция которая возвращает строку в верхнем регистре
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Функция которая возвращает строку в нижнем регистре
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Функция которая возвращает строку где строка начинается с большой буквы, остальные буквы в нижнем регистре
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	first := strings.ToUpper(s[:1])
	if len(s) == 1 {
		return first
	}

	return first + strings.ToLower(s[1:])
}

// Функция которая проверяет является ли строка словом
func isWordOrNumber(s string) bool {
	for i, value := range s {
		if !(unicode.IsLetter(value) || unicode.IsDigit(value)) && i == 0 {
			return false
		}
	}
	return true
}

// Вспомогательная функция, которая выполняет нужные модификации, как параметр передается функция
func modify(ptrToAnswer *[]string, i int, transform func(string) string, intValue int) {
	switch {
	case intValue == 0:
		if transform((*ptrToAnswer)[len(*ptrToAnswer)-1]) == "" {
			*ptrToAnswer = append((*ptrToAnswer)[:i-1], (*ptrToAnswer)[i:]...)
			log.Printf("Invalid input, skipping the modification")
		} else {
			(*ptrToAnswer)[len(*ptrToAnswer)-1] = transform((*ptrToAnswer)[len(*ptrToAnswer)-1])
		}
	// Если у нас после модификатора есть цифра, то проходимся циклом по предыдущим элементам
	case intValue > 0:
		for j := 1; j <= intValue; j++ {
			// Если значение будетя являтся пунктуацией или группой пунктуаций, то нужно пропустить этот элемент и увеличить intValue на один
			if isWordOrNumber((*ptrToAnswer)[len(*ptrToAnswer)-j]) {
				(*ptrToAnswer)[len(*ptrToAnswer)-j] = transform((*ptrToAnswer)[len(*ptrToAnswer)-j])
			} else if len((*ptrToAnswer)) > intValue {
				intValue += 1
			}
		}
	}
}

func CheckModifier(ptrToSlice, ptrToAnswer *[]string, currentElement string, i int) bool {
	formattedCurrentElement := ToLower(currentElement)
	switch {
	case formattedCurrentElement == "(bin)" && i > 0:
		modify(ptrToAnswer, i, BinaryToDecimal, 0)
	case formattedCurrentElement == "(hex)" && i > 0:
		modify(ptrToAnswer, i, HexToDecimal, 0)
	case formattedCurrentElement == "(low)" && i > 0:
		modify(ptrToAnswer, i, ToLower, 0)
	case formattedCurrentElement == "(up)" && i > 0:
		modify(ptrToAnswer, i, ToUpper, 0)
	case formattedCurrentElement == "(cap)" && i > 0:
		modify(ptrToAnswer, i, Capitalize, 0)
	case formattedCurrentElement == "(up," && i > 0:
		intValue, err := strconv.Atoi(string((*ptrToSlice)[i+1][:len((*ptrToSlice)[i+1])-1]))
		if err != nil || intValue <= 0 {
			log.Printf("Invalid integer value in %s%v) skipping modification", currentElement, intValue)
			return true
		}
		if len((*ptrToAnswer)) < intValue {
			intValue = len(*ptrToAnswer)
		}
		modify(ptrToAnswer, i, ToUpper, intValue)
		return true
	case formattedCurrentElement == "(low," && i > 0:
		intValue, err := strconv.Atoi(string((*ptrToSlice)[i+1][:len((*ptrToSlice)[i+1])-1]))
		if err != nil || intValue <= 0 {
			log.Printf("Invalid integer value in %s%v) skipping modification", currentElement, intValue)
			return true
		}
		if len(*ptrToAnswer) < intValue {
			intValue = len(*ptrToAnswer)
		}
		modify(ptrToAnswer, i, ToLower, intValue)
		return true
	case formattedCurrentElement == "(cap," && i > 0:
		intValue, err := strconv.Atoi(string((*ptrToSlice)[i+1][:len((*ptrToSlice)[i+1])-1]))
		if err != nil || intValue <= 0 {
			log.Printf("Invalid integer value in %s%s) skipping modification", currentElement, (*ptrToSlice)[i+1])
			return true
		}
		if len(*ptrToAnswer) < intValue {
			intValue = len(*ptrToAnswer)
		}
		modify(ptrToAnswer, i, Capitalize, intValue)
		return true
	default:
		*ptrToAnswer = append(*ptrToAnswer, currentElement)
		return false
	}
	return false
}

func CheckPunctuation(ptrToAnswer, finalAnswer *[]string, i int) {
	punctChars := []string{".", ",", "!", "?", ":", ";"}
	str := (*ptrToAnswer)[i]
	punct := ""
	for _, c := range str {
		if contains(punctChars, string(c)) {
			punct += string(c)
		} else {
			break
		}
	}
	if punct != "" && i > 0 {
		// Добавить пунктуацию к последнему слову в слайсе finalAnswer
		(*finalAnswer)[len(*finalAnswer)-1] += punct
		// Добавить оставшийся текст без знаков препинания к текущему слову в слайсе finalAnswer
		if len(*finalAnswer) <= i {
			if str[len(punct):] != "" {
				(*finalAnswer)[len(*finalAnswer)-1] += " "
			}
			(*finalAnswer)[len(*finalAnswer)-1] += str[len(punct):]
		} else {
			(*finalAnswer)[i] = str[len(punct):]
		}
	} else {
		// Если пунктуация не найдена, скопировать слово как есть в слайс finalAnswer
		if len(*finalAnswer) <= i {
			*finalAnswer = append(*finalAnswer, str)
		} else {
			(*finalAnswer)[i] = str
		}
	}
	if i+1 < len(*ptrToAnswer) && startsWithVowelSound((*ptrToAnswer)[i+1]) && ToLower((*ptrToAnswer)[i]) == "a" {
		if len(*finalAnswer) > 2 {
			(*finalAnswer)[len(*finalAnswer)-1] = "an"
		} else {
			(*finalAnswer)[len(*finalAnswer)-1] = "An"
		}
	} else if i+1 < len(*ptrToAnswer) && !startsWithVowelSound((*ptrToAnswer)[i+1]) && ToLower((*ptrToAnswer)[i]) == "an" {
		if len(*finalAnswer) > 2 {
			(*finalAnswer)[len(*finalAnswer)-1] = "a"
		} else {
			(*finalAnswer)[len(*finalAnswer)-1] = "A"
		}
	}
}

func CheckQuotes(s *[]string) []string {
	newSlice := []string{}
	flagSingle := false
	flagDouble := false
	for i := 0; i < len(*s); i++ {
		if (*s)[i][0] == '\'' {
			if len((*s)[i]) > 1 && !flagSingle {
				newSlice = append(newSlice, (*s)[i])
				continue
			}
			if !flagSingle {
				// Открывающая одинарная кавычка, должна добавиться к началу следующего элемента
				if i < len(*s)-1 {
					newSlice = append(newSlice, "'"+(*s)[i+1])
					i++
				} else {
					newSlice = append(newSlice, "'")
				}
			} else {
				// Закрывающая одинарная кавычка, должна добавиться в конец предыдущего элемента
				newSlice[len(newSlice)-1] += (*s)[i]
			}
			flagSingle = !flagSingle
		} else if (*s)[i][0] == '"' {
			if len((*s)[i]) > 1 && !flagDouble {
				newSlice = append(newSlice, (*s)[i])
				continue
			}
			if !flagDouble {
				// Открывающая двойная кавычка, должна добавиться к началу следующего элемента
				if i < len(*s)-1 {
					newSlice = append(newSlice, "\""+(*s)[i+1])
					i++
				} else {
					newSlice = append(newSlice, "\"")
				}
			} else {
				// Закрывающая двойная кавычка, должна добавиться в конец предыдущего элемента
				newSlice[len(newSlice)-1] += (*s)[i]
			}
			flagDouble = !flagDouble
		} else {
			newSlice = append(newSlice, (*s)[i])
		}
	}
	return newSlice
}

func AppendQuotes(s *[]string) []string {
	newSlice := []string{}
	flag := false
	for i := 0; i < len(*s); i++ {
		if (*s)[i][0] == '\'' && !flag && len((*s)[i]) > 1 {
			flag = !flag
			newSlice = append(newSlice, (*s)[i])
		} else if (*s)[i][0] == '\'' && flag {
			newSlice[len(newSlice)-1] += (*s)[i]
			flag = !flag
		} else {
			newSlice = append(newSlice, (*s)[i])
		}
	}
	return newSlice
}

func CheckSpecialCase(s []string) []string {
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == "'" && s[i-1][len(s[i-1])-1] == 'n' && i+1 < len(s) && s[i+1] == "t" {
			s[i-1] = s[i-1] + s[i] + s[i+1]
			s = append(s[:i], s[i+2:]...)
		}
	}
	return s
}

// Простая функция contains, которая возвращает истину если строка имеется в слайсе строк
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Простая функция, которая возвращает истину если строка начинается на гласную букву.
func startsWithVowelSound(str string) bool {
	vowelSounds := []string{"a", "e", "i", "o", "u", "h"}
	firstChar := string([]rune(str)[0])
	for _, s := range vowelSounds {
		if s == firstChar {
			return true
		}
	}
	return false
}
