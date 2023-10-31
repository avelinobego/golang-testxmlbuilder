package element_test

import (
	"bytes"
	"element"
	"fmt"
	"testing"
)

func TestElement(t *testing.T) {

	e := element.Make("eSocial")
	e.Space("xml:bego")
	e.Comment("Este é um teste")

	buffer := bytes.Buffer{}

	a := make(map[string]interface{})
	b := make(map[string]interface{})

	a["nome"] = "nome"

	for i := 0; i < 10; i++ {
		b["id"] = i
		nome := element.Make("nome")
		nome.Value("Avelino", a, b)
		buffer.WriteString(nome.String())
	}

	line := buffer.String()
	e.Value(line)
	value, err := e.ToXml()
	if err == nil {
		fmt.Println(value)
	}

}

func TestOmmited(t *testing.T) {
	e := element.Make("eSocial")
	e.Space("xml:bego")
	e.Comment("Este é um teste")
	e.Ommit(true)

	value, err := e.ToXml()
	if err == nil {
		fmt.Println(value)
	}

}
