package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var CompaniaDTO = Type("CompaniaDTO", func() {
	Description("Compania Type")
	Attribute("id", String)
	Attribute("nome", String)
	Attribute("pais_id", String)
	Attribute("created_at", String)
	Required("id", "nome", "pais_id", "created_at")
})

var CreateCompaniaDTO = Type("CreateCompaniaDTO", func() {
	Description("New Compania Type")
	Attribute("nome", String, func() {
		MinLength(1)
	})
	Attribute("pais_id", String, func() {
		Format(FormatUUID)
	})
	Required("nome", "pais_id")
})
