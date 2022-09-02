// text to lower case
// remove stop words
// replace contrations with the full words
// remove punctuation marks special, accented characters
// text correction

package textnormalizer

import "strings"


func ToLowerCase(text_content string) string{
	var textToLowerCase string = strings.ToLower(text_content)

	return textToLowerCase
}


func RemoveStopWords(text_content string)string{
	return text_content
}

func ExpandContrations(text_contents string)string{
	return text_contents
}


func RemovePunctuationSpecialCharacters(text_contents string)string{
	return text_contents
}