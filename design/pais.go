package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var PaisDTO = Type("PaisDTO", func() {
	Description("Pais Type")
	Attribute("id", String)
	Attribute("nome", String)
	Required("id", "nome")
})

var CreatePaisDTO = Type("CreatePaisDTO", func() {
	Description("New Pais Type")
	Attribute("nome", String)
	Required("nome")
})
