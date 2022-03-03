package main

import "fmt"
func HelperFunctions(keyword string) (string, error){
	var key_word string = keyword + " "+ "this is the keyword"
	return key_word, fmt.Errorf(key_word)
}
// import (
//     "fmt"
//     "log"

//     "github.com/jdkato/prose/v2"
//     "github.com/lu4p/cat"
// )

// func main() {
//     // Create a new document with the default configuration:
//     doc, err := prose.NewDocument(`Then at another part of the
//     judgment he stated as follows; “It has been earlier held that the security created with
//     the account receivable is not enforceable against the 3rd defendant. In accordance with
//     section 25(3) of the Borrowers and Lenders Act, 2008 (Act 773) though the said
//     security is not enforceable, the money which was secured shall immediately become
//     payable...” But what must be noted is that John Oseku Ankrah did not guarantee the
//     loans that were extended to 2 nd defendant in this case. It was the 1st defendant who
//     guaranteed the loans. Sections 1 and 2 of Act 773 gives the matters the Act applies to
//     as follows;
//     Section 1. Application
//     (1) This Act applies to
//     (a) a credit agreement between parties who deal at arm’s length except;- (i)
//     a credit agreement covering an amount of less than one hundred Ghana
//     Cedis (GH¢ 100) or an amount determined by the Bank of Ghana; (ii) any
//     other credit agreement exempted by the Bank of Ghana by notice under this
//     Act.
//     (b) a credit guarantee where the guarantee is only in respect of a credit
//     facility or credit transaction covered by this Act.
//     2. Meaning of credit agreement
//     For the purpose of this Act a credit agreement is an agreement in the nature
//     of a credit facility, a credit transaction, a credit guarantee or any
//     combination of these.
//     A credit guarantee as stated under section 5 of Act 773 is; “an agreement is a credit
//     guarantee if in that agreement a third party undertakes to satisfy on demand
//     an obligation of a borrower in a credit facility or credit transaction to which
//     this Act applies.” (emphasis supplied).
//     `)
//     if err != nil {
//         log.Fatal(err)
//     }

//     // Iterate over the doc's tokens:
//     for _, tok := range doc.Tokens() {
//         fmt.Println(tok.Text, tok.Tag, tok.Label)
//         // Go NNP B-GPE
//         // is VBZ O
//         // an DT O
//         // ...
//     }

//     // Iterate over the doc's named-entities:
//     for _, ent := range doc.Entities() {
//         fmt.Println(ent.Text, ent.Label)
//         // Go GPE
//         // Google GPE
//     }

//     // Iterate over the doc's sentences:
//     for _, sent := range doc.Sentences() {
//         fmt.Println(sent.Text)
//         fmt.Println("------------------------")
//         // Go is an open-source programming language created at Google.


//     txt, err := cat.File("filename")
//     fmt.Println(txt)
//     if err != nil{
//         log.Fatal(err)
//     }
//     }
// }