package demo

import "fmt"

func Method(a int, b int) (string, int) {
	return fmt.Sprintf("%d%d", a, b), a + b
}
