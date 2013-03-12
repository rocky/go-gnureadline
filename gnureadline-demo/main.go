package main
import ( . "gnureadline"; "fmt"; "os")

func print_edit_mode() {
	if Rl_editing_mode() == Emacs {
		fmt.Println("Editing mode is emacs")
	} else {
		fmt.Println("Editing mode is vi")
	}
	fmt.Printf("Double check: %s\n", Rl_variable_value("editing-mode"))

}

func print_insert_mode() {
	if Rl_insert_mode() == 1 {
		fmt.Println("Insert mode on")
	} else {
		fmt.Println("Overwrite mode on")
	}
}
func main() {
	var line string
	term := os.ExpandEnv("TERM")

	fmt.Println("Genuine GNU Readline?", Rl_gnu_readline_p())
	fmt.Printf("Your readline library version is %s\n", 
		Rl_readline_library_version())
	fmt.Printf("Your readline version is %x\n", 
		Rl_readline_version())
	fmt.Printf("Your readline terminal is %s\n", 
		Rl_readline_terminal_name())
	fmt.Printf("Your readline terminal is %s\n", 
		Rl_readline_name())
	fmt.Printf("rl_prefer_env_winsize is %d\n", 
		   Rl_prefer_env_winsize())
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
			variable,  Rl_variable_value(variable))
	}

	print_edit_mode()
	print_insert_mode()
	fmt.Printf("History length %d\n",  HistoryLength())
	fmt.Printf("History max entries %d\n",  HistoryMaxEntries())
	fmt.Println("History is stifled", HistoryIsStifled())
	StifleHistory(3)
	fmt.Println("History is stifled", HistoryIsStifled())
	fmt.Println("Reading data/undo.inputrc")
	Rl_read_init_file("data/undo.inputrc")

	line = Readline("Enter something without history: ")
	for i:=1; line != "quit"; i++ {
		line = Readline(fmt.Sprintf("Enter something %d: ", i), true)
		switch line {
		case "vi":
			Rl_editing_mode_set(Vi)
			print_edit_mode()
		case "emacs":
			Rl_editing_mode_set(Emacs)
			print_edit_mode()
		case "insert":
			Rl_insert_mode_set(1)
			print_insert_mode()
		case "overwrite":
			Rl_insert_mode_set(0)
			print_insert_mode()
		}
		fmt.Printf("You typed: %s\n", line)
		fmt.Printf("Byte in history %d, position %d\n", 
			HistoryTotalBytes(), WhereHistory())
	}
	fmt.Printf("History length %d\n",  HistoryLength())
	fmt.Printf("History max entries %d\n",  HistoryMaxEntries())
	fmt.Println("writing: data/deleteme.history")
	WriteHistory("data/deleteme.history")
	Rl_reset_terminal(term)
}
