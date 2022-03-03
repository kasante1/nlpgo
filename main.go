package main

import (
    "log"
	"bufio"
	"fmt"
	"os"
	"strings"
    "user/nlplearn/getpdfcontent"
    "user/nlplearn/gettextdocxcontent"
    "user/nlplearn/sentencesegmenter"
)

func main(){
    fmt.Println("Enter the file directory :")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    filename := scanner.Text()
    if len(filename) != 0{
        if strings.HasSuffix(filename, "txt"){
            fmt.Println("this is a txt file")
            txt, err := gettextdocxcontent.Gettextdocxcontent(filename)
            if err != nil{
                log.Fatal(err)
            }
            txt_contents, err := sentencesegmenter.SegmentSentence(txt)
            if err != nil{
                log.Fatal(err)
            }
            fmt.Println(txt_contents)
        }else if strings.HasSuffix(filename, "docx"){
            fmt.Println("this is a docx file")
            docx, err := gettextdocxcontent.Gettextdocxcontent(filename)
            if err != nil{
                log.Fatal(err)
            }

            docx_contents, err := sentencesegmenter.SegmentSentence(docx)
            if err != nil{
                log.Fatal(err)
            }
            fmt.Println(docx_contents)
        }else if strings.HasSuffix(filename, "pdf"){
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
        }
    }

}