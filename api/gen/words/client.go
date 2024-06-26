// Code generated by goa v3.16.2, DO NOT EDIT.
//
// words client
//
// Command:
// $ goa gen github.com/kormiltsev/proofofwork/api/design -o ./api/

package words

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "words" service client.
type Client struct {
	WordsEndpoint   goa.Endpoint
	RequestEndpoint goa.Endpoint
}

// NewClient initializes a "words" service client given the endpoints.
func NewClient(words, request goa.Endpoint) *Client {
	return &Client{
		WordsEndpoint:   words,
		RequestEndpoint: request,
	}
}

// Words calls the "words" endpoint of the "words" service.
// Words may return the following errors:
//   - "not_found" (type *NotFoundError): is a common error response for not found
//   - "bad_request" (type *BadRequestError): is a common error response for bad request
//   - "internal" (type *InternalError): is a common error response for internal error
//   - "conflict" (type *ConflictError): is a common error response for conflict error
//   - "forbidden" (type *ForbiddenError): is a common error response for forbidden error
//   - error: internal error
func (c *Client) Words(ctx context.Context, p *WordsPayload) (res *WordsResult, err error) {
	var ires any
	ires, err = c.WordsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*WordsResult), nil
}

// Request calls the "request" endpoint of the "words" service.
// Request may return the following errors:
//   - "not_found" (type *NotFoundError): is a common error response for not found
//   - "bad_request" (type *BadRequestError): is a common error response for bad request
//   - "internal" (type *InternalError): is a common error response for internal error
//   - "conflict" (type *ConflictError): is a common error response for conflict error
//   - "forbidden" (type *ForbiddenError): is a common error response for forbidden error
//   - error: internal error
func (c *Client) Request(ctx context.Context) (res *WordsTask, err error) {
	var ires any
	ires, err = c.RequestEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*WordsTask), nil
}
