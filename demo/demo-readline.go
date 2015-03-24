// +build ignore
package main
import (
	"fmt"; "os"; "io"
	"github.com/rocky/go-gnureadline"
)

func print_edit_mode() {
	if gnureadline.Rl_editing_mode() == gnureadline.Emacs {
		fmt.Println("Editing mode is emacs")
	} else {
		fmt.Println("Editing mode is vi")
	}
	fmt.Printf("Double check: %s\n",
		gnureadline.Rl_variable_value("editing-mode"))

}

func print_insert_mode() {
	if gnureadline.Rl_insert_mode() {
		fmt.Println("Insert mode on")
	} else {
		fmt.Println("Overwrite mode on")
	}
}
func main() {
	var line string
	term := os.ExpandEnv("TERM")

	fmt.Println("Genuine GNU Readline?", gnureadline.Rl_gnu_readline_p())
	fmt.Printf("Your readline library version is %s\n",
		gnureadline.Rl_readline_library_version())
	fmt.Printf("Your readline version is %x\n",
		gnureadline.Rl_readline_version())
	fmt.Printf("Your readline terminal is %s\n",
		gnureadline.Rl_readline_terminal_name())
	fmt.Printf("Your readline terminal is %s\n",
		gnureadline.Rl_readline_name())
	fmt.Printf("rl_prefer_env_winsize is %d\n",
		   gnureadline.Rl_prefer_env_winsize())
	varlist := []string{
			"bell_style", "comment-begin", "completion-prefix-display-length",
			"completion-query-items", "editing-mode", "history-size",
			"isearch-terminators", "keymap",
			"bind-tty-special-chars",
			"blink-matching-paren",
			"byte-oriented",
			"completion-ignore-case",
			"convert-meta",
			"disable-completion",
			"enable-keypad",
			"expand-tilde",
			"history-preserve-point",
			"horizontal-scroll-mode",
			"input-meta",
			"mark-directories",
			"mark-modified-lines",
			"mark-symlinked-directories",
			"match-hidden-files",
			"meta-flag",
			"output-meta",
			"page-completions",
			"prefer-visible-bell",
			"print-completions-horizontally",
			"revert-all-at-newline",
			"show-all-if-ambiguous",
			"show-all-if-unmodified",
		}
	for _, variable := range(varlist) {
	     fmt.Printf("variable_value for %s is: %s\n",
			variable,  gnureadline.Rl_variable_value(variable))
	}

	print_edit_mode()
	print_insert_mode()
	fmt.Println("Reading data/undo.inputrc")
	gnureadline.Rl_read_init_file("data/undo.inputrc")

	line, err := gnureadline.Readline("Enter something without history: ")
	for i:=1; err == nil && line != "quit"; i++ {
		line, err = gnureadline.Readline(fmt.Sprintf("Enter something %d: ",
			i), true)
		switch line {
		case "vi":
			gnureadline.Rl_editing_mode_set(gnureadline.Vi)
			print_edit_mode()
		case "emacs":
			gnureadline.Rl_editing_mode_set(gnureadline.Emacs)
			print_edit_mode()
		case "insert":
			gnureadline.Rl_insert_mode_set(true)
			print_insert_mode()
		case "overwrite":
			gnureadline.Rl_insert_mode_set(false)
			print_insert_mode()
		}
		fmt.Printf("You typed: %s\n", line)
		fmt.Printf("Byte in history %d, position %d\n",
			gnureadline.HistoryTotalBytes(),
			gnureadline.WhereHistory())
	}
	if err == io.EOF {
		fmt.Println("Got EOF")
	}
	fmt.Printf("History length %d\n",  gnureadline.HistoryLength())
	fmt.Printf("History max entries %d\n",  gnureadline.HistoryMaxEntries())
	fmt.Println("writing: data/deleteme.history")
	gnureadline.WriteHistory("data/deleteme.history")
	gnureadline.Rl_reset_terminal(term)
}
