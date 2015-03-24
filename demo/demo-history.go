package main
import (
	"fmt"
	"github.com/rocky/go-gnureadline"
)

func main() {
	fmt.Printf("history_expansion_char is %c\n",
		   gnureadline.HistoryExpansionChar())
	fmt.Printf("history_subst_char is %c\n",
		   gnureadline.HistorySubstChar())
	fmt.Printf("history_comment_char is %c\n",
		   gnureadline.HistoryCommentChar())
	fmt.Printf("History length %d\n",  gnureadline.HistoryLength())
	fmt.Printf("History max entries %d\n",  gnureadline.HistoryMaxEntries())
	fmt.Println("History is stifled", gnureadline.HistoryIsStifled())
	gnureadline.StifleHistory(3)
	fmt.Println("History is stifled", gnureadline.HistoryIsStifled())

}
