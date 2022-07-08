package a5er

import (
	"strings"

	"gopkg.in/ini.v1"
)

type Entity struct {
	pName            string
	lName            string
	convertedPName   string
	comment          string
	tableOption      string
	page             string
	left             string
	top              string
	fields           []string
	convertedField   []string
	indexes          []string
	effectMode       string
	color            string
	bkColor          string
	modifiedDateTime string
	position         string
	zOrder           string
}

// 本来全部のキーを取得する必要はないが、
// gopkg.in/ini.v1 の仕様上、複数項目存在する同一キー（Field）に対して
// 想定している更新ができないため、全項目を取得して、本ツールで書き換える
func NewEntity(section *ini.Section) *Entity {
	entity := &Entity{
		pName:            section.Key("PName").String(),
		lName:            section.Key("LName").String(),
		comment:          section.Key("Comment").String(),
		tableOption:      section.Key("TableOption").String(),
		page:             section.Key("Page").String(),
		left:             section.Key("Left").String(),
		top:              section.Key("Top").String(),
		fields:           section.Key("Field").ValueWithShadows(),
		indexes:          section.Key("Index").ValueWithShadows(),
		effectMode:       section.Key("EffectMode").String(),
		color:            section.Key("Color").String(),
		bkColor:          section.Key("BkColor").String(),
		modifiedDateTime: section.Key("ModifiedDateTime").String(),
		position:         section.Key("Position").String(),
		zOrder:           section.Key("ZOrder").String(),
	}
	return entity
}

func (e *Entity) Convert(c *Config, conv *Convertor, dict *Dictionary) {
	// テーブル名
	if convertedPName, ok := conv.Logical2Physical(strings.Trim(e.lName, `"`), dict); ok {
		e.convertedPName = convertedPName
	} else {
		e.convertedPName = strings.Trim(e.pName, `"`)
	}
	if c.TablePlural {
		e.convertedPName = c.PluralClient.Plural(e.convertedPName)
	}

	// カラム名
	fields := e.extractFields()
	for _, field := range fields {
		if convertedFieldPhysicalName, ok := conv.Logical2Physical(strings.Trim(field.logicalName, `"`), dict); ok {
			field.convertedPhysicalName = convertedFieldPhysicalName
		} else {
			field.convertedPhysicalName = strings.Trim(field.physicalName, `"`)
		}
		e.convertedField = append(e.convertedField, field.String())
	}
}

func (e *Entity) extractFields() []Field {
	var fs []Field
	for _, f := range e.fields {
		ss := strings.SplitN(f, ",", 3)
		fs = append(fs, Field{
			logicalName:  ss[0],
			physicalName: ss[1],
			other:        ss[2],
		})
	}
	return fs
}

type Field struct {
	logicalName           string
	convertedPhysicalName string
	physicalName          string
	other                 string
}

func (f Field) String() string {
	pname := f.physicalName
	if f.convertedPhysicalName != "" {
		pname = `"` + f.convertedPhysicalName + `"`
	}
	return strings.Join([]string{f.logicalName, pname, f.other}, ",")
}

func clearEntitySectionKey(section *ini.Section) {
	section.DeleteKey("PName")
	section.DeleteKey("LName")
	section.DeleteKey("Comment")
	section.DeleteKey("TableOption")
	section.DeleteKey("Page")
	section.DeleteKey("Left")
	section.DeleteKey("Top")
	section.DeleteKey("Field")
	section.DeleteKey("Index")
	section.DeleteKey("EffectMode")
	section.DeleteKey("Color")
	section.DeleteKey("BkColor")
	section.DeleteKey("ModifiedDateTime")
	section.DeleteKey("Position")
	section.DeleteKey("ZOrder")
}

func (e *Entity) writeSectionKey(section *ini.Section) {
	section.Key("PName").SetValue(e.convertedPName)
	section.Key("LName").SetValue(e.lName)
	section.Key("Comment").SetValue(e.comment)
	section.Key("TableOption").SetValue(e.tableOption)
	section.Key("Page").SetValue(e.page)
	section.Key("Left").SetValue(e.left)
	section.Key("Top").SetValue(e.top)
	for _, field := range e.convertedField {
		section.Key("Field").AddShadow(field)
	}
	for _, index := range e.indexes {
		section.Key("Index").AddShadow(index)
	}
	section.Key("Field").ValueWithShadows()
	section.Key("EffectMode").SetValue(e.effectMode)
	section.Key("Color").SetValue(e.color)
	section.Key("BkColor").SetValue(e.bkColor)
	section.Key("ModifiedDateTime").SetValue(e.modifiedDateTime)
	section.Key("Position").SetValue(e.position)
	section.Key("ZOrder").SetValue(e.zOrder)
}

type Relation struct {
	fields1         string
	convertedField1 string
	fields2         string
	convertedField2 string
}

func NewRelation(section *ini.Section) *Relation {
	return &Relation{
		fields1: section.Key("Fields1").String(),
		fields2: section.Key("Fields2").String(),
	}
}

func (r *Relation) Convert(c *Config, conv *Convertor, dict *Dictionary) {
	r.convertedField1 = convertRelationField(c, r.fields1, conv, dict)
	r.convertedField2 = convertRelationField(c, r.fields2, conv, dict)
}

func convertRelationField(c *Config, field string, conv *Convertor, dict *Dictionary) string {
	ss := strings.Split(strings.Trim(field, `"`), ",")
	var convertedField []string
	for _, s := range ss {
		if cf, ok := conv.Logical2Physical(s, dict); ok {
			convertedField = append(convertedField, cf)
		} else {
			convertedField = append(convertedField, s)
		}
	}
	return strings.Join(convertedField, ",")
}

func (r *Relation) writeSectionKey(section *ini.Section) {
	section.Key("Fields1").SetValue(r.convertedField1)
	section.Key("Fields2").SetValue(r.convertedField2)
}
