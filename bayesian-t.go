package main
import (
	. "github.com/jbrukh/bayesian"
	"fmt"

)

const (
    Good Class = "Good"
    Bad Class = "Bad"
)

func main() {
	classifier := NewClassifier(Good, Bad)
	goodStuff := []string{"tall", "rich", "handsome"}
	badStuff  := []string{"poor", "smelly", "ugly"}
	goodStuff1 := []string{"bueno", "nueve", "pretty"}
	classifier.Learn(goodStuff, Good)
	classifier.Learn(badStuff,  Bad)
	classifier.Learn(goodStuff1, Good)

	scores, likely1, _ := classifier.LogScores([]string{"bueno", "random"})
	probs, likely2, _ := classifier.ProbScores([]string{"random", "poor"})

	fmt.Println("%q, %q; %q, %q", scores, likely1, probs, likely2)
}
