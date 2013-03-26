package main
import ( . "gnureadline"; "fmt")

func main() {
	fmt.Printf("Keymap name is %s\n", GetKeymapNameFromEditMode())
}
