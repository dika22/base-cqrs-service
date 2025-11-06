package cmd

import (
	"context"
	"cqrs-base/internal/command/service"
	"cqrs-base/package/config"
	"fmt"
	"net/http"
	"os"
	"time"

	"os/signal"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"

	commandDelivery "cqrs-base/internal/command/delivery"
	queryDelivery "cqrs-base/internal/query/delivery"
	userServiceQuery "cqrs-base/internal/query/service"
)

const CmdServeHTTP = "serve-http"

type HTTP struct {
	cfg *config.Config
	usersvc *service.UserService
	userQuerySvc *userServiceQuery.UserQueryService
}

func (h HTTP) ServeAPI(c *cli.Context) error {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	userAPI := e.Group("api/v1/users")

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	

	commandDelivery.NewUserHandler(userAPI, h.usersvc)
	queryDelivery.NewUserQueryHandler(userAPI, h.userQuerySvc)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", h.cfg.ServerPort)); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	return nil
}

func ServeAPI(cfg *config.Config, usersvc *service.UserService, userQuerySvc *userServiceQuery.UserQueryService) []*cli.Command {
	h := &HTTP{cfg: cfg, usersvc: usersvc}
	return []*cli.Command{
		{
			Name:   CmdServeHTTP,
			Usage:  "Serve Document Service",
			Action: h.ServeAPI,
		},
	}
}
