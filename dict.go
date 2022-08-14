package a5er

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode/utf8"
)

type Dictionary []*dict

func (d Dictionary) String() string {
	var o string
	for _, dd := range d {
		o += fmt.Sprintf("[%s, %s]", dd.Key, dd.Value)
	}
	return o
}

type dict struct {
	Key, Value string
}

func (d *Dictionary) LoadCSV(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		key, value := rec[0], rec[1]
		*d = append(*d, &dict{key, value})
	}

	// より長くマッチした単語を優先して置換対象とするため、
	// 日本語文字列の降順でソート
	d.sort()

	return nil
}

func (d *Dictionary) sort() {
	sort.SliceStable(*d, func(i, j int) bool {
		return utf8.RuneCountInString((*d)[i].Key) > utf8.RuneCountInString((*d)[j].Key)
	})
}

func (d *Dictionary) containsValue(value string) bool {
	for _, dd := range *d {
		if value == dd.Value {
			return true
		}
	}
	return false
}
