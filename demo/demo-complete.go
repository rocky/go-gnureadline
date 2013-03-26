package main
import ( . "gnureadline"; "fmt")

func main() {
	fmt.Printf("Completer word-break characters: '%s'\n", 
		CompleterWordBreakCharacters())
	CompleterWordBreakCharacters_("X")
	fmt.Printf("Completer word-break characters: '%s'\n", 
		CompleterWordBreakCharacters())
}
