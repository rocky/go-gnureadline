package main
import ( . "gnureadline"; "fmt")
func main() {
	var line string
	for i:=1; line != "quit"; i++ {
		line = Readline(fmt.Sprintf("Enter something %d: ", i))
		fmt.Printf("You typed: %s\n", line)
	}
}
