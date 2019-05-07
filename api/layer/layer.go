package layer

import (
	"github.com/pedafy/pedafy-assignments/api"
	"github.com/pedafy/pedafy-assignments/api/apiv1"
)

const (
	Version1Beta = "v.b.1"
	Version1     = "v.1"
)

func NewApiManager(apiVersion string) api.APIHandler {
	switch apiVersion {
	case Version1:
		return apiv1.APIv1{}
	}
	return nil
}
