// Code generated by go-swagger; DO NOT EDIT.

package unidentified_containers

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

// CountByDateRangeReader is a Reader for the CountByDateRange structure.
type CountByDateRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountByDateRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountByDateRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCountByDateRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCountByDateRangeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCountByDateRangeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/unidentified-containers/count-by-date/v1] CountByDateRange", response, response.Code())
	}
}

// NewCountByDateRangeOK creates a CountByDateRangeOK with default headers values
func NewCountByDateRangeOK() *CountByDateRangeOK {
	return &CountByDateRangeOK{}
}

/*
CountByDateRangeOK describes a response with status code 200, with default header values.

OK
*/
type CountByDateRangeOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsAggregateValuesByFieldResponse
}

// IsSuccess returns true when this count by date range o k response has a 2xx status code
func (o *CountByDateRangeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this count by date range o k response has a 3xx status code
func (o *CountByDateRangeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count by date range o k response has a 4xx status code
func (o *CountByDateRangeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this count by date range o k response has a 5xx status code
func (o *CountByDateRangeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this count by date range o k response a status code equal to that given
func (o *CountByDateRangeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the count by date range o k response
func (o *CountByDateRangeOK) Code() int {
	return 200
}

func (o *CountByDateRangeOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count-by-date/v1][%d] countByDateRangeOK  %+v", 200, o.Payload)
}

func (o *CountByDateRangeOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count-by-date/v1][%d] countByDateRangeOK  %+v", 200, o.Payload)
}

func (o *CountByDateRangeOK) GetPayload() *models.ModelsAggregateValuesByFieldResponse {
	return o.Payload
}

func (o *CountByDateRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsAggregateValuesByFieldResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountByDateRangeForbidden creates a CountByDateRangeForbidden with default headers values
func NewCountByDateRangeForbidden() *CountByDateRangeForbidden {
	return &CountByDateRangeForbidden{}
}

/*
CountByDateRangeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CountByDateRangeForbidden struct {

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

// IsSuccess returns true when this count by date range forbidden response has a 2xx status code
func (o *CountByDateRangeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this count by date range forbidden response has a 3xx status code
func (o *CountByDateRangeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count by date range forbidden response has a 4xx status code
func (o *CountByDateRangeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this count by date range forbidden response has a 5xx status code
func (o *CountByDateRangeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this count by date range forbidden response a status code equal to that given
func (o *CountByDateRangeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the count by date range forbidden response
func (o *CountByDateRangeForbidden) Code() int {
	return 403
}

func (o *CountByDateRangeForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count-by-date/v1][%d] countByDateRangeForbidden  %+v", 403, o.Payload)
}

func (o *CountByDateRangeForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count-by-date/v1][%d] countByDateRangeForbidden  %+v", 403, o.Payload)
}

func (o *CountByDateRangeForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CountByDateRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCountByDateRangeTooManyRequests creates a CountByDateRangeTooManyRequests with default headers values
func NewCountByDateRangeTooManyRequests() *CountByDateRangeTooManyRequests {
	return &CountByDateRangeTooManyRequests{}
}

/*
CountByDateRangeTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CountByDateRangeTooManyRequests struct {

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

// IsSuccess returns true when this count by date range too many requests response has a 2xx status code
func (o *CountByDateRangeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this count by date range too many requests response has a 3xx status code
func (o *CountByDateRangeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count by date range too many requests response has a 4xx status code
func (o *CountByDateRangeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this count by date range too many requests response has a 5xx status code
func (o *CountByDateRangeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this count by date range too many requests response a status code equal to that given
func (o *CountByDateRangeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the count by date range too many requests response
func (o *CountByDateRangeTooManyRequests) Code() int {
	return 429
}

func (o *CountByDateRangeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count-by-date/v1][%d] countByDateRangeTooManyRequests  %+v", 429, o.Payload)
}

func (o *CountByDateRangeTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count-by-date/v1][%d] countByDateRangeTooManyRequests  %+v", 429, o.Payload)
}

func (o *CountByDateRangeTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CountByDateRangeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCountByDateRangeInternalServerError creates a CountByDateRangeInternalServerError with default headers values
func NewCountByDateRangeInternalServerError() *CountByDateRangeInternalServerError {
	return &CountByDateRangeInternalServerError{}
}

/*
CountByDateRangeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CountByDateRangeInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CoreEntitiesResponse
}

// IsSuccess returns true when this count by date range internal server error response has a 2xx status code
func (o *CountByDateRangeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this count by date range internal server error response has a 3xx status code
func (o *CountByDateRangeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this count by date range internal server error response has a 4xx status code
func (o *CountByDateRangeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this count by date range internal server error response has a 5xx status code
func (o *CountByDateRangeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this count by date range internal server error response a status code equal to that given
func (o *CountByDateRangeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the count by date range internal server error response
func (o *CountByDateRangeInternalServerError) Code() int {
	return 500
}

func (o *CountByDateRangeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count-by-date/v1][%d] countByDateRangeInternalServerError  %+v", 500, o.Payload)
}

func (o *CountByDateRangeInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/unidentified-containers/count-by-date/v1][%d] countByDateRangeInternalServerError  %+v", 500, o.Payload)
}

func (o *CountByDateRangeInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *CountByDateRangeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CoreEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}