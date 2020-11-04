package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = Service("pais", func() {
	Description("O serviço pais executa operações em pais")

	Error("invalid_fields")

	Method("create_pais", func() {
		Payload(CreatePaisDTO)
		Result(PaisDTO)
		HTTP(func() {
			POST("/pais")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})
})

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
