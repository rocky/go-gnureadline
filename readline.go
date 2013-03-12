/* Copyright (C) 2013 Rocky Bernstein

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

// This file contains things from readline/readline.h

package gnureadline

/*
#cgo LDFLAGS: -lreadline
#include <stdio.h>
#include <readline/readline.h>
#include <readline/history.h>  // for readline() history parameter
#include <stdlib.h> // for free()
*/
import "C"
import "unsafe"

/*
Readline will read a line from the terminal and return it, using prompt
as a prompt.  If prompt the  empty  string,  no  prompt  is
issued.  The  line  returned  has  the  final  newline
removed, so only the text of the line remains.

Readline offers editing capabilities while the user is entering the
line.  By default, the line editing commands are similar to those of
emacs.  A vi-style line editing interface is also available.

*/
func Readline(prompt string add_history ... bool) string {
	c_prompt := C.CString(prompt)
	defer C.free(unsafe.Pointer(c_prompt))
	c_line := C.readline(c_prompt)
	defer C.free(unsafe.Pointer(c_line))
	if len(add_history) > 0 && add_history[0] == true {
		C.add_history(c_line)
	}
		
	return C.GoString(c_line)
}

/*
     Modify the terminal settings for Readline's use, so `readline()'
     can read a single character at a time from the keyboard.  The
     META_FLAG argument should be non-zero if Readline should read
     eight-bit input.
*/
func Rl_prep_terminal(meta_flag int) {
	C.rl_prep_terminal(C.int(meta_flag))
}

/*
     Undo the effects of `rl_prep_terminal()', leaving the terminal in
     the state in which it was before the most recent call to
     `rl_prep_terminal()'.
*/
func Rl_deprep_terminal() {
	C.rl_deprep_terminal()
}

/*
     Ring the terminal bell, obeying the setting of `bell-style'.
*/
func Rl_ding () int {
	return int(C.rl_ding())
}

/*
     Reinitialize Readline's idea of the terminal settings using
     TERMINAL_NAME as the terminal type (e.g., `vt100').  If
     TERMINAL_NAME is the empty string, the value of the `TERM' environment
     variable is used.
*/
func Rl_reset_terminal(terminal_name string) int {
	if len(terminal_name) == 0 {
		return int(C.rl_reset_terminal(nil))
	}
	c_terminal_name := C.CString(terminal_name)
	defer C.free(unsafe.Pointer(c_terminal_name))
	return int(C.rl_reset_terminal(c_terminal_name))
}

/*  
 Update Readline's internal screen size by reading values from the
 kernel.
*/
func Rl_resize_terminal() {
	C.rl_resize_terminal()
}

type EditingMode int

const (
	Vi    EditingMode = 0
	Emacs EditingMode = 1
)

/* Says which editing mode readline is currently using: Emacs or Vi */
func Rl_editing_mode() EditingMode {
	return EditingMode(C.rl_editing_mode)
}

// I miss Ruby's attr_reader

/* 
 True if this is real GNU readline. (It's probably true here.)
*/
func Rl_gnu_readline_p() bool {
	if (int(C.rl_gnu_readline_p) == 0) {
		return false
	}
	return true
}

/*
     If non-zero, Readline gives values found in the `LINES' and
     `COLUMNS' environment variables greater precedence than values
     fetched from the kernel when computing the screen dimensions.
*/
func Rl_prefer_env_winsize() int {
	return int(C.rl_prefer_env_winsize)
}

/* 
 Read keybindings and variable assignments from FILENAME 
*/
func Rl_read_init_file(filename string) int {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	return int(C.rl_read_init_file(c_filename))
}

/* The version number of this revision of the library. e.g. "6.2" */
func Rl_readline_library_version() string {
	return C.GoString(C.rl_library_version)
}

/*
     This variable is set to a unique name by each application using
     Readline.  The value allows conditional parsing of the inputrc file
*/
func Rl_readline_name() string {
	return C.GoString(C.rl_readline_name)
}

/*
     The terminal type, used for initialization.  If not set by the
     application, Readline sets this to the value of the `TERM'
     environment variable the first time it is called.
*/
func Rl_readline_terminal_name() string {
	return C.GoString(C.rl_terminal_name)
}

/*
     Return a string representing the value of the Readline variable
     VARIABLE.  For boolean variables, this string is either `on' or
     `off'.
*/
func Rl_variable_value(variable string) string {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))
	return C.GoString(C.rl_variable_value(c_variable))
}

/*
     An integer encoding the current version of the library.  The
     encoding is of the form 0xMMMM, where MM is the two-digit major
     version number, and MM is the two-digit minor version number.  For
     example, for Readline-4.2, `rl_readline_version' would have the
     value 0x0402.
*/
func Rl_readline_version() int {
	return int(C.rl_readline_version)
}

/* Insert or overwrite mode for emacs mode.  1 means insert mode; 0 means
   overwrite mode.  Reset to insert mode on each input line. */
func Rl_insert_mode() int {
	return int(C.rl_insert_mode)
}

// I miss Ruby's attr_accessor

/* Set editing mode readline is currently using.  1 means emacs mode;
   0 means vi mode. */
func Rl_editing_mode_set(new_value EditingMode) EditingMode {
	C.rl_editing_mode = C.int(new_value)
	return EditingMode(C.rl_editing_mode)
}

/* Set insert or overwrite mode for emacs mode.  1 means insert mode; 0 means
   overwrite mode.  Reset to insert mode on each input line. */
func Rl_insert_mode_set(new_value int) int {
	C.rl_insert_mode = C.int(new_value)
	return int(C.rl_insert_mode)
}

