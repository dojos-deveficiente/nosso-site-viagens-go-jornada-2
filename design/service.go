package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = Service("decolar", func() {
	Description("The decolar service")

	Error("invalid_fields")
	Error("not_found")

	Method("create_pais", func() {
		Payload(CreatePaisDTO)
		Result(PaisDTO)
		HTTP(func() {
			POST("/pais")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("create_compania", func() {
		Payload(CreateCompaniaDTO)
		Result(CompaniaDTO)
		HTTP(func() {
			POST("/compania")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})

	Method("create_aeroporto", func() {
		Payload(CreateAeroportoDTO)
		Result(AeroportoDTO)
		HTTP(func() {
			POST("/aeroporto")
			Response(StatusCreated)
			Response("invalid_fields", StatusBadRequest)
		})
	})
})
