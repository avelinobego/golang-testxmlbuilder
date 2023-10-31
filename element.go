package element

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func Make(name string) *Element {
	return &Element{root: name, comment: nil}
}

type Element struct {
	root    string
	value   string
	space   string
	attribs []xml.Attr
	comment xml.Comment
	ommit   bool
}

func (el *Element) Ommit(ommit bool) {
	el.ommit = ommit
}

func (el *Element) Comment(text string) {
	el.comment = xml.Comment(text)
}

func (el *Element) Space(val string) {
	el.space = val
}

func (el *Element) Value(val string, attribs ...map[string]interface{}) {
	internal := make(map[string]interface{})
	for _, np := range attribs {
		for k, v := range np {
			internal[k] = v
		}
	}

	el.value = val
	for k, v := range internal {
		el.attribs = append(el.attribs, xml.Attr{
			Name:  xml.Name{Space: "", Local: k},
			Value: fmt.Sprintf("%v", v)})
	}
}

func (el *Element) ToXml() (result string, err error) {

	if el.ommit && el.value == "" {
		return
	}

	name := xml.Name{Local: el.root, Space: el.space}
	start := xml.StartElement{Name: name, Attr: el.attribs}

	wr := &bytes.Buffer{}
	enc := xml.NewEncoder(wr)

	if el.comment != nil {
		err = enc.EncodeToken(el.comment)
		if err != nil {
			return
		}
	}

	err = enc.EncodeToken(start)
	if err != nil {
		return
	}

	enc.Flush()

	wr.WriteString(fmt.Sprintf("%v", el.value))

	err = enc.EncodeToken(start.End())
	if err != nil {
		return
	}
	enc.Close()

	result = wr.String()
	return
}

func (el *Element) String() string {
	result, _ := el.ToXml()
	return result
}
