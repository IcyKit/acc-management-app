package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный код ошибки")
	}
}

func sum[T int | float32 | float64 | string](a, b T) T {
	return a + b
}
