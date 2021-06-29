// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/jakub-dzon/k4e-operator/models"
)

// GetDeviceConfigurationOKCode is the HTTP code returned for type GetDeviceConfigurationOK
const GetDeviceConfigurationOKCode int = 200

/*GetDeviceConfigurationOK Success.

swagger:response getDeviceConfigurationOK
*/
type GetDeviceConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.WorkloadList `json:"body,omitempty"`
}

// NewGetDeviceConfigurationOK creates GetDeviceConfigurationOK with default headers values
func NewGetDeviceConfigurationOK() *GetDeviceConfigurationOK {

	return &GetDeviceConfigurationOK{}
}

// WithPayload adds the payload to the get device configuration o k response
func (o *GetDeviceConfigurationOK) WithPayload(payload *models.WorkloadList) *GetDeviceConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get device configuration o k response
func (o *GetDeviceConfigurationOK) SetPayload(payload *models.WorkloadList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeviceConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDeviceConfigurationUnauthorizedCode is the HTTP code returned for type GetDeviceConfigurationUnauthorized
const GetDeviceConfigurationUnauthorizedCode int = 401

/*GetDeviceConfigurationUnauthorized Unauthorized

swagger:response getDeviceConfigurationUnauthorized
*/
type GetDeviceConfigurationUnauthorized struct {
}

// NewGetDeviceConfigurationUnauthorized creates GetDeviceConfigurationUnauthorized with default headers values
func NewGetDeviceConfigurationUnauthorized() *GetDeviceConfigurationUnauthorized {

	return &GetDeviceConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *GetDeviceConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetDeviceConfigurationForbiddenCode is the HTTP code returned for type GetDeviceConfigurationForbidden
const GetDeviceConfigurationForbiddenCode int = 403

/*GetDeviceConfigurationForbidden Forbidden.

swagger:response getDeviceConfigurationForbidden
*/
type GetDeviceConfigurationForbidden struct {
}

// NewGetDeviceConfigurationForbidden creates GetDeviceConfigurationForbidden with default headers values
func NewGetDeviceConfigurationForbidden() *GetDeviceConfigurationForbidden {

	return &GetDeviceConfigurationForbidden{}
}

// WriteResponse to the client
func (o *GetDeviceConfigurationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetDeviceConfigurationNotFoundCode is the HTTP code returned for type GetDeviceConfigurationNotFound
const GetDeviceConfigurationNotFoundCode int = 404

/*GetDeviceConfigurationNotFound Error.

swagger:response getDeviceConfigurationNotFound
*/
type GetDeviceConfigurationNotFound struct {
}

// NewGetDeviceConfigurationNotFound creates GetDeviceConfigurationNotFound with default headers values
func NewGetDeviceConfigurationNotFound() *GetDeviceConfigurationNotFound {

	return &GetDeviceConfigurationNotFound{}
}

// WriteResponse to the client
func (o *GetDeviceConfigurationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetDeviceConfigurationInternalServerErrorCode is the HTTP code returned for type GetDeviceConfigurationInternalServerError
const GetDeviceConfigurationInternalServerErrorCode int = 500

/*GetDeviceConfigurationInternalServerError Error.

swagger:response getDeviceConfigurationInternalServerError
*/
type GetDeviceConfigurationInternalServerError struct {
}

// NewGetDeviceConfigurationInternalServerError creates GetDeviceConfigurationInternalServerError with default headers values
func NewGetDeviceConfigurationInternalServerError() *GetDeviceConfigurationInternalServerError {

	return &GetDeviceConfigurationInternalServerError{}
}

// WriteResponse to the client
func (o *GetDeviceConfigurationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
