package main 

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "regexp"
    "github.com/kljensen/snowball"
    . "github.com/jbrukh/bayesian"
    "github.com/sjwhitworth/golearn/base"
    "github.com/sjwhitworth/golearn/evaluation"

)

// readLines reads a whole file into memory 
// and returns a slice of its lines.

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}

func tokLines(lines []string) [][]string {
	var toklines [][]string

	for _, line := range lines {
		toklines = append(toklines, strings.Fields(line))
	}
    return toklines

}


func main() {
	lines, err := readLines("./smsspamcollection/SMSSpamCollection")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//for i, line := range lines {
	//	fmt.Println(i, line)
	//}

    for i := 0; i < 5; i++ {
    	fmt.Println(lines[i])
    }

    fmt.Println("--------------------")
	

    //shorten toklines for print purpose
	toklines := tokLines(lines)
	//fmt.Println("%q", toklines)

	var labels []string
	var messages [][]string
	for _, message := range toklines {
		labels = append(labels, message[0])
		messages = append(messages, message[1:])
	}

	fmt.Println("%q - %q", labels[:5], messages[:5])
	fmt.Println("--------------------")


	var stemmed_messages [][]string
	var stemmed_message []string
	for _, message := range messages {
		stemmed_message = []string{}
		for _, token := range message {
			re := regexp.MustCompile("[^A-Za-z0-9 ]")
			stemmed, err := snowball.Stem(token, "english", true)
			if err == nil {
				stemmed_message = append(stemmed_message, re.ReplaceAllLiteralString(stemmed, ""))
			}
		}
		stemmed_messages = append(stemmed_messages, stemmed_message)
		//fmt.Println("%q", stemmed_message)
	}

    fmt.Println("%q", stemmed_messages[:5])

    //Do a training-test split
    trainData, testData := base.InstancesTrainTestSplit(stemmed_messages, 0.25)


    //build the labels
    
    const (
    	spam Class = "spam"
    	ham Class = "ham"
    )


    var clabels []Class

    for _, l := range labels {
    	if l == "spam" {
    		clabels = append(clabels, spam)
    	} else if l == "ham" {
    		clabels = append(clabels, ham)
    	}
    }

    // train the classifier
    classifier := NewClassifier(ham, spam)
	for i := 0; i < len(clabels) - 500; i++  {
		classifier.Learn(stemmed_messages[i], clabels[i])
	}


	for i:= len(clabels) - 500; i < len(clabels); i++ {
		fmt.Print(labels[i])
		probs, likely, _ := classifier.ProbScores(stemmed_messages[i])
		fmt.Println(" --- probs: %q, classfied to: %q", probs, likely)
	}
    

	//fmt.Println("%q, %q; %q, %q", scores, likely1, probs, likely2)
	//fmt.Println("%q", len(labels))

}