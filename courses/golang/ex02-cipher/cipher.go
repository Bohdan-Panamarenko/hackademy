package main

import (
	"fmt"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct {
	Code int
}

func NewCaesar() *Caesar {
	return &Caesar{Code: 3}
}

func NewShift(shift int) *Caesar {
	if shift < -25 || shift > 25 || shift == 0 {
		return nil
	}
	return &Caesar{Code: shift}
}

func (c Caesar) Encode(message string) (codedMessage string) {
	for _, value := range message {
		if value >= 'A' && value <= 'Z' {
			value += 32
		}
		if value >= 'a' && value <= 'z' {
			value += rune(c.Code)

			if value > 'z' {
				value -= 26
			}

			if value < 'a' {
				value += 26
			}

			codedMessage += string(value)
		}
	}

	return codedMessage
}

func (c Caesar) Decode(codedMessage string) (message string) {
	for _, value := range codedMessage {
		value -= rune(c.Code)

		if value > 'z' {
			value -= 26
		}

		if value < 'a' {
			value += 26
		}

		message += string(value)
	}

	return message
}

type Vigenere struct {
	Code string
}

func NewVigenere(code string) *Vigenere {
	isValid, onlyA := true, true
	for _, char := range code {
		if !(char >= 'a' && char <= 'z') {
			isValid = false
			break
		}
		if char != 'a' {
			onlyA = false
		}
	}

	if !isValid || onlyA || code == "" {
		return nil
	}

	return &Vigenere{Code: code}
}

func (v Vigenere) getCode(messageLen int) string {
	currCode := v.Code
	codeInitialLen := len(currCode)
	codeLen := codeInitialLen

	for messageLen > codeLen {
		currCode += v.Code
		codeLen += codeInitialLen
	}

	return currCode[:messageLen]
}

func removeSymbols(str string) string {
	strWithoutSymbols := ""
	for _, value := range str {
		if value >= 'A' && value <= 'Z' || value >= 'a' && value <= 'z' {
			strWithoutSymbols += string(value)
		}
	}

	return strWithoutSymbols
}

func (v Vigenere) Encode(message string) (codedMessage string) {
	message = removeSymbols(message)
	currCode := v.getCode(len(message))

	for index, value := range message {
		if value >= 'A' && value <= 'Z' {
			value += rune(32)
		}
		if value >= 'a' && value <= 'z' {
			value += rune(currCode[index] - 'a')

			if value > 'z' {
				value -= rune(26)
			}

			codedMessage += string(value)
		}
	}

	return codedMessage
}

func (v Vigenere) Decode(codedMessage string) (message string) {
	currCode := v.getCode(len(codedMessage))
	for index, value := range codedMessage {
		value -= rune(currCode[index] - 'a')

		if value < 'a' {
			value += rune(26)
		}

		message += string(value)
	}

	return message
}

func main() {
	// c := Vigenere{ Code: "lemon" }
	// a := "Hello world!"
	// b := c.Encode(a)
	// fmt.Printf("%q | %q | %q\n", a, b, c.Decode(b))

	// a := Vigenere{ Code: "lemon" }
	// fmt.Printf("%q | %q\n", a, a.getCode(7))

	// fmt.Printf("%q\n", generateCode())

	a := NewVigenere("lemon")
	str := "Hello world!"
	b := a.Encode(str)
	fmt.Printf("%q | %q | %q\n", str, b, a.Decode(b))
}
