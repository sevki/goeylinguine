// Package goeylinguine is a package that uses data from
// github.com/github/linguist data to return a file's lexer type.
package goeylinguine

import (
	"bufio"
	"encoding/gob"
	"os"
	"path/filepath"
)

// Language information for linguist
type Language struct {
	Language   string
	Extensions []string
	FileNames  []string
	Color      string
	Lexer      string
	Type       string
}
type Languages struct {
	Languages []Language
}

var langs Languages

func init() {
	langs = Languages{}
	fi, _ := os.Open(os.Getenv("GOPATH") + "/src/github.com/sevki/goeylinguine/languages.gob")
	dec := gob.NewDecoder(bufio.NewReader(fi))
	dec.Decode(&langs)
}

// GetLanguageFromFileName will return it's linguist information,
// color, lexer, type, etc of a given file name.
func GetLanguageFromFileName(fname string) *Language {
	return getFileLanguage(fname)
}

// GetFileLanguage will return language information for the given
// os.File, useful for iterating trough readdir().
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

// GetLanguage will return language information for the given
// language.
func GetLanguage(name string) *Language {
	for _, lang := range langs.Languages {
		if lang.Language == name {
			return lang
		}
	}
	return &Language{Language: "Text", Color: "#ccc"}
}
