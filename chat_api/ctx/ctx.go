package ctx

import (
	"chat/pkg/http/engine"
	"chat/pkg/log"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"

	grpcClient "chat/pkg/grpc/client"
	"github.com/sirupsen/logrus"
)

type ServiceContext interface {
	Conf() *Conf
	Logger() *logrus.Logger
	Start(ginEngine *engine.GinEngine)
	RailsGrpcClient() *grpc.ClientConn
}

type defaultServiceContext struct {
	conf            *Conf
	logger          *logrus.Logger
	railsGrpcClient *grpc.ClientConn
}

func NewDefaultServiceContext() ServiceContext {
	op := "NewDefaultServiceContext"
	conf := loadConf()
	logger := log.NewLogger(conf)

	railsGrpcClient, err := grpcClient.NewGrpcClient(conf.railsGrpcHost, conf.railsGrpcPort)
	if err != nil {
		logger.Fatalf("%s.railsGrpcClient: %s", op, err.Error())
	}

	ctx := &defaultServiceContext{
		conf:            conf,
		logger:          logger,
		railsGrpcClient: railsGrpcClient,
	}

	return ctx
}

func (ctx *defaultServiceContext) Start(ginEngine *engine.GinEngine) {
	ctx.ShutdownHook(ginEngine.Shutdown)
	ginEngine.RunHttpServer()
}

func (ctx *defaultServiceContext) ShutdownHook(shutdownFunctions ...func() error) {
	go func() {
		ctx.Logger().Info("installing shutdown hook")
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		ctx.Logger().Info("received shutdown hook")
		ctx.shutdown()
		for i, function := range shutdownFunctions {
			err := function()
			if err != nil {
				ctx.Logger().Warnf("#%d error while trying to shutdown: %v", i, err)
			}
		}
	}()
}

func (ctx *defaultServiceContext) Conf() *Conf {
	return ctx.conf
}

func (ctx *defaultServiceContext) Logger() *logrus.Logger {
	return ctx.logger
}

func (ctx *defaultServiceContext) shutdown() {
	ctx.logger.Info("shutting down")
}

func (ctx *defaultServiceContext) RailsGrpcClient() *grpc.ClientConn {
	return ctx.railsGrpcClient
}
