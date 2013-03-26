package gnureadline
import "testing"

func assert_true(b bool, msg string, t *testing.T) {
	if !b  {
		t.Errorf("failed %s\n", msg)
	}
}
	
func assert_false(b bool, msg string, t *testing.T) {
	if b  {
		t.Errorf("failed %s\n", msg)
	}
}

func assert_equal_int(expect, got int, msg string, t *testing.T) {
	if expect != got  {
		t.Errorf("failed expected %d, got %d: %s\n", 
			expect, got, msg)
	}
}

func assert_equal_rune(expect, got rune, msg string, t *testing.T) {
	if expect != got  {
		t.Errorf("failed expected %c, got %c: %s\n", 
			expect, got, msg)
	}
}

func assert_equal_string(expect, got string, msg string, t *testing.T) {
	if expect != got  {
		t.Errorf("failed expected %s, got %s: %s\n", 
			expect, got, msg)
	}
}

func assert_error(got error, msg string, t *testing.T) {
	if got == nil  {
		t.Errorf("failed expected an error\n")
	}
}

func assert_not_error(got error, msg string, t *testing.T) {
	if got != nil  {
		t.Errorf("failed expected no error\n")
	}
}
