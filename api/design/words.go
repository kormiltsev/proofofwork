package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("words", func() {
	Error("not_found", NotFoundErrorMedia, "is a common error response for not found")
	Error("bad_request", BadRequestErrorMedia, "is a common error response for bad request")
	Error("internal", InternalErrorMedia, "is a common error response for internal error")
	Error("conflict", ConflictErrorMedia, "is a common error response for conflict error")
	Error("forbidden", ForbiddenErrorMedia, "is a common error response for forbidden error")

	Method("words", func() {
		Description(`Returns smartypents quote`)

		Payload(func() {
			Field(1, "solution", String, "Solution block", func() {
				Description("Solution block")
			})
		})

		Result(WordsMedia)

		HTTP(func() {
			POST("/words")

			Response(StatusOK)
			Response(StatusForbidden)
			Response("bad_request", StatusBadRequest)
			Response("internal", StatusInternalServerError)

			Meta("swagger:summary", "Return smartypents quote")
		})
	})

	Method("request", func() {
		Description(`First step for new quote, returns task`)

		Payload(func() {})

		Result(TaskMedia)

		HTTP(func() {
			GET("/words")

			Response(StatusOK)
			Response(StatusUnauthorized)
			Response(StatusForbidden)
			Response("bad_request", StatusBadRequest)
			Response("internal", StatusInternalServerError)

			Meta("swagger:summary", "Request for new quote, returns task")
		})
	})
})
