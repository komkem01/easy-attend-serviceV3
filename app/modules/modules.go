package modules

import (
	"log/slog"
	"sync"

	"mcop/internal/config"
	configDTO "mcop/internal/config/dto"
	"mcop/internal/database"
	"mcop/internal/log"
	"mcop/internal/otel/collector"

	"mcop/app/modules/entities"
	"mcop/app/modules/example"
	exampletwo "mcop/app/modules/example-two"
	appConf "mcop/config"
	// "mcop/app/modules/kafka"
)

type Modules struct {
	Conf *config.Module[appConf.Config]
	Log  *log.Module
	OTEL *collector.Module
	DB   *database.DatabaseModule
	ENT  *entities.Module
	// Kafka *kafka.Module
	Example  *example.Module
	Example2 *exampletwo.Module
}

func modulesInit() {
	confMod := config.New(&appConf.App)
	conf := confMod.Svc.Config()

	logMod := log.New(configDTO.Conf[log.Option](confMod.Svc))
	otel := collector.New(configDTO.Conf[collector.Config](confMod.Svc))
	log := log.With(slog.String("module", "modules"))
	log.Infof("otel module initialized")

	db := database.New(conf.Database.Sql)
	log.Infof("database module initialized")

	entitiesMod := entities.New(db.Svc.DB())
	log.Infof("entities module initialized")

	exampleMod := example.New(configDTO.Conf[example.Config](confMod.Svc), entitiesMod.Svc)
	log.Infof("example module initialized")

	exampleMod2 := exampletwo.New(configDTO.Conf[exampletwo.Config](confMod.Svc), entitiesMod.Svc)
	log.Infof("example module initialized")

	// kafka := kafka.New(&conf.Kafka)
	// log.Infof("kafka module initialized")

	mod = &Modules{
		Conf:     confMod,
		Log:      logMod,
		OTEL:     otel,
		DB:       db,
		ENT:      entitiesMod,
		Example:  exampleMod,
		Example2: exampleMod2,
	}

	log.Infof("all modules initialized")
}

var (
	once sync.Once
	mod  *Modules
)

func Get() *Modules {
	once.Do(modulesInit)

	return mod
}
