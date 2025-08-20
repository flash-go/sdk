package middleware

import (
	"github.com/flash-go/flash/http/server"
	"github.com/flash-go/sdk/errors"
)

func ParseJsonBody[T interface{ Validate() error }]() func(server.ReqHandler) server.ReqHandler {
	return func(handler server.ReqHandler) server.ReqHandler {
		return func(ctx server.ReqCtx) {
			// Parse body
			var body T
			if err := ctx.ReadJson(&body); err != nil {
				ctx.WriteErrorResponse(errors.ErrBadRequest)
				return
			}

			// Validate body
			if err := body.Validate(); err != nil {
				ctx.WriteErrorResponse(err)
				return
			}

			// Set body
			ctx.SetUserValue("body", body)

			handler(ctx)
		}
	}
}
