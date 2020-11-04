package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

var _ = API("dojo-1", func() {
	Title("Dojo-1 Service")
	Description("Dojo-1 HTTP service")
	Server("dojo-1", func() {
		Description("hosts for Dojo-1 Service.")
		Services("pais")
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
