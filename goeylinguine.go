// goeylinguine is a package that uses data from
// github.com/github/linguist data to return a file's lexer type.
package goeylinguine

import (
	"bufio"
	"encoding/gob"
	"os"
	"path/filepath"
)

type Language struct {
	Language   string
	Extensions []string
	FileNames  []string
	Color      string
	Lexer      string
	Type       string
}
type languages struct {
	Languages []Language
}
var langs languages

func init() {
	langs = languages{}
	fi, _ := os.Open(os.Getenv("GOPATH") + "/src/github.com/sevki/goeylinguine/languages.gob")
	dec := gob.NewDecoder(bufio.NewReader(fi))
	dec.Decode(&langs)
}
// Given a file name returns it's linguist information, color, lexer,
// type, etc.
func GetLanguageFromFileName(fname string) *Language {
	return getFileLanguage(fname)
}
// Given os.File returns it's linguist information, usefull when
// iterating trough a file range in a directory.
func GetFileLanguage(f os.File) *Language {
	fstat, _ := f.Stat()
	return getFileLanguage(fstat.Name())
}

func getFileLanguage(fname string) *Language {

	ext := filepath.Ext(fname)
	for _, lang := range langs.Languages {
		for _, fn := range lang.FileNames {
			if fname == fn {
				return &lang
			}
		}
		for _, xt := range lang.Extensions {
			if ext == xt {
				return &lang
			}
		}
	}
	return &Language{Language: "Text", Color: "#ccc"}
}
