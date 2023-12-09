# How to write Unit Test for Tightly coupled code in Go.

we can use a variable to handle this in golang as mentioned in the Job.go

```go
var elasticServiceFactory = elastic.NewInstance

func SetElasticServiceFactory(factory func(x, y, z, a string) elastic.IElasticService) {
	elasticServiceFactory = factory
}

func ResetElasticServiceFactory() {
	elasticServiceFactory = elastic.NewInstance
}
```

---

## To Generate Mocks

Firstly get required packages.

> go get github.com/golang/mock/mockgen@v1.6.0

> go get github.com/golang/mock/gomock

Add below line on the file for which mocks has to be created.

> //go:generate mockgen -source=elastic/elasticservice.go -destination=mocks/mock_elasticservice.go -package=mocks

Run below command for generating mocks.

> go generate ./...

Run below command if any import issues occur.

> go mod tidy

---

## Unit Test

Write your unit test - please refer the unit test code.

Run below command for Coverage.

> go test -cover ./...

Thanks ! Happy Unit Testing
