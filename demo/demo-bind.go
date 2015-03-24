package main
import (
	"fmt"
	"github.com/rocky/go-gnureadline"
)

func main() {
	fmt.Printf("Keymap name is %s\n",
		gnureadline.GetKeymapNameFromEditMode())
}
