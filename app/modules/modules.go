package modules

import (
	"log/slog"
	"sync"

	"github.com/easy-attend-serviceV3/internal/config"
	configDTO "github.com/easy-attend-serviceV3/internal/config/dto"
	"github.com/easy-attend-serviceV3/internal/database"
	"github.com/easy-attend-serviceV3/internal/log"
	"github.com/easy-attend-serviceV3/internal/otel/collector"

	"github.com/easy-attend-serviceV3/app/modules/classroom"
	"github.com/easy-attend-serviceV3/app/modules/entities"
	"github.com/easy-attend-serviceV3/app/modules/example"
	exampletwo "github.com/easy-attend-serviceV3/app/modules/example-two"
	"github.com/easy-attend-serviceV3/app/modules/gender"
	"github.com/easy-attend-serviceV3/app/modules/prefix"
	"github.com/easy-attend-serviceV3/app/modules/school"
	appConf "github.com/easy-attend-serviceV3/config"
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

	Gender    *gender.Module
	Prefix    *prefix.Module
	School    *school.Module
	Classroom *classroom.Module
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

	genderMod := gender.New(entitiesMod.Svc)
	log.Infof("gender module initialized")

	prefixMod := prefix.New(entitiesMod.Svc)
	log.Infof("prefix module initialized")

	schoolMod := school.New(entitiesMod.Svc)
	log.Infof("school module initialized")

	classroomMod := classroom.New(entitiesMod.Svc)
	log.Infof("classroom module initialized")

	// kafka := kafka.New(&conf.Kafka)
	// log.Infof("kafka module initialized")

	mod = &Modules{
		Conf:      confMod,
		Log:       logMod,
		OTEL:      otel,
		DB:        db,
		ENT:       entitiesMod,
		Example:   exampleMod,
		Example2:  exampleMod2,
		Gender:    genderMod,
		Prefix:    prefixMod,
		School:    schoolMod,
		Classroom: classroomMod,
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
