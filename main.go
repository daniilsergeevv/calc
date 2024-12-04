package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Удаляем символ новой строки

	result, err := calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}

func calculate(input string) (string, error) {
	// Регулярное выражение для разделения строки по символам + - * /
	re := regexp.MustCompile(`(["][^"]*["])\s*([-+*/])\s*([0-9]+|["][^"]*["])`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 4 {
		return "", fmt.Errorf("неверный формат ввода")
	}

	a := matches[1]
	operator := matches[2]
	b := matches[3]

	// Проверяем, что первый аргумент является строкой
	if !strings.HasPrefix(a, "\"") || !strings.HasSuffix(a, "\"") {
		return "", fmt.Errorf("первый аргумент должен быть строкой")
	}
	a = strings.Trim(a, "\"")

	// Проверяем, что второй аргумент является числом или строкой в зависимости от оператора
	if operator == "*" || operator == "/" {
		if _, err := strconv.Atoi(b); err != nil {
			return "", fmt.Errorf("второй аргумент должен быть числом для операций умножения и деления")
		}
	} else {
		if !strings.HasPrefix(b, "\"") || !strings.HasSuffix(b, "\"") {
			return "", fmt.Errorf("второй аргумент должен быть строкой для операций сложения и вычитания")
		}
		b = strings.Trim(b, "\"")
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
		return "", fmt.Errorf("неверный оператор")
	}

	// Ограничиваем длину результата до 40 символов
	if len(result) > 40 {
		result = result[:41] + "..."
	}

	return result, nil
}

func add(a, b string) string {
	result := a + b
	return "\"" + result + "\""
}

func substraction(a, b string) string {
	result := strings.ReplaceAll(a, b, "")
	return "\"" + result + "\""
}

func multiplication(a, b string) string {
	intVal, _ := strconv.Atoi(b)
	if intVal < 1 || intVal > 10 {
		panic("Число должно быть от 1 до 10")
	}
	copy := strings.Repeat(a, intVal)
	return "\"" + copy + "\""
}

func division(a, b string) string {
	intVal, _ := strconv.Atoi(b)
	if intVal < 1 || intVal > 10 {
		panic("Число должно быть от 1 до 10")
	}
	length := len(a)
	res := a[:length/intVal]
	return "\"" + res + "\""
}
