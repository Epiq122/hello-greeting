package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, eg. en, es, fr, etc.")
	flag.Parse()
	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// represents the languages code
type language string

// phrasebook holds the translations for a given language.
var phrasebook = map[language]string{
	"el": "Γειά σου Κόσμε!", // Greek
	"en": "Hello, World!",   // English
	"es": "¡Hola Mundo!",    // Spanish
	"fr": "Bonjour, Monde!", // French
	"it": "Ciao, Mondo!",    // Italian
	"la": "Salve, Orbis!",   // Latin
	"pt": "Olá Mundo!",      // Portuguese
	"ru": "Привет, мир!",    // Russian
	"tr": "Selam Dünya!",    // Turkish
	"ur": "ہیلو دنیا",       // Urdu
}

// greet says hello to the world in a given language.
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %s", l)
	}
	return greeting
}
