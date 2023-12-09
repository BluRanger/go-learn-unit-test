package service

import "sample/elastic"

var elasticServiceFactory = elastic.NewInstance

func SetElasticServiceFactory(factory func(x, y, z, a string) elastic.IElasticService) {
	elasticServiceFactory = factory
}

func ResetElasticServiceFactory() {
	elasticServiceFactory = elastic.NewInstance
}

func Job() {
	instance := elasticServiceFactory("a", "a", "a", "a")
	instance.ESscroll("hello")
	instance.Method2("hello")

	instance1 := elasticServiceFactory("b", "b", "b", "b")
	instance1.ESscroll("hello")
	instance1.Method2("hello")
}
