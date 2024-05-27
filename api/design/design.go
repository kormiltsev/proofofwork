package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("word-of-wisdom", func() {
	Version("1.0")
	Title("Word of Wosdom")
	Description("RESTful API for Word of Wisdom service.")
	Server("word-of-wisdom", func() {
		Services(
			"words",
		)

		Host("localhost", func() {
			URI("https://localhost:8080/words")
		})
	})
})
