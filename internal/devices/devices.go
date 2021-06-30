package devices

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jakub-dzon/k4e-operator/models"
	operations "github.com/jakub-dzon/k4e-operator/restapi/operations/devices"
)

type DeviceHandler struct {
}

func NewDeviceHandler() *DeviceHandler {
	return &DeviceHandler{}
}

func (h *DeviceHandler) GetDeviceConfiguration(ctx context.Context, params operations.GetDeviceConfigurationParams) middleware.Responder {
	// TODO: retrieve CRs from the cluster and respond
	return operations.NewGetDeviceConfigurationOK().WithPayload(&models.DeviceConfiguration{DeviceID: params.DeviceID})
}
