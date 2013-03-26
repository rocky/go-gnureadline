/* 
   GNU Readline key binding and startup file support

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

/* Thanks to Sebastien Binet from which these routines were started from. */

package gnureadline
/*
#cgo darwin CFLAGS: -I/opt/local/include
#cgo darwin LDFLAGS: -L/opt/local/lib
#cgo LDFLAGS: -lreadline
#include <stdio.h>
#include <readline/readline.h>
#include <stdlib.h> // for free()
*/
import "C"
import "unsafe"
import "syscall"

/* 
func GetKeymap() Keymap {
	return C.GoString(C.rl_get_keymap())
}
*/

/**** 
Should we include? 
*****/
func GetKeymapNameFromEditMode() string {
	return C.GoString(C.rl_get_keymap_name_from_edit_mode())
}

func SetKeymapNameFromEditMode() {
	C.rl_set_keymap_from_edit_mode()
}

/* 
 Read the binding command from STRING and perform it.
 A key binding command looks like: Keyname: function-name\0,
 a variable binding command looks like: set variable value.
 A new-style keybinding looks like "\C-x\C-x": exchange-point-and-mark. 
 
 True is returned if there wasn't an error, false otherwise.
 */
func ParseAndBind(s string) bool {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	success := C.rl_parse_and_bind(p)
	if (success == 0) { return true }
	return false
}

/* 
 Parse a readline initialization file. The default filename is the
 last filename used.

 Do key bindings from a file.  If FILENAME is NULL it defaults
 to the first non-null filename from this list:
 1. the filename used for the previous call
 2. the value of the shell variable `INPUTRC'
 3. ~/.inputrc
 4. /etc/inputrc
 If the file existed and could be opened and read, 0 is returned,
   otherwise errno is returned. */

func ReadInitFile(filename string) error {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	errno := C.rl_read_init_file(c_filename)
	if errno == 0 {	return nil }
	return syscall.Errno(errno)
}

/* Re-read the current keybindings file. */
func ReReadInitFile() error {
	/* readline has these two bogus "count" and "ignore" 
	 parameters we have to supply. */
	errno := C.rl_re_read_init_file(-1, -1)
	if errno == 0 {	return nil }
	return syscall.Errno(errno)
}

