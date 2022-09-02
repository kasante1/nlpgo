package main

import (
	"bufio"
	"fmt"
	"os"
	"user/nlplearn/filextension"
    "user/nlplearn/checkfile"
)

func main(){
    fmt.Println("[Text Normalizer!]")
    fmt.Println("Enter the file directory :")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    filename := scanner.Text()
    if len(filename) != 0 && checkfile(filename) == true{
        filename_extension := filextension.getfilextention(filename)
    }else{

    }
        
}