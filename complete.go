/* 
   GNU Readline completion functions

   Copyright (C) 2013 Rocky Bernstein

   gnureadline is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   gnureadline is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with gnureadline.  If not, see <http://www.gnu.org/licenses/>.
*/

/* This file contains things from readline/complete.c */

/* Thanks to Sebastien Binet from which these routines were started
from. */


package gnureadline
/*
#cgo darwin CFLAGS: -I/opt/local/include
#cgo darwin LDFLAGS: -L/opt/local/lib
#cgo LDFLAGS: -lreadline
#include <stdio.h>
#include <readline/readline.h>
#include <stdlib.h> // for free()

char* _go_readline_strarray_at(char **strarray, int idx) 
 {
   return strarray[idx];
 }

 int _go_readline_strarray_len(char **strarray)
 {
   int sz = 0;
   while (strarray[sz] != NULL) {
     sz += 1;
   }
   return sz;
 }
*/
import "C"
import "unsafe"

/* 
 The list of characters that signal a break between words.
 The default list is the contents of
 rl_basic_word_break_characters.
  */
func CompleterWordBreakCharacters() string {
	cstr := C.rl_completer_word_break_characters
	delims := C.GoString(cstr)
	return delims
}

/* Setter for CompleteWordBreakCharacters */
func CompleterWordBreakCharacters_(break_chars string) {
	p := C.CString(break_chars)
	//defer C.free(unsafe.Pointer(p))
	C.free(unsafe.Pointer(C.rl_completer_word_break_characters))
	C.rl_completer_word_break_characters = p
}

/* 

Return a slice which is a list of completions for TEXT.  If there are
no completions, return [].

The first entry in the returned array is the substitution for TEXT.
The remaining entries are the possible completions.  The array is
terminated with a NULL pointer.
 
ENTRY_FUNCTION is a function of two args, and returns a (char *).  The
first argument is TEXT.
 
The second is a state argument; it should be zero on the first call,
and non-zero on subsequent calls.  It returns a NULL pointer to the
caller when there are no more matches.  */

func CompletionMatches(text string, 
	entry_func func(text string, state int) string) []string {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))
	c_entry_func := (*C.rl_compentry_func_t)(unsafe.Pointer(&entry_func))
	c_matches := C.rl_completion_matches(c_text, c_entry_func)
	n_matches := int(C._go_readline_strarray_len(c_matches))
	matches := make([]string, n_matches)
	for i := 0; i < n_matches; i++ {
		matches[i] = C.GoString(C._go_readline_strarray_at(c_matches, 
			C.int(i)))
	}
	return matches
}

//
func SetAttemptedCompletionFunction(entry_func func(text string, 
	start, end int) []string) {
	c_entry_func := (*C.rl_completion_func_t)(unsafe.Pointer(&entry_func))
	C.rl_attempted_completion_function = c_entry_func
}
