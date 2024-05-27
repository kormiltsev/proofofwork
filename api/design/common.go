package design

import (
	. "goa.design/goa/v3/dsl"
)

var WordsErrorType = Type("vndError", func() {
	Attribute("msgCode", String, "Key to select a localized message", func() {
		Example("JWT_TOKEN_EXPIRED")
	})
	Attribute("msg", String, "Detailed error message", func() {
		Example("JWT token has expired.")
	})
	Attribute("attributes", Any, "Message attributes", func() {
		Example("{\"entityId\":\"xxxx-xxxx-xxxx-xxxx\"}")
	})
	Required("msg", "msgCode")
})

var InternalErrorMedia = ResultType("application/vnd.internal.error+json", func() {
	ContentType("application/json")
	Reference(WordsErrorType)
	Attributes(func() {
		Attribute("msgCode")
		Attribute("msg")
		Attribute("attributes")
	})
	Required("msg", "msgCode")
})

var NotFoundErrorMedia = ResultType("application/vnd.not.found.error+json", func() {
	ContentType("application/json")
	Reference(WordsErrorType)
	Attributes(func() {
		Attribute("msgCode")
		Attribute("msg")
		Attribute("attributes")
	})
	Required("msg", "msgCode")
})

var BadRequestErrorMedia = ResultType("application/vnd.bad.request.error+json", func() {
	ContentType("application/json")
	Reference(WordsErrorType)
	Attributes(func() {
		Attribute("msgCode")
		Attribute("msg")
		Attribute("attributes")
	})
	Required("msg", "msgCode")
})

var ConflictErrorMedia = ResultType("application/vnd.conflict.error+json", func() {
	ContentType("application/json")
	Reference(WordsErrorType)
	Attributes(func() {
		Attribute("msgCode")
		Attribute("msg")
		Attribute("attributes")
	})
	Required("msg", "msgCode")
})

var ForbiddenErrorMedia = ResultType("application/vnd.forbidden.error+json", func() {
	ContentType("application/json")
	Reference(WordsErrorType)
	Attributes(func() {
		Attribute("msgCode")
		Attribute("msg")
		Attribute("attributes")
	})
	Required("msg", "msgCode")
})
