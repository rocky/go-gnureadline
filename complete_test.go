package gnureadline
import "testing"

func TestComplete(t *testing.T) {
	completer_characters := "X"
	CompleterWordBreakCharacters_(completer_characters)
	assert_equal_string(completer_characters,
		CompleterWordBreakCharacters(), 
		"Completer characters should match what was just set", t)
}
