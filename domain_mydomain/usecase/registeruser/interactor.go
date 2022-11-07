package registeruser

import (
	"context"
	"demo-app/domain_mydomain/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type registerUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &registerUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *registerUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	UserObj, err := entity.NewUser(entity.UserRequest{
		Username: req.Username,
	})

	if err != nil {
		return nil, err
	}
	err = r.outport.SaveUser(ctx, UserObj)
	if err != nil {
		return nil, err
	}

	//!

	return res, nil
}

//type loginInteractor struct {
//	outport Outport
//}

// NewUsecase is constructor for create default implementation of usecase
