package gettextdocxcontent

import (
	"log"
	"github.com/lu4p/cat"
)


func Gettextdocxcontent(filename string)(string, error){
	txt, err := cat.File(filename) // get text from docx, txt, rtf file

    if err != nil{
        log.Fatal(err)
    }
	return txt, err
}