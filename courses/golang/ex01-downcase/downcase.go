package main

func Downcase(str string) (answ string, err error) {
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			char += 32
		}
		answ += string(char)
	} 
	return answ, nil
}

func main() {
	str := "Hello World!"
	print(Downcase(str))
}