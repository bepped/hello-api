package translation_test

import (
	"testing"
	"github.com/bepped/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	// Arrange
	word := "hello"
	language := "english"

	// Act 
	res := translation.Translate(word, language)

	// Arrange
	if res != "hello" {
		t.Errorf(`expected "hello" but received "%s`, %res)
	}
}