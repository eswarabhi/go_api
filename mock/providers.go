package mock

import (
	"fmt"

	"../model"
)

var Providers []model.ConnectorProvider = []model.ConnectorProvider{
	model.ConnectorProvider{
		DisplayName: fmt.Sprint(model.DSCRIBE),
		Name:        fmt.Sprint(model.SCRIBE),
	},
	model.ConnectorProvider{
		DisplayName: fmt.Sprint(model.DFLOGO),
		Name:        fmt.Sprint(model.FLOGO),
	},
}
