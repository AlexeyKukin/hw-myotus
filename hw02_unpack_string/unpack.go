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
"aaa0b" => "aab"        "" => ""        "d\n5abc" => "d\n\n\n\n\nabc".
*/
func Unpack(input string) (string, error) {
	var cache int32
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
				if cache != 0 {
					unpacked.WriteRune(cache)
				}
				cache = v
			}
		// Если символ цифра (n) и перед ней есть буква в кеше, "распаковываем кеш" n раз.
		case unicode.IsDigit(v):
			{
				if !unicode.IsDigit(cache) && cache != 0 {
					for i := v - 48; i > 0; i-- {
						unpacked.WriteRune(cache)
					}
					cache = 0
				} else {
					return "", ErrInvalidString
				}
			}

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
