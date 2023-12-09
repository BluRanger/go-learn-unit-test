package elastic

type Elastic struct {
}

func NewInstance(x, y, z, a string) IElasticService {
	return ElasticService{test: "1", num: 2}
}
