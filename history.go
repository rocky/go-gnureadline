/* 
   GNU Readline History functions

   Copyright (C) 2013 Rocky Bernstein

   Readline is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   Readline is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with Readline.  If not, see <http://www.gnu.org/licenses/>.
*/

package gnureadline
/*
#cgo LDFLAGS: -lreadline
#include <stdio.h>
#include <readline/readline.h>
#include <readline/history.h>
#include <stdlib.h> // for free()
*/
import "C"

/* Place LINE at the end of the history list.
   The associated data field (if any) is set to NULL. */
func AddHistory(line string) {
	c_line := C.CString(line)
	C.add_history(c_line)
}

/* Clear the history list and start over. */
func ClearHistory() {
	C.clear_history()
}

/* Stifle the history list, remembering only MAX number of entries. */
func StifleHistory(max int) {
	C.stifle_history(C.int(max))
}

/* Stop stifling the history.  This returns the previous amount the
   history was stifled by.  The value is positive if the history was
   stifled, negative if it wasn't. */
func UnstifleHistory() int {
	return int(C.unstifle_history())
}

/* Return true if the history is stifled, false if it is not. */
func HistoryIsStifled() bool {
	var b = int(C.history_is_stifled())
	if b == 0 { return false }
	return true
}


