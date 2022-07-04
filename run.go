package a5er

import (
	"context"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func Run(ctx context.Context, c *Config) error {

	// CSV辞書の読込
	dict := new(Dictionary)

	if err := dict.LoadCSV(c.InputDictionaryFilePath); err != nil {
		return fmt.Errorf("load csv, path = %s", c.InputDictionaryFilePath)
	}

	// A5erファイルの読込
	ini.PrettyFormat = false
	opt := ini.LoadOptions{
		SpaceBeforeInlineComment: true,
		AllowShadows:             true,
		AllowNonUniqueSections:   true,
	}
	a5er, err := ini.LoadSources(opt, c.InputA5erFilePath)
	if err != nil {
		return fmt.Errorf("load a5er file, path = %s", c.InputA5erFilePath)
	}

	// 論物変換
	conv := NewConvertor()
	for _, section := range a5er.Sections() {
		switch section.Name() {
		case "Entity":
			entity := NewEntity(section)
			entity.Convert(c, conv, dict)
			clearEntitySectionKey(section)
			entity.writeSectionKey(section)
		case "Relation":
			entity := NewRelation(section)
			entity.Convert(c, conv, dict)
			entity.writeSectionKey(section)
		}
	}

	// ファイル出力
	out, err := os.Create(c.OutputA5erFilePath)
	if err != nil {
		return fmt.Errorf("create converted a5er file, path = %s", c.OutputA5erFilePath)
	}
	defer out.Close()

	if _, err := a5er.WriteTo(out); err != nil {
		return err
	}

	return nil
}
