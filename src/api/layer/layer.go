package layer

import (
	"github.com/pedafy/pedafy-assignments/src/api"
	"github.com/pedafy/pedafy-assignments/src/api/apiv1"
	"github.com/pedafy/pedafy-assignments/src/version"
)

func NewApiManager(apiVersion string) api.APIHandler {
	switch apiVersion {
	case version.Version1:
		return &apiv1.APIv1{}
	}
	return nil
}
