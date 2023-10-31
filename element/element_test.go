package element_test

import (
	"testing"

	"github.com/avelinobego/xml/element"
)

func TestChild(t *testing.T) {

	eSocial := element.Make("eSocial")
	evtInfoEmpregador := element.Make("evtInfoEmpregador")
	ideEvento := element.Make("ideEvento").
		Values("tpAmb", 2).
		Values("procEmi", 1).
		Values("verProc", "QuartaRH eSocial 1.0")

	avelino := element.Make("avelino")
	bego := element.Make("bego")
	avelino.Child(bego)
	ideEvento.Child(avelino)

	evtInfoEmpregador.Child(ideEvento)
	eSocial.Child(evtInfoEmpregador)

	f := func(c *element.Element) {
		if c.Name() == "evtInfoEmpregador" {
			c.
				Attrib("Id", "ID1123456780000002017082410324100001").
				Attrib("fake", true)
		}
	}
	eSocial.Through(f)

	f = func(c *element.Element) {
		if c.Name() == "bego" {
			c.Val("Sobrenome\ndo Avelino").Cdata(true)
		}
	}
	eSocial.Through(f)

	value, err := eSocial.ToXml()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(value)
	}

	t.Log("Fim teste")
}
