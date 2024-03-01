package app

import (
	"basic-backend/internal/config"
	"context"
	"log"
	"net"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initApp(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) initApp(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	log.Println("init service provider")

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("./config/config.yaml")
	if err != nil {
		return err
	}

	log.Println("init config")

	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {
	mux := http.NewServeMux()

	userHandlers := a.serviceProvider.UserAPI().Handlers()
	for pattern, _ := range userHandlers {
		mux.HandleFunc(pattern, a.serviceProvider.UserAPI().Get)
	}

	a.httpServer = &http.Server{
		Handler: mux,
	}

	log.Println("init http server")

	return nil
}

func (a *App) runHTTPServer() error {
	l, err := net.Listen("tcp", a.serviceProvider.HTTPConfig().Address())
	if err != nil {
		return err
	}

	log.Println("run http server")

	return a.httpServer.Serve(l)
}
