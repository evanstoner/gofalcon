// Code generated by go-swagger; DO NOT EDIT.

package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// PromoteSystemDefinitionReader is a Reader for the PromoteSystemDefinition structure.
type PromoteSystemDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PromoteSystemDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPromoteSystemDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPromoteSystemDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPromoteSystemDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPromoteSystemDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPromoteSystemDefinitionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPromoteSystemDefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /workflows/system-definitions/promote/v1] promote.system-definition", response, response.Code())
	}
}

// NewPromoteSystemDefinitionOK creates a PromoteSystemDefinitionOK with default headers values
func NewPromoteSystemDefinitionOK() *PromoteSystemDefinitionOK {
	return &PromoteSystemDefinitionOK{}
}

/*
PromoteSystemDefinitionOK describes a response with status code 200, with default header values.

OK
*/
type PromoteSystemDefinitionOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSystemDefinitionCreateResponse
}

// IsSuccess returns true when this promote system definition o k response has a 2xx status code
func (o *PromoteSystemDefinitionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this promote system definition o k response has a 3xx status code
func (o *PromoteSystemDefinitionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote system definition o k response has a 4xx status code
func (o *PromoteSystemDefinitionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this promote system definition o k response has a 5xx status code
func (o *PromoteSystemDefinitionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this promote system definition o k response a status code equal to that given
func (o *PromoteSystemDefinitionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the promote system definition o k response
func (o *PromoteSystemDefinitionOK) Code() int {
	return 200
}

func (o *PromoteSystemDefinitionOK) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionOK  %+v", 200, o.Payload)
}

func (o *PromoteSystemDefinitionOK) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionOK  %+v", 200, o.Payload)
}

func (o *PromoteSystemDefinitionOK) GetPayload() *models.ClientSystemDefinitionCreateResponse {
	return o.Payload
}

func (o *PromoteSystemDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ClientSystemDefinitionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteSystemDefinitionBadRequest creates a PromoteSystemDefinitionBadRequest with default headers values
func NewPromoteSystemDefinitionBadRequest() *PromoteSystemDefinitionBadRequest {
	return &PromoteSystemDefinitionBadRequest{}
}

/*
PromoteSystemDefinitionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PromoteSystemDefinitionBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSystemDefinitionCreateResponse
}

// IsSuccess returns true when this promote system definition bad request response has a 2xx status code
func (o *PromoteSystemDefinitionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this promote system definition bad request response has a 3xx status code
func (o *PromoteSystemDefinitionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote system definition bad request response has a 4xx status code
func (o *PromoteSystemDefinitionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this promote system definition bad request response has a 5xx status code
func (o *PromoteSystemDefinitionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this promote system definition bad request response a status code equal to that given
func (o *PromoteSystemDefinitionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the promote system definition bad request response
func (o *PromoteSystemDefinitionBadRequest) Code() int {
	return 400
}

func (o *PromoteSystemDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *PromoteSystemDefinitionBadRequest) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *PromoteSystemDefinitionBadRequest) GetPayload() *models.ClientSystemDefinitionCreateResponse {
	return o.Payload
}

func (o *PromoteSystemDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ClientSystemDefinitionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteSystemDefinitionForbidden creates a PromoteSystemDefinitionForbidden with default headers values
func NewPromoteSystemDefinitionForbidden() *PromoteSystemDefinitionForbidden {
	return &PromoteSystemDefinitionForbidden{}
}

/*
PromoteSystemDefinitionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PromoteSystemDefinitionForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this promote system definition forbidden response has a 2xx status code
func (o *PromoteSystemDefinitionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this promote system definition forbidden response has a 3xx status code
func (o *PromoteSystemDefinitionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote system definition forbidden response has a 4xx status code
func (o *PromoteSystemDefinitionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this promote system definition forbidden response has a 5xx status code
func (o *PromoteSystemDefinitionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this promote system definition forbidden response a status code equal to that given
func (o *PromoteSystemDefinitionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the promote system definition forbidden response
func (o *PromoteSystemDefinitionForbidden) Code() int {
	return 403
}

func (o *PromoteSystemDefinitionForbidden) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *PromoteSystemDefinitionForbidden) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *PromoteSystemDefinitionForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PromoteSystemDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteSystemDefinitionNotFound creates a PromoteSystemDefinitionNotFound with default headers values
func NewPromoteSystemDefinitionNotFound() *PromoteSystemDefinitionNotFound {
	return &PromoteSystemDefinitionNotFound{}
}

/*
PromoteSystemDefinitionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PromoteSystemDefinitionNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSystemDefinitionCreateResponse
}

// IsSuccess returns true when this promote system definition not found response has a 2xx status code
func (o *PromoteSystemDefinitionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this promote system definition not found response has a 3xx status code
func (o *PromoteSystemDefinitionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote system definition not found response has a 4xx status code
func (o *PromoteSystemDefinitionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this promote system definition not found response has a 5xx status code
func (o *PromoteSystemDefinitionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this promote system definition not found response a status code equal to that given
func (o *PromoteSystemDefinitionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the promote system definition not found response
func (o *PromoteSystemDefinitionNotFound) Code() int {
	return 404
}

func (o *PromoteSystemDefinitionNotFound) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *PromoteSystemDefinitionNotFound) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *PromoteSystemDefinitionNotFound) GetPayload() *models.ClientSystemDefinitionCreateResponse {
	return o.Payload
}

func (o *PromoteSystemDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ClientSystemDefinitionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteSystemDefinitionTooManyRequests creates a PromoteSystemDefinitionTooManyRequests with default headers values
func NewPromoteSystemDefinitionTooManyRequests() *PromoteSystemDefinitionTooManyRequests {
	return &PromoteSystemDefinitionTooManyRequests{}
}

/*
PromoteSystemDefinitionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PromoteSystemDefinitionTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this promote system definition too many requests response has a 2xx status code
func (o *PromoteSystemDefinitionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this promote system definition too many requests response has a 3xx status code
func (o *PromoteSystemDefinitionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote system definition too many requests response has a 4xx status code
func (o *PromoteSystemDefinitionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this promote system definition too many requests response has a 5xx status code
func (o *PromoteSystemDefinitionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this promote system definition too many requests response a status code equal to that given
func (o *PromoteSystemDefinitionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the promote system definition too many requests response
func (o *PromoteSystemDefinitionTooManyRequests) Code() int {
	return 429
}

func (o *PromoteSystemDefinitionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PromoteSystemDefinitionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PromoteSystemDefinitionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PromoteSystemDefinitionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteSystemDefinitionInternalServerError creates a PromoteSystemDefinitionInternalServerError with default headers values
func NewPromoteSystemDefinitionInternalServerError() *PromoteSystemDefinitionInternalServerError {
	return &PromoteSystemDefinitionInternalServerError{}
}

/*
PromoteSystemDefinitionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PromoteSystemDefinitionInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSystemDefinitionCreateResponse
}

// IsSuccess returns true when this promote system definition internal server error response has a 2xx status code
func (o *PromoteSystemDefinitionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this promote system definition internal server error response has a 3xx status code
func (o *PromoteSystemDefinitionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote system definition internal server error response has a 4xx status code
func (o *PromoteSystemDefinitionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this promote system definition internal server error response has a 5xx status code
func (o *PromoteSystemDefinitionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this promote system definition internal server error response a status code equal to that given
func (o *PromoteSystemDefinitionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the promote system definition internal server error response
func (o *PromoteSystemDefinitionInternalServerError) Code() int {
	return 500
}

func (o *PromoteSystemDefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *PromoteSystemDefinitionInternalServerError) String() string {
	return fmt.Sprintf("[POST /workflows/system-definitions/promote/v1][%d] promoteSystemDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *PromoteSystemDefinitionInternalServerError) GetPayload() *models.ClientSystemDefinitionCreateResponse {
	return o.Payload
}

func (o *PromoteSystemDefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ClientSystemDefinitionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}