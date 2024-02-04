package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func conv(x string) int {
	switch x {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	}
	panic("Неверное значение")
	return 0
}

func revconvert(x int) string {
	switch x {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VII"
	case 9:
		return "IX"
	case 10:
		return "X"
	}
	return ""
}

func calculate(a []string) string {

	if len(a) != 3 {
		panic("Неверное оформление задания")
	}
	v1, error1 := strconv.Atoi(a[0])
	v2, error2 := strconv.Atoi(a[2])
	if (v1 >= 10) || (v1 <= 0) || (v2 >= 10) || (v2 <= 0) {
		panic("Недопустимые значения")
	}
	if error1 != error2 {
		panic("Числа разных типов")
	}
	if error1 != nil {
		v1 = conv(a[0])
		v2 = conv(a[2])
		if a[1] == "-" {
			cal := v1 - v2
			if cal <= 0 {
				panic("Значение меньше 0")
			}
			return revconvert(cal)
		} else if a[1] == "+" {
			return revconvert(v1 + v2)
		} else if a[1] == "*" {
			return revconvert(v1 * v2)
		} else if a[1] == "/" {
			return revconvert(v1 / v2)
		} else {
			panic("Недопустимая операция")
		}
	}
	if a[1] == "-" {
		return strconv.Itoa(v1 - v2)
	} else if a[1] == "+" {
		return string(v1 + v2)
	} else if a[1] == "*" {
		return string(v1 * v2)
	} else if a[1] == "/" {
		return string(v1 / v2)
	} else {
		panic("Недопустимая операция")
	}

	return ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		_text, _ := reader.ReadString('\n')
		words := strings.Fields(_text)
		fmt.Println(calculate(words))
	}
	return
}
