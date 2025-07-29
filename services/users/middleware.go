package users

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/flash-go/flash/http/client"
	"github.com/flash-go/flash/http/server"
	"github.com/flash-go/sdk/errors"
)

type MiddlewareConfig struct {
	UsersService string
	HttpClient   client.Client
}

func NewMiddleware(config *MiddlewareConfig) Middleware {
	return &middleware{
		usersService: config.UsersService,
		httpClient:   config.HttpClient,
	}
}

type Middleware interface {
	Auth(mfa bool, roles ...string) func(server.ReqHandler) server.ReqHandler
}

type middleware struct {
	usersService string
	httpClient   client.Client
}

type tokenValidateRequest struct {
	AccessToken string `json:"access_token"`
}

type tokenValidateResult struct {
	Id       string   `json:"id"`
	User     uint     `json:"user"`
	Role     string   `json:"role"`
	Mfa      bool     `json:"mfa"`
	Expires  int64    `json:"expires"`
	Issued   int64    `json:"issued"`
	Issuer   string   `json:"issuer"`
	Audience []string `json:"audience"`
}

func (m *middleware) Auth(mfa bool, roles ...string) func(server.ReqHandler) server.ReqHandler {
	return func(handler server.ReqHandler) server.ReqHandler {
		return func(ctx server.ReqCtx) {
			// Get auth token
			token, err := ctx.GetBearerToken()
			if err != nil {
				ctx.WriteErrorResponse(ErrAuthInvalidToken)
				return
			}

			// Encode body json
			body, err := json.Marshal(
				tokenValidateRequest{
					AccessToken: token,
				},
			)
			if err != nil {
				ctx.WriteErrorResponse(errors.ErrServiceUnavailable)
				return
			}

			// Send service request
			res, err := m.httpClient.ServiceRequest(
				// Context
				ctx.Context(),
				// Method
				http.MethodPost,
				// Service
				m.usersService+"-http",
				// Path
				"/users/auth/token/validate",
				// Body opt
				client.WithRequestBodyOption(body),
			)
			if err != nil {
				ctx.WriteErrorResponse(errors.ErrServiceUnavailable)
				return
			}

			// Check status code
			switch res.StatusCode() {
			case 200:
				// Parse response
				var response tokenValidateResult
				if err := json.Unmarshal(res.Body(), &response); err != nil {
					ctx.WriteErrorResponse(errors.ErrServiceUnavailable)
					return
				}

				// Set data to ctx
				ctx.SetUserValue("user", response.User)
				ctx.SetUserValue("role", response.Role)
				ctx.SetUserValue("mfa", response.Mfa)

				// Check two factor
				if mfa && response.Mfa {
					ctx.WriteErrorResponse(ErrAuth2faRequired)
					return
				}

				// Check role permissions
				if len(roles) > 0 && !slices.Contains(roles, response.Role) {
					ctx.WriteErrorResponse(ErrAuthInsufficientPermissions)
					return
				}

				handler(ctx)
			case 400:
				ctx.WriteErrorResponse(ErrAuthInvalidToken)
			default:
				ctx.WriteErrorResponse(errors.ErrServiceUnavailable)
			}
		}
	}
}

var (
	ErrAuthInsufficientPermissions = errors.New(errors.ErrForbidden, "insufficient_role_permissions")
	ErrAuth2faRequired             = errors.New(errors.ErrUnauthorized, "2fa_required")
	ErrAuthInvalidToken            = errors.New(errors.ErrUnauthorized, "invalid_token")
)
