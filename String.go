package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func exampleString() {
	// Создадим строковый литерал s, значение которого "Это строка".
	// Строка состоит из 10 символов.
	var s string = "Это строка"

	// Однако длина строки len(s) составит 19 байт, т.к. использованные кирилические символы
	// занимают 2 байта, а пробел занимает 1 байт.
	fmt.Printf("Длина строки: %d байт\n", len(s))

	// Посмотрим как взять подстроку
	fmt.Printf("Напечатаем только второе слово в кавычках: \"%v\"\n", s[7:])

	/*
		Попробуем изменить что-то встроке:
		s[3] = 12
		Ошибка компиляции: cannot assign to s[3], потому что строки - неизменяемые последовательности.
	*/

	// "Изменим строку", создав новую
	s = s + " Новая строка"
	fmt.Printf("%v\n", s)

	// А теперь проитерируемся по этой строке
	for _, b := range s {
		fmt.Printf("%v ", b)
	}
	fmt.Print("\n")

	// Output:
	// Длина строки: 19 байт
	// Напечатаем только второе слово в кавычках: "строка"
	// Это строка Новая строка
	// 1069 1090 1086 32 1089 1090 1088 1086 1082 1072 32 1053 1086 1074 1072 1103 32 1089 1090 1088 1086 1082 1072

}
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func methodString() {
	/*	// Содержится ли подстрока в строке
		strings.Contains("test", "es"),
		// результат: true

		// Кол-во подстрок в строке
			strings.Count("test", "t"),
		// результат: 2

		// Начинается ли строка с префикса
			strings.HasPrefix("test", "te"),
		// результат: true

		// Заканчивается ли строка суффиксом
			strings.HasSuffix("test", "st"),
		// результат: true

		// Возвращает начальный индекс подстроки в строке, а при отсутствии вхождения возвращает -1
			strings.Index("test", "e"),
		// результат: 1

		// объединяет массив строк через символ
			strings.Join([]string{"hello","world"}, "-"),
		// результат: "hello-world"

		// Повторяет строку n раз подряд
			strings.Repeat("a", 5),
		// результат: "aaaaa"

		// Функция Replace заменяет любое вхождение old в вашей строке на new
		// Если значение n равно -1, то будут заменены все вхождения.
		// Общий вид: func Replace(s, old, new string, n int) string
		// Пример:
			strings.Replace("blanotblanot", "not", "***", 	-1),
		// результат: "bla***bla***"

		// Разбивает строку согласно разделителю
			strings.Split("a-b-c-d-e", "-"),
		// результат: []string{"a","b","c","d","e"}

		// Возвращает строку c нижним регистром
			strings.ToLower("TEST"),
		// результат: "test"

		// Возвращает строку c верхним регистром
			strings.ToUpper("test"),
		// результат: "TEST"

		// Возвращает строку с вырезанным набором
			strings.Trim("tetstet", "te"),
		// результат: s

	*/
}
func ExampleRune() {
	// Поступим также, как в работе с типом []byte
	rs := []rune("Это срез рун")

	// Итерируясь мы будем заменять символ 'р' на '*'
	for i := range rs {
		if rs[i] == 'р' {
			rs[i] = '*'
		}
	}
	fmt.Printf("Измененнный срез в виде строки: %s\n", string(rs))

	// Output:
	// Измененнный срез в виде строки: Это с*ез *ун
}

func format() {
	var a float64 = 1.0123456789

	// 1 параметр - число для конвертации
	// fmt - форматирование
	// prec - точность (кол-во знаков после запятой)
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	fmt.Println(strconv.FormatFloat(a, 'f', 2, 64)) // 1.01

	// если мы хотим учесть все цифры после запятой, то можем в prec передать -1
	fmt.Println(strconv.FormatFloat(a, 'f', -1, 64)) // 1.0123456789

	// Возможные форматы fmt:
	// 'f' (-ddd.dddd, no exponent),
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'g' ('e' for large exponents, 'f' otherwise),
	// 'G' ('E' for large exponents, 'f' otherwise),
	// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
	// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
	var b float64 = 2222 * 1023 * 245 * 2 * 52
	fmt.Println(strconv.FormatFloat(b, 'e', -1, 64)) // 5.791874088e+10

}

func unicodeStr() {
	// функции ниже принимают на вход тип rune

	// проверка символа на цифру
	fmt.Println(unicode.IsDigit('1')) // true
	// проверка символа на букву
	fmt.Println(unicode.IsLetter('a')) // true
	// проверка символа на нижний регистр
	fmt.Println(unicode.IsLower('A')) // false
	// проверка символа на верхний регистр
	fmt.Println(unicode.IsUpper('A')) // true
	// проверка символа на пробел
	// пробел это не только ' ', но и:
	//  '\t', '\n', '\v', '\f', '\r' - подробнее читайте в документации
	fmt.Println(unicode.IsSpace('\t')) // true

	// С помощью функции Is можно проверять на кастомный RangeTable:
	// например, проверка на латиницу:
	fmt.Println(unicode.Is(unicode.Latin, 'ы')) // false

	// функции преобразований
	fmt.Println(string(unicode.ToLower('F'))) // f
	fmt.Println(string(unicode.ToUpper('f'))) // F

	// input := bufio.NewScanner(os.Stdin)
	// for {
	// 	if input.Scan() {
	// 		result := input.Text()
	// 		if result == "" {
	// 			break
	// 		}
	// 		fmt.Println(result)
	// 	}
	// }
}
