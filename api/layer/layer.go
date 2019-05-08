package layer

import (
	"github.com/pedafy/pedafy-assignments/api"
	"github.com/pedafy/pedafy-assignments/api/apiv1"
	"github.com/pedafy/pedafy-assignments/version"
)

func NewApiManager(apiVersion string) api.APIHandler {
	switch apiVersion {
	case version.Version1:
		return &apiv1.APIv1{}
	}
	return nil
}
