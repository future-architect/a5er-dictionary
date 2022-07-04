package a5er

import "github.com/gertd/go-pluralize"

// Config holds information about the configuration of a5er-dictionary.
type Config struct {
	// input
	InputA5erFilePath       string
	InputDictionaryFilePath string

	// output
	OutputA5erFilePath string

	// plural
	TablePlural  bool
	PluralClient *pluralize.Client
}
