package handler

import (
	"fmt"
	"net/http"

	"github.com/PereRohit/util/request"
	"github.com/PereRohit/util/response"

	"github.com/PereRohit/test-service/internal/logic"
	"github.com/PereRohit/test-service/internal/model"
	"github.com/PereRohit/test-service/internal/repo/datasource"
)

const TestServiceName = "testService"

//go:generate mockgen --build_flags=--mod=mod --destination=./../../pkg/mock/mock_handler.go --package=mock github.com/PereRohit/test-service/internal/handler TestServiceHandler

type TestServiceHandler interface {
	HealthChecker
	Ping(w http.ResponseWriter, r *http.Request)
}

type testService struct {
	logic logic.TestServiceLogicIer
}

func NewTestService(ds datasource.DataSource) TestServiceHandler {
	svc := &testService{
		logic: logic.NewTestServiceLogic(ds),
	}
    AddHealthChecker(svc)
	return svc
}

func (svc testService) HealthCheck() (svcName string, msg string, stat bool) {
	set := false
	defer func() {
		svcName = TestServiceName
		if !set {
			msg = ""
			stat = true
		}
	}()
	stat = svc.logic.HealthCheck()
    set = true
	return
}

func (svc testService) Ping(w http.ResponseWriter, r *http.Request) {
	req := &model.PingRequest{}

	suggestedCode, err := request.FromJson(r, req)
	if err != nil {
		response.ToJson(w, suggestedCode, fmt.Sprintf("FAILED: %s", err.Error()), nil)
		return
	}
	// call logic
	resp := svc.logic.Ping(req)
	response.ToJson(w, resp.Status, resp.Message, resp.Data)
	return
}