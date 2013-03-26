package gnureadline
import "testing"

func TestBind(t *testing.T) {
	assert_not_error(ReadInitFile("./data/undo.inputrc"), 
		"Should be able to read ./data/undo.inputrc", t)
	assert_error(ReadInitFile("bogus!"), 
		"Should not be able to read ./bogus!", t)
	ReReadInitFile()

}
