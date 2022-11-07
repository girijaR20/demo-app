package payload

import (
	"demo-app/shared/driver"
)

type Payload struct {
	Data      any                    `json:"data"`
	Publisher driver.ApplicationData `json:"publisher"`
	TraceID   string                 `json:"traceId"`
}
