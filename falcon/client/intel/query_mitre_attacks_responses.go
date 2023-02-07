// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// QueryMitreAttacksReader is a Reader for the QueryMitreAttacks structure.
type QueryMitreAttacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryMitreAttacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryMitreAttacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryMitreAttacksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryMitreAttacksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryMitreAttacksOK creates a QueryMitreAttacksOK with default headers values
func NewQueryMitreAttacksOK() *QueryMitreAttacksOK {
	return &QueryMitreAttacksOK{}
}

/*
QueryMitreAttacksOK describes a response with status code 200, with default header values.

OK
*/
type QueryMitreAttacksOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this query mitre attacks o k response has a 2xx status code
func (o *QueryMitreAttacksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query mitre attacks o k response has a 3xx status code
func (o *QueryMitreAttacksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query mitre attacks o k response has a 4xx status code
func (o *QueryMitreAttacksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query mitre attacks o k response has a 5xx status code
func (o *QueryMitreAttacksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query mitre attacks o k response a status code equal to that given
func (o *QueryMitreAttacksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query mitre attacks o k response
func (o *QueryMitreAttacksOK) Code() int {
	return 200
}

func (o *QueryMitreAttacksOK) Error() string {
	return fmt.Sprintf("[GET /intel/queries/mitre/v1][%d] queryMitreAttacksOK ", 200)
}

func (o *QueryMitreAttacksOK) String() string {
	return fmt.Sprintf("[GET /intel/queries/mitre/v1][%d] queryMitreAttacksOK ", 200)
}

func (o *QueryMitreAttacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewQueryMitreAttacksForbidden creates a QueryMitreAttacksForbidden with default headers values
func NewQueryMitreAttacksForbidden() *QueryMitreAttacksForbidden {
	return &QueryMitreAttacksForbidden{}
}

/*
QueryMitreAttacksForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryMitreAttacksForbidden struct {

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

// IsSuccess returns true when this query mitre attacks forbidden response has a 2xx status code
func (o *QueryMitreAttacksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query mitre attacks forbidden response has a 3xx status code
func (o *QueryMitreAttacksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query mitre attacks forbidden response has a 4xx status code
func (o *QueryMitreAttacksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query mitre attacks forbidden response has a 5xx status code
func (o *QueryMitreAttacksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query mitre attacks forbidden response a status code equal to that given
func (o *QueryMitreAttacksForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query mitre attacks forbidden response
func (o *QueryMitreAttacksForbidden) Code() int {
	return 403
}

func (o *QueryMitreAttacksForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/queries/mitre/v1][%d] queryMitreAttacksForbidden  %+v", 403, o.Payload)
}

func (o *QueryMitreAttacksForbidden) String() string {
	return fmt.Sprintf("[GET /intel/queries/mitre/v1][%d] queryMitreAttacksForbidden  %+v", 403, o.Payload)
}

func (o *QueryMitreAttacksForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryMitreAttacksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryMitreAttacksTooManyRequests creates a QueryMitreAttacksTooManyRequests with default headers values
func NewQueryMitreAttacksTooManyRequests() *QueryMitreAttacksTooManyRequests {
	return &QueryMitreAttacksTooManyRequests{}
}

/*
QueryMitreAttacksTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryMitreAttacksTooManyRequests struct {

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

// IsSuccess returns true when this query mitre attacks too many requests response has a 2xx status code
func (o *QueryMitreAttacksTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query mitre attacks too many requests response has a 3xx status code
func (o *QueryMitreAttacksTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query mitre attacks too many requests response has a 4xx status code
func (o *QueryMitreAttacksTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query mitre attacks too many requests response has a 5xx status code
func (o *QueryMitreAttacksTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query mitre attacks too many requests response a status code equal to that given
func (o *QueryMitreAttacksTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query mitre attacks too many requests response
func (o *QueryMitreAttacksTooManyRequests) Code() int {
	return 429
}

func (o *QueryMitreAttacksTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/queries/mitre/v1][%d] queryMitreAttacksTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryMitreAttacksTooManyRequests) String() string {
	return fmt.Sprintf("[GET /intel/queries/mitre/v1][%d] queryMitreAttacksTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryMitreAttacksTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryMitreAttacksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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