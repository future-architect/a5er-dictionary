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

func (c *Convertor) Logical2Physical(logicalName string, dict *Dictionary) (string, bool) {
	var converted []string
	r := []rune(logicalName)
	// マッチしない語句が複数ある場合は、カンマ区切りでそれぞれの単語を出力する
	var misses []string
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
				if len(miss) > 0 {
					misses = append(misses, string(miss))
					miss = nil
				}
				break
			}
		}

		if !exists {
			miss = append(miss, r[i])
			i++
		}
	}

	if len(miss) > 0 {
		misses = append(misses, string(miss))
	}

	if len(misses) > 0 {
		// 中途半端に物理名が設定されることを避けるために、変換に失敗した語句がある場合は物理名を設定しない
		log.Printf("Fail to logical to physical [#%s]. remain [#%s]\n", logicalName, strings.Join(misses, ","))
		c.missing[string(miss)] = struct{}{}
		return "", false
	}

	return strings.Join(converted, "_"), true
}

// IsConverted はlogicalNameがすでに論物変換済かどうか判定します
func (c *Convertor) IsConverted(logicalName string, dict *Dictionary) bool {
	ss := strings.Split(logicalName, "_")

	startIdx := 0
	lastIdx := len(ss)
	// 論物変換はマッチング候補が複数ある場合はなるべく長い単語にマッチングする
	// よって変換判定もなるべく長い値から順に辞書に登録されているか判定する
	for 0 < lastIdx && startIdx < len(ss) && startIdx <= lastIdx {
		v := strings.Join(ss[startIdx:lastIdx], "_")
		if dict.containsValue(v) {
			startIdx = lastIdx
			lastIdx = len(ss)
			continue
		}
		lastIdx--
	}
	return startIdx == len(ss)
}
