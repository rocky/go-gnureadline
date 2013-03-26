package main
import ( . "gnureadline"; "fmt")

func main() {
	fmt.Printf("history_expansion_char is %c\n", 
		   HistoryExpansionChar())
	fmt.Printf("history_subst_char is %c\n", 
		   HistorySubstChar())
	fmt.Printf("history_comment_char is %c\n", 
		   HistoryCommentChar())
	fmt.Printf("History length %d\n",  HistoryLength())
	fmt.Printf("History max entries %d\n",  HistoryMaxEntries())
	fmt.Println("History is stifled", HistoryIsStifled())
	StifleHistory(3)
	fmt.Println("History is stifled", HistoryIsStifled())

}
