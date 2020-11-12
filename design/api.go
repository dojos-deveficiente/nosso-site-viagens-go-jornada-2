package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = API("decolar", func() {
	Title("Decolar Service")
	Description("Decolar HTTP service")
	Server("Decolar", func() {
		Description("hosts for Dojo-1 Service.")
		Services("decolar")
		Host("development", func() {
			Description("Development hosts.")
			URI("http://localhost:3333")
		})
	})
})

var _ = Service("swagger", func() {
	Description("Swagger service")
	Files("/swagger.json", "../../gen/http/openapi.json")
})
