package element

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func Make(name string) *Element {
	return &Element{root: name, comment: nil}
}

type ChildCallable func(*Element)

type Elements []*Element
type Attribs map[string]*xml.Attr

type Element struct {
	root    string
	value   any
	cdata   bool
	child   Elements
	space   string
	attribs Attribs
	comment xml.Comment
	ommit   bool
}

func (el *Element) Name() string {
	return el.root
}

func (el *Element) Child(child ...*Element) *Element {
	el.child = append(el.child, child...)
	return el
}

func (el *Element) Ommit(ommit bool) *Element {
	el.ommit = ommit
	return el
}

func (el *Element) Comment(text string) *Element {
	el.comment = xml.Comment(text)
	return el
}

func (el *Element) Space(val string) *Element {
	el.space = val
	return el
}

func (el *Element) Cdata(val bool) *Element {
	el.cdata = val
	return el
}

func (el *Element) Attrib(key string, val any) *Element {
	if el.attribs == nil {
		el.attribs = make(Attribs)
	}
	if v, ok := el.attribs[key]; ok {
		v.Value = fmt.Sprintf("%v", val)
	} else {
		el.attribs[key] = &xml.Attr{
			Name:  xml.Name{Space: "", Local: key},
			Value: fmt.Sprintf("%v", val)}
	}

	return el
}

func (el *Element) Val(val any) *Element {
	el.value = val
	return el
}

func (el Element) GetVal() any {
	return el.value
}

func (el *Element) Values(tag string, value any) *Element {
	temp := Make(tag)
	temp.value = value
	el.child = append(el.child, temp)
	return el
}

func (el *Element) Through(f ChildCallable) (result *Element) {
	result = el
	f(el)
	for _, c := range el.child {
		if c.child != nil {
			c.Through(f)
		} else {
			f(c)
		}
	}
	return
}

func (el *Element) ToXml() (result string, err error) {

	if el.ommit && el.value == nil {
		return
	}

	name := xml.Name{Local: el.root, Space: el.space}

	var attribs []xml.Attr
	if el.attribs != nil {
		for _, v := range el.attribs {
			attribs = append(attribs, *v)
		}
	}
	start := xml.StartElement{Name: name, Attr: attribs}

	wr := &bytes.Buffer{}
	enc := xml.NewEncoder(wr)

	err = enc.EncodeToken(start)
	if err != nil {
		return
	}

	if el.comment != nil {
		err = enc.EncodeToken(el.comment)
		if err != nil {
			return
		}
	}

	enc.Flush()

	if el.cdata {
		wr.WriteString("<![CDATA[")
	}

	if el.value != nil {
		wr.WriteString(fmt.Sprintf("%v", el.value))
	}

	var temp string
	for _, e := range el.child {
		temp, err = e.ToXml()
		if err != nil {
			return
		}
		wr.WriteString(temp)

	}

	if el.cdata {
		wr.WriteString("]]>")
	}

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
