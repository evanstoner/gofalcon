// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// GetCIDGroupMembersByV2Reader is a Reader for the GetCIDGroupMembersByV2 structure.
type GetCIDGroupMembersByV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCIDGroupMembersByV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCIDGroupMembersByV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCIDGroupMembersByV2MultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCIDGroupMembersByV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCIDGroupMembersByV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCIDGroupMembersByV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCIDGroupMembersByV2OK creates a GetCIDGroupMembersByV2OK with default headers values
func NewGetCIDGroupMembersByV2OK() *GetCIDGroupMembersByV2OK {
	return &GetCIDGroupMembersByV2OK{}
}

/*
GetCIDGroupMembersByV2OK describes a response with status code 200, with default header values.

OK
*/
type GetCIDGroupMembersByV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

// IsSuccess returns true when this get c Id group members by v2 o k response has a 2xx status code
func (o *GetCIDGroupMembersByV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c Id group members by v2 o k response has a 3xx status code
func (o *GetCIDGroupMembersByV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by v2 o k response has a 4xx status code
func (o *GetCIDGroupMembersByV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c Id group members by v2 o k response has a 5xx status code
func (o *GetCIDGroupMembersByV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by v2 o k response a status code equal to that given
func (o *GetCIDGroupMembersByV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get c Id group members by v2 o k response
func (o *GetCIDGroupMembersByV2OK) Code() int {
	return 200
}

func (o *GetCIDGroupMembersByV2OK) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2OK  %+v", 200, o.Payload)
}

func (o *GetCIDGroupMembersByV2OK) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2OK  %+v", 200, o.Payload)
}

func (o *GetCIDGroupMembersByV2OK) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupMembersByV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupMembersByV2MultiStatus creates a GetCIDGroupMembersByV2MultiStatus with default headers values
func NewGetCIDGroupMembersByV2MultiStatus() *GetCIDGroupMembersByV2MultiStatus {
	return &GetCIDGroupMembersByV2MultiStatus{}
}

/*
GetCIDGroupMembersByV2MultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetCIDGroupMembersByV2MultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

// IsSuccess returns true when this get c Id group members by v2 multi status response has a 2xx status code
func (o *GetCIDGroupMembersByV2MultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get c Id group members by v2 multi status response has a 3xx status code
func (o *GetCIDGroupMembersByV2MultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by v2 multi status response has a 4xx status code
func (o *GetCIDGroupMembersByV2MultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get c Id group members by v2 multi status response has a 5xx status code
func (o *GetCIDGroupMembersByV2MultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by v2 multi status response a status code equal to that given
func (o *GetCIDGroupMembersByV2MultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get c Id group members by v2 multi status response
func (o *GetCIDGroupMembersByV2MultiStatus) Code() int {
	return 207
}

func (o *GetCIDGroupMembersByV2MultiStatus) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2MultiStatus  %+v", 207, o.Payload)
}

func (o *GetCIDGroupMembersByV2MultiStatus) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2MultiStatus  %+v", 207, o.Payload)
}

func (o *GetCIDGroupMembersByV2MultiStatus) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupMembersByV2MultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupMembersByV2BadRequest creates a GetCIDGroupMembersByV2BadRequest with default headers values
func NewGetCIDGroupMembersByV2BadRequest() *GetCIDGroupMembersByV2BadRequest {
	return &GetCIDGroupMembersByV2BadRequest{}
}

/*
GetCIDGroupMembersByV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCIDGroupMembersByV2BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get c Id group members by v2 bad request response has a 2xx status code
func (o *GetCIDGroupMembersByV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c Id group members by v2 bad request response has a 3xx status code
func (o *GetCIDGroupMembersByV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by v2 bad request response has a 4xx status code
func (o *GetCIDGroupMembersByV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c Id group members by v2 bad request response has a 5xx status code
func (o *GetCIDGroupMembersByV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by v2 bad request response a status code equal to that given
func (o *GetCIDGroupMembersByV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get c Id group members by v2 bad request response
func (o *GetCIDGroupMembersByV2BadRequest) Code() int {
	return 400
}

func (o *GetCIDGroupMembersByV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetCIDGroupMembersByV2BadRequest) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetCIDGroupMembersByV2BadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetCIDGroupMembersByV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupMembersByV2Forbidden creates a GetCIDGroupMembersByV2Forbidden with default headers values
func NewGetCIDGroupMembersByV2Forbidden() *GetCIDGroupMembersByV2Forbidden {
	return &GetCIDGroupMembersByV2Forbidden{}
}

/*
GetCIDGroupMembersByV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCIDGroupMembersByV2Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get c Id group members by v2 forbidden response has a 2xx status code
func (o *GetCIDGroupMembersByV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c Id group members by v2 forbidden response has a 3xx status code
func (o *GetCIDGroupMembersByV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by v2 forbidden response has a 4xx status code
func (o *GetCIDGroupMembersByV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c Id group members by v2 forbidden response has a 5xx status code
func (o *GetCIDGroupMembersByV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by v2 forbidden response a status code equal to that given
func (o *GetCIDGroupMembersByV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get c Id group members by v2 forbidden response
func (o *GetCIDGroupMembersByV2Forbidden) Code() int {
	return 403
}

func (o *GetCIDGroupMembersByV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetCIDGroupMembersByV2Forbidden) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetCIDGroupMembersByV2Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetCIDGroupMembersByV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupMembersByV2TooManyRequests creates a GetCIDGroupMembersByV2TooManyRequests with default headers values
func NewGetCIDGroupMembersByV2TooManyRequests() *GetCIDGroupMembersByV2TooManyRequests {
	return &GetCIDGroupMembersByV2TooManyRequests{}
}

/*
GetCIDGroupMembersByV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCIDGroupMembersByV2TooManyRequests struct {

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

// IsSuccess returns true when this get c Id group members by v2 too many requests response has a 2xx status code
func (o *GetCIDGroupMembersByV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get c Id group members by v2 too many requests response has a 3xx status code
func (o *GetCIDGroupMembersByV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get c Id group members by v2 too many requests response has a 4xx status code
func (o *GetCIDGroupMembersByV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get c Id group members by v2 too many requests response has a 5xx status code
func (o *GetCIDGroupMembersByV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get c Id group members by v2 too many requests response a status code equal to that given
func (o *GetCIDGroupMembersByV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get c Id group members by v2 too many requests response
func (o *GetCIDGroupMembersByV2TooManyRequests) Code() int {
	return 429
}

func (o *GetCIDGroupMembersByV2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCIDGroupMembersByV2TooManyRequests) String() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v2][%d] getCIdGroupMembersByV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCIDGroupMembersByV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCIDGroupMembersByV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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