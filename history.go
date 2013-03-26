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

// This file contains things from readline/history.h

package gnureadline
/*
#cgo LDFLAGS: -lreadline
#include <stdio.h>
#include <readline/history.h>
#include <stdlib.h> // for free()
*/
import "C"
import "unsafe"

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

/* Return true if the history is stifled, false if it is not. */
func HistoryIsStifled() bool {
	var b = int(C.history_is_stifled())
	if b == 0 { return false }
	return true
}

/* Return the number of bytes that the primary history entries are using.
   This just adds up the lengths of the_history->lines. */
func HistoryTotalBytes() int {
	return int(C.history_total_bytes())
}

/* Set the position in the history list to POS. */
func HistorySetPos(pos int) {
	C.history_set_pos(C.int(pos))
}

/* Add the contents of FILENAME to the history list, a line at a time.
   If FILENAME is '', then read from ~/.history.  Returns 0 if
   successful, or errno if not. */
func ReadHistory(filename string) int {
	if len(filename) == 0 {
		return int(C.read_history(nil))
	}
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	return int(C.read_history(c_filename))
}

/* Stifle the history list, remembering only MAX number of entries. */
func StifleHistory(max int) {
	C.stifle_history(C.int(max))
}

/* Returns the number which says what history element we are now
   looking at.  */
func WhereHistory() int {
	return int(C.where_history())
}

/* Write the current history to FILENAME.  If FILENAME is '',
   then write the history list to ~/.history.  Values returned
   are as in ReadHistory().  */
func WriteHistory(filename string) int {
	if len(filename) == 0 {
		return int(C.write_history(nil))
	}
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	return int(C.write_history(c_filename))
}

/* Stop stifling the history.  This returns the previous amount the
   history was stifled by.  The value is positive if the history was
   stifled, negative if it wasn't. */
func UnstifleHistory() int {
	return int(C.unstifle_history())
}

/* Begin a session in which the history functions might be used.  This
   just initializes the interactive variables. */
func UsingHistory() {
	C.using_history()
}


// I miss Ruby's attr_reader and attr_accessor

/* The logical `base' of the history array.  It defaults to 1. */
func HistoryBase() int {
	return int(C.history_base)
}

func HistoryBase_(base int) int {
	C.history_base = C.int(base)
	return base
}

/* The number of strings currently stored in the history list. */
func HistoryLength() int {
	return int(C.history_length)
}

/* If HISTORY_STIFLED is non-zero, then this is the maximum number of
   entries to remember. */
func HistoryMaxEntries() int {
	return int(C.history_max_entries)
}

/* The character that represents the start of a history expansion
   request.  This is usually `!'. */
func HistoryExpansionChar() rune {
	return rune(C.history_expansion_char)
}

func HistoryExpansionChar_(c rune) rune {
	C.history_expansion_char = C.char(c) 
	return c
}

/* The character that represents the start of a history expansion
   request.  This is usually `!'. */
func HistorySubstChar() rune {
	return rune(C.history_subst_char)
}

func HistorySubstChar_(c rune) rune {
	C.history_subst_char = C.char(c)
	return c
}

/* During tokenization, if this character is seen as the first character
   of a word, then it, and all subsequent characters upto a newline are
   ignored.  For a Bourne shell, this should be '#'.  Bash special cases
   the interactive comment character to not be a comment delimiter. */
func HistoryCommentChar() rune {
	return rune(C.history_comment_char)
}

func HistoryCommentChar_(c rune) rune {
	C.history_comment_char = C.char(c)
	return c
}




