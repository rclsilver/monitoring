package mqtt

import (
	"github.com/gin-gonic/gin"
)

type healthOutStatus string

const (
	healthOutStatusOK          healthOutStatus = "OK"
	healthOutStatusUnavailable healthOutStatus = "UNAVAILABLE"
	healthOutStatusUnknown     healthOutStatus = "UNKNOWN"
)

type healthOut struct {
	Status healthOutStatus `json:"status"`
	Error  error           `json:"error"`
}

func (mqtt *MQTTComponent) handlerPing(c *gin.Context) (*healthOut, error) {
	mqtt.healthMut.Lock()
	defer mqtt.healthMut.Unlock()

	return &healthOut{
		Status: mqtt.health.Status,
		Error:  mqtt.health.Error,
	}, nil
}
