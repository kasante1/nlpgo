 
package sentencesegmenter



import (
    "fmt"
    "log"
    "github.com/jdkato/prose/v2"

)


func SegmentSentence(textcontent string)([]string, error){
	doc, err := prose.NewDocument(textcontent)
    if err != nil {
        log.Fatal(err)
    }

	textcontents := make([]string, 0)

        // Iterate over the doc's sentences:
        for _, sent := range doc.Sentences() {
            fmt.Println(sent.Text)
            // Go is an open-source programming language created at Google.
			textcontents = append(textcontents, sent.Text)
        }
	return textcontents, err

}