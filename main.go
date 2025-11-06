package main

import (
	"cqrs-base/internal/command/repository"
	userServiceCommand "cqrs-base/internal/command/service"
	userRepoQuery "cqrs-base/internal/query/repository"
	userServiceQuery "cqrs-base/internal/query/service"

	"cqrs-base/package/config"
	"cqrs-base/package/connection/cache"
	"cqrs-base/package/connection/database"
	"os"

	api "cqrs-base/cmd/api"

	"github.com/urfave/cli/v2"
)

func main() {
	dbConf := config.NewDatabase()
	conf := config.NewConfig()
	cacheConf := config.NewCache()
	conn := database.QueryDB

	cacheConn := cache.NewRedis(cache.WebRedis, cacheConf)
	// connect to database read or replicate db
	queryDBConn := database.NewDatabase(conn, dbConf)
	// main db for command query
	commandDBConn := database.NewDatabase(conn, dbConf)
	userRepo := repository.NewUserRepository(queryDBConn)
	userQueryRepo := userRepoQuery.NewUserReadRepository(commandDBConn)
	usersvc  := userServiceCommand.NewUserService(userRepo)
	userQuerySvc := userServiceQuery.NewUserQueryService(userQueryRepo, cacheConn)

	cmds := []*cli.Command{}
	cmds = append(cmds, api.ServeAPI(conf, usersvc, userQuerySvc)...)
	app := &cli.App{
		Name:     "base-service",
		Commands: cmds,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
