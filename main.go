package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(a, b string) string {
	result := strings.ReplaceAll(a+b, "\"", "")
	return "\"" + result + "\""
}

func substraction(a, b string) string {
	split := strings.Split(a, b)

	newSplit := (strings.Join(split, ""))
	newSplit1 := strings.ReplaceAll(newSplit, "\"", "")
	return "\"" + newSplit1 + "\""
}

func multiplication(a, b string) string {
	int_val, _ := strconv.Atoi(b)
	copy := strings.Repeat(a, int_val)
	copy1 := strings.ReplaceAll(copy, "\"", "")
	return "\"" + copy1 + "\""
}

func division(a, b string) string {
	int_val1, _ := strconv.Atoi(b)
	length := len(a)
	res := a[:length/int_val1]
	res1 := strings.ReplaceAll(res, "\"", "")
	return "\"" + res1 + "\""
}

func getInput() (string, string, string, error) {
	var a, operator, b string
	fmt.Scan(&a, &operator, &b)

	// Проверка длины входных данных
	if len(a) > 10 || len(b) > 10 {
		return "", "", "", fmt.Errorf("входные данные слишком длинны")
	}

	// Проверка оператора
	validOperators := []string{"+", "-", "*", "/"}
	if !contains(validOperators, operator) {
		return "", "", "", fmt.Errorf("неверный оператор")
	}

	// Проверка строк на целостность и диапазон от 1 до 10
	//numA, err := strconv.Atoi(a)
	//if err != nil || numA < 1 || numA > 10 {
	//return "", "", "", fmt.Errorf("число A должно быть от 1 до 10")
	//}

	//numB, err := strconv.Atoi(b)
	//if err != nil || numB < 1 || numB > 10 {
	//return "", "", "", fmt.Errorf("число B должно быть от 1 до 10")
	//}

	return a, operator, b, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	a, operator, b, err := getInput()
	if err != nil {
		panic(err)
	}

	var result string

	switch operator {
	case "+":
		result = add(a, b)
	case "-":
		result = substraction(a, b)
	case "*":
		result = multiplication(a, b)
	case "/":
		result = division(a, b)
	default:
		panic(fmt.Errorf("неверный оператор"))

	}

	// Обрезка результата до 40 символов
	if len(result) > 40 {
		result = result[:41] + "..."
	}
	fmt.Printf(result)
}
