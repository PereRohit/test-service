package logic

import (
	"net/http"

	"github.com/PereRohit/util/log"
    respModel "github.com/PereRohit/util/model"

	"github.com/PereRohit/test-service/internal/model"
	"github.com/PereRohit/test-service/internal/repo/datasource"
)

//go:generate mockgen --build_flags=--mod=mod --destination=./../../pkg/mock/mock_logic.go --package=mock github.com/PereRohit/test-service/internal/logic TestServiceLogicIer

type TestServiceLogicIer interface {
	Ping(*model.PingRequest) *respModel.Response
    HealthCheck() bool
}

type testServiceLogic struct{
	dummyDsSvc datasource.DataSource
}

func NewTestServiceLogic(ds datasource.DataSource) TestServiceLogicIer {
	return &testServiceLogic{
		dummyDsSvc: ds,
    }
}

func (l testServiceLogic) Ping(req *model.PingRequest) *respModel.Response {
	// add business logic here
	res, err := l.dummyDsSvc.Ping(&model.PingDs{
    	Data: req.Data,
    })
    if err != nil {
        log.Error("datasource error", err)
    	return &respModel.Response{
    		Status:  http.StatusInternalServerError,
    		Message: "",
    		Data:    nil,
    	}
    }
    return &respModel.Response{
    	Status:  http.StatusOK,
    	Message: "Pong",
    	Data:    res,
    }
}

func (l testServiceLogic) HealthCheck() bool {
	// check all internal services are working fine
	return l.dummyDsSvc.HealthCheck()
}