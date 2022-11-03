package helpers

import "fmt"

func RgbaToHex(R uint32, G uint32, B uint32, A uint32) string {
	return fmt.Sprintf("#%02x%02x%02x", R>>8, G>>8, B>>8)
}
