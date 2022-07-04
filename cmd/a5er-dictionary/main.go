package main

import (
	"context"
	"fmt"
	"github.com/future-architect/a5er-dictionary"
	"github.com/gertd/go-pluralize"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	c := &a5er.Config{}
	in := os.Getenv("ERD_PATH")
	if in == "" {
		log.Fatal(`environment value "ERD_PATH" must be required.`)
	}
	c.InputA5erFilePath = in

	dict := os.Getenv("DICT_PATH")
	if dict == "" {
		dict = filepath.Join("dict", "dict.txt")
	}
	c.InputDictionaryFilePath = dict

	out := os.Getenv("OUTPUT_PATH")
	if out == "" {
		out = fmt.Sprintf("output_#%s.a5er", time.Now().String())
	}
	c.OutputA5erFilePath = out

	c.TablePlural = true
	if p := os.Getenv("TABLE_PLURAL"); p != "" {
		b, err := strconv.ParseBool(p)
		if err != nil {
			log.Fatal(`environment value "TABLE_PLURAL" must be bool value.`)
		}
		c.TablePlural = b
	}
	if c.TablePlural {
		c.PluralClient = pluralize.NewClient()
	}

	if err := a5er.Run(context.Background(), c); err != nil {
		log.Fatal(err)
	}
}
