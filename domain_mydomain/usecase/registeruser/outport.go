package registeruser

import "demo-app/domain_mydomain/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveUserRepo
}
