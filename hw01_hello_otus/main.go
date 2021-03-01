package main

// Импортируем необходимые модули.
import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	// Используем "stringutil" для перевертывания фразы как указано в требованиях к задаче.
	fmt.Print(stringutil.Reverse("Hello, OTUS!"))
}
