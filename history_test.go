package gnureadline
import "testing"

func TestReadlineHistory(t *testing.T) {
	assert_equal_int(0, HistoryLength(), "History length should be 0", t)
	assert_equal_int(0, HistoryMaxEntries(), "History max entries be 0", t)

	StifleHistory(3)
	assert_true(HistoryIsStifled(), "History should be stifled", t)
	UnstifleHistory()
	assert_false(HistoryIsStifled(), "History should not be stifled", t)

	subst_char := HistorySubstChar()
        test_char  := 'X'
	if subst_char == test_char { test_char = 'Y' }
	assert_equal_rune(test_char, HistorySubstChar_(test_char),
		"Set HistorySubstChar", t)

        test_char = 'Z'
	comment_char := HistoryCommentChar()
	if comment_char == test_char { test_char = 'A' }
	assert_equal_rune(test_char, HistoryCommentChar_(test_char),
		"Set HistorySubstChar", t)

	AddHistory("testing")
	assert_equal_int(1, HistoryLength(), "History length should be 1", t)
	if err := RemoveHistory(0); err != nil {
		assert_true(false, "Got error in removing history item 0", t)
	}
	if err := RemoveHistory(0); err == nil {
		assert_true(false, "Show have gotten an error in removing empty history item 0", t)
	}
	ClearHistory()
	assert_equal_int(0, HistoryLength(), "History length should be back at 0 after clearing",
		t)

}
