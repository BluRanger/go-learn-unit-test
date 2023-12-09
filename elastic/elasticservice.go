//go:generate mockgen -source=elastic/elasticservice.go -destination=mocks/mock_elasticservice.go -package=mocks
package elastic

type IElasticService interface {
	ESscroll(string) string
	Method2(string) string
}

type ElasticService struct {
	test string
	num  int
}

func (ElasticService) ESscroll(test string) string {
	return "esscroll"
}

func (ElasticService) Method2(test string) string {
	return "method2"
}
