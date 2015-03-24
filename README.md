[![Build Status](https://travis-ci.org/rocky/go-gnureadline.png)](https://travis-ci.org/rocky/go-gnureadline)

Go bindings to the [GNU Readline](http://cnswww.cns.cwru.edu/php/chet/readline/rltop.html) library

This provides command-line entry, editing and command history


Synopsis
--------

```go
package main
import (
    "code.google.com/p/go-gnureadline"
    "fmt"
    "os"
)

const HISTORY_FILE string = "my.history"

func main() {
     var err error
     term := os.Getenv("TERM")
     gnureadline.ReadHistory(HISTORY_FILE)
     gnureadline.StifleHistory(10)  // Maximum 10 history entries
     gnureadline.ReadInitFile(".inputrc")  // Read in a keybinding initialization file
     line := ""
     for i:=1; err == nil && line != "quit"; i++ {
             line, err = gnureadline.Readline(fmt.Sprintf("Enter something [%d]: ", i), true)
             switch line {
                case "vi":
                        gnureadline.Rl_editing_mode_set(Vi)
                case "emacs":
                        gnureadline.Rl_editing_mode_set(Emacs)
                case "insert":
                        gnureadline.Rl_insert_mode_set(true)
                case "overwrite":
                        gnureadline.Rl_insert_mode_set(false)
                }
                fmt.Printf("You typed: %s\n", line)
     }
     fmt.Printf("History length %d\n",  gnureadline.HistoryLength())
     fmt.Printf("History max entries %d\n",  gnureadline.HistoryMaxEntries())
     gnureadline.WriteHistory(gnureadline.HISTORY_FILE)
     gnureadline.Rl_reset_terminal(term)
}
```

Description
-----------

readline will read a line from the terminal and return it, using
prompt as a prompt.  If prompt he empty string, no prompt is issued.
The line returned has the final newline removed, so only the text of
the line remains.

readline offers editing capabilities while the user is entering the
line.  By default, the line editing commands are similar to those of
emacs.  A vi-style line editing interface is also available.
