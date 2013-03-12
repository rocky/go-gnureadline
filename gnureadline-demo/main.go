package main
import ( . "gnureadline"; "fmt"; "os")

func print_edit_mode() {
	if Rl_editing_mode() == Emacs {
		fmt.Println("Editing mode is emacs")
	} else {
		fmt.Println("Editing mode is vi")
	}
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
	print_edit_mode()
	print_insert_mode()
	fmt.Println("History is stifled", HistoryIsStifled())
	StifleHistory(3)
	fmt.Println("History is stifled", HistoryIsStifled())
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
	fmt.Println("writing: deleteme.history")
	WriteHistory("deleteme.history")
	Rl_reset_terminal(term)
}
