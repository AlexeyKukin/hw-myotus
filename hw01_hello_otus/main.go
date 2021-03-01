package main

// Импортируем модули необходимые для выполнения задания.
import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	// Выводим перевенутую фразу "Hello, OTUS!" используя модуль stringutil.
	fmt.Print(stringutil.Reverse("Hello, OTUS!"))
}
