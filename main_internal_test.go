package main

import "testing"

func Example_main() {
	main()
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"Greek": {
			lang: "el",
			want: "Γειά σου Κόσμε!",
		},

		"English": {
			lang: "en",
			want: "Hello, World!",
		},

		"Spanish": {
			lang: "es",
			want: "¡Hola Mundo!",
		},

		"French": {
			lang: "fr",
			want: "Bonjour, Monde!",
		},

		"Italian": {
			lang: "it",
			want: "Ciao, Mondo!",
		},

		"Latin": {
			lang: "la",
			want: "Salve, Orbis!",
		},

		"Portuguese": {
			lang: "pt",
			want: "Olá Mundo!",
		},

		"Russian": {
			lang: "ru",
			want: "Привет, мир!",
		},

		"Turkish": {
			lang: "tr",
			want: "Selam Dünya!",
		},
	}
	// range over the map
	for name, tc := range tests {
		// run the test
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}

// func TestGreet_English(t *testing.T) {
// 	lang := language("en")
// 	want := "Hello, World!"
// 	got := greet(lang)

// 	if got != want {
// 		// mark this as failed
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }
// func TestGreet_French(t *testing.T) {
// 	lang := language("fr")
// 	want := "Bonjour, Monde!"
// 	got := greet(lang)

// 	if got != want {
// 		// mark this as failed
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

// func TestGreet_Akkadian(t *testing.T) {
// 	// not implemented yet
// 	lang := language("akk")
// 	want := ""
// 	got := greet(lang)

// 	if got != want {
// 		// mark this as failed
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }
