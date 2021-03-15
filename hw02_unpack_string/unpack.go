package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

/*
Функция осуществляет примитивную распаковку строки, содержащую повторяющиеся символы/руны примеры:
"a4bc2d5e" => "aaaabccddddde"   "abcd" => "abcd"
"3abc" => "" (некорректная строка)      "45" => "" (некорректная строка)        "aaa10b" => "" (некорректная строка)
"aaa0b" => "aab"        "" => ""        "d\n5abc" => "d\n\n\n\n\nabc"
Дополнительно:
"qwe\4\5" => "qwe45"	"qwe\45" => "qwe44444"	"qwe\\5a" => "qwe\\\\\a".
*/
func Unpack(input string) (string, error) {
	var cache int32
	var quoted bool
	var unpacked strings.Builder
	// Запускаем перебор по строке
	for k, v := range input {
		switch {
		// Если первый символ цифра - выдаем ошибку.
		case (k == 0 && unicode.IsDigit(v)):
			{
				return "", ErrInvalidString
			}
		// Если символ буква или управляющий символ выводим символ из кеша, либо вносим его в кеш если он пуст.
		case (unicode.IsLetter(v) || unicode.IsControl(v)):
			{
				if quoted {
					return "", ErrInvalidString
				}
				if cache != 0 {
					unpacked.WriteRune(cache)
				}
				cache = v
			}
		// Если символ цифра (n) и перед ней есть буква в кеше, "распаковываем кеш" n раз.
		case unicode.IsDigit(v) && !quoted:
			{
				if /*!unicode.IsDigit(cache) && */ cache != 0 {
					for i := v - 48; i > 0; i-- {
						unpacked.WriteRune(cache)
					}
					cache = 0
				} else {
					return "", ErrInvalidString
				}
			}
		// Если цифра экранирована
		case unicode.IsDigit(v) && quoted:
			{
				cache = v
				quoted = false
			}
		// Обрабатываем экранирование через \ используя флаг quoted.
		case v == 92:
			{
				if !quoted {
					unpacked.WriteRune(cache)
					quoted = true
				} else {
					cache = v
					quoted = false
				}
			}
		// Если встретили непонятный символ - выводим ошибку.
		default:
			{
				return "", ErrInvalidString
			}
		}
	}
	// Если в кеше остался последний символ выводим его.
	if cache != 0 {
		unpacked.WriteRune(cache)
	}
	return unpacked.String(), nil
}
