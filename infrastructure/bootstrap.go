package infrastructure

import (
	"hulk/go-webservice/infrastructure/persist"
	"hulk/go-webservice/infrastructure/services"
)

func RegisterAppServices() {
	persist.InitDB()
	services.InitSchedulerService()
	services.InitPubSubService()
	services.InitRedisService()
	services.InitPubSubService()
	services.InitFTSService()
}
