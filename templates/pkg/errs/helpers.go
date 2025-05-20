package errs

import (
	"encoding/json"
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IsCode(err error, code string) bool {
	for err != nil {
		var ae *AppError
		if errors.As(err, &ae) && ae.Code == code {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}

// HTTPErrorResponse — the structure of the JSON response for the HTTP API
type HTTPErrorResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
}

func (e *AppError) HTTPStatus() int {
	switch e.Code {
	// 400 Bad Request
	case CodeBadRequest, CodeValidationError, CodeMissingParameter, CodeInvalidParameter,
		CodeParsingError, CodeSerializationError, CodePasswordTooWeak:
		return http.StatusBadRequest

	// 401 Unauthorized
	case CodeUnauthorized, CodeInvalidCredentials, CodeTokenExpired, CodeTokenInvalid,
		CodeOAuthError, CodeJWTError:
		return http.StatusUnauthorized

	// 403 Forbidden
	case CodeForbidden, CodeAccountDisabled, CodeAccountLocked, CodePermissionDenied:
		return http.StatusForbidden

	// 404 Not Found
	case CodeNotFound, CodeRecordNotFound, CodeFileNotFound:
		return http.StatusNotFound

	// 405 Method Not Allowed
	case CodeMethodNotAllowed:
		return http.StatusMethodNotAllowed

	// 409 Conflict
	case CodeConflict, CodeRecordAlreadyExists:
		return http.StatusConflict

	// 413 Payload Too Large
	case CodePayloadTooLarge:
		return http.StatusRequestEntityTooLarge

	// 414 URI Too Long
	case CodeRequestURITooLong:
		return http.StatusRequestURITooLong

	// 415 Unsupported Media Type
	case CodeUnsupportedMediaType:
		return http.StatusUnsupportedMediaType

	// 429 Too Many Requests
	case CodeTooManyRequests:
		return http.StatusTooManyRequests

	// 503 Service Unavailable
	case CodeServiceUnavailable, CodeDependencyFailure, CodeExternalServiceError,
		CodeNetworkError, CodeConnectionError, CodeDNSError, CodeTLSHandshakeError:
		return http.StatusServiceUnavailable

	// 408 Request Timeout
	case CodeTimeout:
		return http.StatusRequestTimeout

	// 402 Payment Required
	case CodePaymentDeclined, CodeInsufficientFunds:
		return http.StatusPaymentRequired

	// 409 Conflict
	case CodeBusinessRuleViolation:
		return http.StatusConflict

	// 423 Locked
	case CodeResourceLocked:
		return http.StatusLocked

	// Default: 500 Internal Server Error
	default:
		return http.StatusInternalServerError
	}
}

// ToHTTPResponseBody returns the JSON-encoded bytes for an HTTPErrorResponse
func (e *AppError) ToHTTPResponseBody() []byte {
	resp := HTTPErrorResponse{
		Code:    e.Code,
		Message: e.Message,
	}
	if len(e.Fields) > 0 {
		resp.Fields = make(map[string]interface{}, len(e.Fields))
		for _, f := range e.Fields {
			resp.Fields[f.Key] = f.Value
		}
	}
	data, _ := json.Marshal(resp)
	return data
}

// ToGRPCError конвертирует ошибку в gRPC ошибку
func ToGRPCError(err error) error {
	if err == nil {
		return nil
	}

	appErr, ok := err.(*AppError)
	if !ok {
		return status.Error(codes.Internal, err.Error())
	}

	code := codes.Internal
	switch appErr.Code {
	case CodeBadRequest:
		code = codes.InvalidArgument
	case CodeNotFound:
		code = codes.NotFound
	case CodeUnauthorized:
		code = codes.Unauthenticated
	case CodeForbidden:
		code = codes.PermissionDenied
	}

	return status.Error(code, appErr.Message)
}
