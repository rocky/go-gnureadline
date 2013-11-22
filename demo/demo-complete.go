package main
import (. "gnureadline"; "sort"; "fmt"; "strings")

var completions = []string{
	"spoken",
	"spokes",
	"spokesman",
	"spokesmen",
	"spokesperson",
	"sponge",
	"sponged",
	"sponger",
	"spongers",
	"sponges",
	"sponging",
	"spongy",
	"sponsor",
	"sponsored",
	"sponsoring",
	"sponsors",
	"sponsorship",
	"spontaneity",
	"spontaneous",
	"spontaneously",
}

func main() {
	fmt.Printf("Completer word-break characters: '%s'\n", 
		CompleterWordBreakCharacters())
	CompleterWordBreakCharacters_("X ")
	fmt.Printf("Completer word-break characters: '%s'\n", 
		CompleterWordBreakCharacters())

	SetAttemptedCompletionFunction(func(text string, start, end int) []string {

		// Binary search for the range of completions with a matching prefix
		n := len(completions)
		top := sort.Search(n, func(i int) bool { return text <= completions[i] })
		bot := sort.Search(n, func(i int) bool { return i > top && !strings.HasPrefix(completions[i], text) })

		if bot == top {
			// Returning nil indicates no completions
			return nil
		} else if bot - top == 1 {
			// One match found, it will be tab completed
			return completions[top:bot]
		} else {
			// More than one possible match.
			// The first element of the returned slice will be inserted on the command line
			// The remaining elements will be listed if the user presses tab twice

			completion := commonPrefix(completions[top], completions[bot-1])
			return append([]string{completion}, completions[top:bot]...)
		}
	})

	var line string

	line, err := Readline("something beginning with s> ")
	for err == nil && line != "quit" {
		line, err = Readline("something beginning with s> ")
	}
}

func commonPrefix(a, b string) string {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	var i int
	for i = 0; i < n && a[i] == b[i]; i += 1 {}

	return a[:i]
}
