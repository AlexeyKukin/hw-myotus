package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var (
	onespace = regexp.MustCompile(`\s+`)
	reg      = regexp.MustCompile(`[,.!:";]`)
)

func Top10(rawstr string) []string {
	toprate := make(map[string]int, 200)
	sortedval := make([]int, 0, 200)
	rate := make([]int, 0, 10)
	keys := make([]string, 0, 200)
	result := make([]string, 0, 10)
	var prev int

	// Если строка пустая, сразу возвращаем результат
	if rawstr == "" {
		return result
	}

	// Удаляем из строки двойные пробелы, символы перевода строк и т.д.
	rawstr = onespace.ReplaceAllString(rawstr, " ")
	top := strings.Split(rawstr, " ")
	for k := range top {
		top[k] = strings.ToLower(top[k])
		top[k] = reg.ReplaceAllString(top[k], "")
	}
	sort.Strings(top)

	// 	Формируем мапу ключ[слово] : значение [int],
	for _, v := range top {
		if v != "-" {
			toprate[v]++
		}
	}

	// Создаем сортированный слайс "sortedval" из значений [int] мапы.
	for _, v := range toprate {
		sortedval = append(sortedval, v)
	}
	sort.Ints(sortedval)

	// Создаем слайс "rate" из десяти самых больших значений [int].
	for k := len(sortedval) - 1; k >= 0; k-- {
		if len(rate) < 10 {
			rate = append(rate, sortedval[k])
		}
	}

	// Создаем слайс ключей отсортированных по алфавиту.
	for k := range toprate {
		keys = append(keys, k)
	}

	// Уникальные стринговые ключи отсортированные по алфавиту.
	sort.Strings(keys)

	// Формируем финальный слайс топ символов. для этого проходим по слайсу rate[int] и range[string].
	for _, val := range rate {
		if prev != val {
			for _, v := range keys {
				if toprate[v] == val {
					result = append(result, v)
					if len(result) == 10 {
						break
					}
				}
				prev = val
			}
		} else {
			continue
		}
		if len(result) == 10 {
			break
		}
	}

	return result
}
