package filextension

import (
    "log"
	"fmt"
	"strings"
    "user/nlplearn/getpdfcontent"
    "user/nlplearn/gettextdocxcontent"
    "user/nlplearn/sentencesegmenter"
)


func getfilextention(filename string){
	switch {
	case strings.HasSuffix(filename, "txt"):
		fmt.Println("this is a txt file")
		txt, err := gettextdocxcontent.getextdocxcontent(filename)
		if err != nil{
			log.Fatal(err)
		}
		txt_contents, err := sentencesegmenter.SegmentSentence(txt)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(txt_contents)
		// TODO write to file txt_contents
	case strings.HasSuffix(filename, "docx"):
		fmt.Println("this is a docx file")
		docx, err := gettextdocxcontent.getextdocxcontent(filename)
		if err != nil{
			log.Fatal(err)
		}

		docx_contents, err := sentencesegmenter.SegmentSentence(docx)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(docx_contents)
		// TODO write to file
	case strings.HasSuffix(filename, "pdf"):
		fmt.Println("this is a pdf file")
		pdf, err := getpdfcontent.Getpdfcontent(filename)
		if err != nil{
			log.Fatal(err)
		}
		pdf_contents, err := sentencesegmenter.SegmentSentence(pdf)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(pdf_contents)
		// TODO write to file
	default:
		fmt.Println("Sorry file type not supported")
	}}