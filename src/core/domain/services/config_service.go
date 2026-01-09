package core_services

import (
	coreModels "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/models"
)

type ConfigService interface {
	Read() *coreModels.Config
}
