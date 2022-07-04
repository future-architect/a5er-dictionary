package a5er

import (
	"log"
	"strings"
	"unicode/utf8"
)

type Convertor struct {
	missing map[string]struct{}
}

func NewConvertor() *Convertor {
	return &Convertor{make(map[string]struct{})}
}

func (c *Convertor) Logical2Physical(logicalName string, dict *Dictionary) string {
	var converted []string
	r := []rune(logicalName)
	var miss []rune

	for i := 0; i < len(r); {
		var exists bool
		for _, d := range *dict {
			l := utf8.RuneCountInString(d.Key)
			if i+l > len(r) {
				continue
			}
			t := r[i : i+l]

			if string(t) == d.Key {
				converted = append(converted, d.Value)
				i += l
				exists = true
				break
			}
		}

		if !exists {
			miss = append(miss, r[i])
			i++
		}
	}

	if len(miss) > 0 {
		// 中途半端に物理名が設定されることを避けるために、変換に失敗した語句がある場合は物理名を設定しない
		log.Printf("Fail to logical to physical [#%s]. remain [#%s]\n", logicalName, string(miss))
		c.missing[string(miss)] = struct{}{}
		return logicalName
	}

	return strings.Join(converted, "_")
}
