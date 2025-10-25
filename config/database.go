package config

import (
	dbdto "github.com/easy-attend-serviceV3/internal/database/dto"
	rddto "github.com/easy-attend-serviceV3/internal/redis/dto"
)

type Database struct {
	Sql   map[string]*dbdto.Option
	Redis map[string]*rddto.Option
}

var (
	database = Database{
		Sql: map[string]*dbdto.Option{"": {
			Host:     "127.0.0.1", // 127.0.0.1
			Port:     5432,        // 5432
			Database: "postgres",  // postgres
			Username: "postgres",  // postgres
			Password: "",
			TimeZone: "Asia/Bangkok", // Asia/Bangkok
		}},
		Redis: map[string]*rddto.Option{"": {
			Db:       0,
			Addr:     "",
			Username: "",
			Password: "",
		}},
	}

	// Database For Test
	databaseTest = Database{
		Sql: map[string]*dbdto.Option{"": {
			Host:     "", // 127.0.0.1
			Port:     0,  // 5432
			Database: "", // postgres
			Username: "", // postgres
			Password: "",
			TimeZone: "", // Asia/Bangkok
		}},
		Redis: map[string]*rddto.Option{"": {
			Db:       0,
			Addr:     "",
			Username: "",
			Password: "",
		}},
	}
)
