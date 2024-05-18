package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	var sliced_input []string
	operators := []string{"+", "-", "*", "/"}
	var operators_count int = 0
	var arabic_numbers_input string = "0123456789"
	var roman_numbers_input string = "IVX"
	contains_arabic_numbers, contains_roman_numbers := false, false
	roman_numbers_result := [101]string{"0_placeholder", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	var number_1 int
	var number_2 int
	var action string = ""
	var result int

	//приветственное сообщение
	fmt.Println("Здравствуйте!\n", "Это программа Калькулятор.\n", "Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами\n", "Формат ввода:\n", "a+b, a-b, a*b, a/b БЕЗ ПРОБЕЛОВ!!!\n", "Калькулятор поддерживает арабские и римские числа от 1 до 10 (от I до X) включительно.\n", "Калькулятор умеет работать только с целыми числами.\n", "Использовать две разные системы записи в одном вводе, а так же совершать более одной операции в рамках ввода запрещено.\n", "При вводе пользователем некорректного ввода программа сообщит об этом и завершит работу.")

	//ввод
	fmt.Scanln(&input)

	//проверка ввода на наличие операнда, при том ТОЛЬКО ОДНОГО
	for x := 0; x < len(input); x++ {
		if input[x] == '+' || input[x] == '-' || input[x] == '*' || input[x] == '/' { //поподробнее почитать про разницу одинарной и двойной кавычки и почему символы в строках являются байтами и их нельзя сравнивать со строками (почему вообще нет типа данных char???????????)
			operators_count++
		}
	}
	if operators_count > 1 {
		fmt.Println("Ошибка, ввод содержит более одного оператора")
		os.Exit(38)
	} else if operators_count == 0 {
		fmt.Println("Ошибка, ввод не содержит ни одного оператора")
		os.Exit(41)
	}

	//проверка ввода на то что операнд стоит там где надо, а именно посередине между числами (иначе строка неправильно поделится на слайсы)
	for _, operator := range operators {
		if strings.HasPrefix(input, operator) || strings.HasSuffix(input, operator) {
			fmt.Println("Ошибка, оператор должен стоять между числами а не в начале или конце строки")
			os.Exit(48)
		}
	}

	//разделяем строку на два слайса содержащие первый и второй операнд, а так же задаём оператор
	if strings.Contains(input, "+") {
		sliced_input = strings.Split(input, "+")
		action = "+" //почему этого не достаточно чтобы он не ругался на переменную action как на 'not used'? потому что внутри if? уточнить.
	} else if strings.Contains(input, "-") {
		sliced_input = strings.Split(input, "-")
		action = "-"
	} else if strings.Contains(input, "*") {
		sliced_input = strings.Split(input, "*")
		action = "*"
	} else if strings.Contains(input, "/") {
		sliced_input = strings.Split(input, "/")
		action = "/"
	}
	//Проверяем что слайсы содержат только допустимые символы а так же что в них совпадают системы записи
	if strings.Trim(sliced_input[0], arabic_numbers_input) == "" && strings.Trim(sliced_input[1], arabic_numbers_input) == "" {
		contains_arabic_numbers = true
	} else if strings.Trim(sliced_input[0], roman_numbers_input) == "" && strings.Trim(sliced_input[1], roman_numbers_input) == "" {
		contains_roman_numbers = true
	} else {
		fmt.Println("Ошибка, ввод содержит недопустимые символы или вы пытаетесь оперировать числами двух разных систем записи (арабские с римскими)")
		os.Exit(73)
	}

	//переводим арабские числа из string в int чтобы провести над ними необходимую математическую операцию
	if contains_arabic_numbers {
		number_1, _ = strconv.Atoi(sliced_input[0])
		number_2, _ = strconv.Atoi(sliced_input[1])
	}

	//переводим римские числа в арабские чтобы провести над ними необходимую математическую операцию
	if contains_roman_numbers {
		for i := 1; i <= 10; i++ {
			if roman_numbers_result[i] == sliced_input[0] {
				number_1 = i
			}
			if roman_numbers_result[i] == sliced_input[1] {
				number_2 = i
			}
		}
	}

	//проверяем что введенные числа входят в диапазон от 1 до 10
	if number_1 < 1 || number_2 < 1 || number_1 > 10 || number_2 > 10 {
		fmt.Println("Ошибка, ввод содержит число меньше единицы или больше десяти")
		os.Exit(97)
	}

	//непосредственно сами вычисления
	switch {
	case action == "+":
		result = number_1 + number_2
	case action == "-":
		result = number_1 - number_2
	case action == "*":
		result = number_1 * number_2
	case action == "/":
		result = number_1 / number_2
	}

	switch {
	case contains_arabic_numbers:
		fmt.Println("РЕЗУЛЬТАТ: ", result, "\n", "____________________________________________")
	case contains_roman_numbers:
		if result < 1 {
			fmt.Println("Ошибка, результат меньше единицы, римская система записи не поддерживает такие числа")
			os.Exit(118)
		} else {
			fmt.Println("РЕЗУЛЬТАТ: ", roman_numbers_result[result], "\n", "____________________________________________")
		}

	}

	// fmt.Println("sliced_input[0]: ", sliced_input[0])
	// fmt.Println("sliced_input[1]: ", sliced_input[1])
	//fmt.Println("action: ", action) //пришлось нагородить эту чушь чтобы компилятор отвалил с ошибкой "not used". уточнить как обойти эту ошибку если переменная использутся в if но объявить я ее хочу глобально, а не внутри if.
	//fmt.Println("contains_arabic_numbers: ", contains_arabic_numbers)
	//fmt.Println("contains_roman_numbers:", contains_roman_numbers)
	//fmt.Println("number_1: ", number_1)
	//fmt.Println("number_2: ", number_2)
}
