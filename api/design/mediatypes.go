package design

import (
	. "goa.design/goa/v3/dsl"
)

var TaskMedia = ResultType("application/vnd.words.task+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("hash", String, "Previous hash", func() {
			Example("006b848c0e6dc6f33b76ec211a7405a5e64d2889df511b8188f1489612bfabc2")
		})
		Attribute("difficulty", Int, "Target difficulty", func() {
			Example(4)
		})
	})
})

var WordsMedia = ResultType("application/vnd.words.result+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("quote", String, "Words of wisdome", func() {
			Example("Life is too short to remove USB safely")
		})
	})
})
