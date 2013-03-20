package gnureadline
import "testing"

func TestReadlineHistory(t *testing.T) {
	assert_true(HistoryLength() == 0, "History length should be 0", t)
	assert_true(HistoryMaxEntries() == 0, "History max entries be 0", t)

	StifleHistory(3)
	assert_true(HistoryIsStifled(), "History should be stifled", t)
	UnstifleHistory()
	assert_false(HistoryIsStifled(), "History should not be stifled", t)
}
