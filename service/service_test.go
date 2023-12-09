package service_test

import (
	"sample/elastic"
	"sample/mocks"
	"sample/service"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestJob(t *testing.T) {
	// Create a new controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock instance of IElasticService
	mockElasticService := mocks.NewMockIElasticService(ctrl)

	// Set expectations on the mock
	mockElasticService.EXPECT().ESscroll("hello").Return("esscroll").Times(2)
	mockElasticService.EXPECT().Method2("hello").Return("method2").Times(2)

	// Override the factory function with the mock in the test
	service.SetElasticServiceFactory(func(x, y, z, a string) elastic.IElasticService {
		return mockElasticService
	})
	defer service.ResetElasticServiceFactory() // Reset the factory after the test

	// Call the Job function
	service.Job()

	// No need to assert anything here, the expectations on the mockElasticService will fail the test if not met
}
