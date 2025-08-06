package internals

import "fmt"

func PadZero(number uint32) string {
	return fmt.Sprintf("%02d", number)
}
