package service

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	jsoncodec "github.com/unistack-org/micro-codec-json/v3"
	consulconfig "github.com/unistack-org/micro-config-consul/v3"
	fileconfig "github.com/unistack-org/micro-config-file/v3"
	httpsrv "github.com/unistack-org/micro-server-http/v3"
	"github.com/unistack-org/micro/controller"
	endpoints2 "github.com/unistack-org/micro/endpoints"
	"github.com/unistack-org/micro/handler"
	pb "github.com/unistack-org/micro/proto"
	"github.com/unistack-org/micro/serviceconfig"
	"github.com/unistack-org/micro/v3"
	"github.com/unistack-org/micro/v3/config"
	"github.com/unistack-org/micro/v3/server"
	"net/http"
)

func StartHTTPService(ctx context.Context, errCh chan<- error) {
	cfg := serviceconfig.NewConfig("test-service", "1.0")

	if err := config.Load(ctx,
		config.NewConfig(config.Struct(cfg)),
		config.NewConfig(
			config.AllowFail(true),
			config.Struct(cfg),
			config.Codec(jsoncodec.NewCodec()),
			fileconfig.Path("../local.json"),
		),
		consulconfig.NewConfig(
			config.AllowFail(true),
			config.Codec(jsoncodec.NewCodec()),
			config.BeforeLoad(func(ctx context.Context, conf config.Config) error {
				return conf.Init(
					consulconfig.Address("localhost:8500"),
					consulconfig.Path("api/test-service"),
				)
			}),
		),
	); err != nil {
		errCh <- err
	}

	options := append([]micro.Option{},
		micro.Servers(httpsrv.NewServer()),
		micro.Context(ctx),
		micro.Name(cfg.Server.Name),
		micro.Version(cfg.Server.Addr),
	)
	svc := micro.NewService(options...)

	err := svc.Init()
	if err != nil {
		errCh <- err
	}

	if err := svc.Init(
		micro.Servers(httpsrv.NewServer(
			server.Name(cfg.Server.Name),
			server.Version(cfg.Server.Version),
			server.Address(cfg.Server.Addr),
			server.Context(ctx),
			server.Codec("application/json", jsoncodec.NewCodec()),
		)),
	); err != nil {
		errCh <- err
	}
	panController := controller.NewPanController(cfg)

	router := mux.NewRouter()

	router.HandleFunc("/visa", panController.PanToCustomer).Methods("POST")
	http.Handle("/", router)
	err = http.ListenAndServe(cfg.Server.Addr, nil)
	if err != nil {
		errCh <- err
	}

	handler := handler.NewHandler(jsoncodec.NewCodec())

	endpoints := pb.NewPantoCustomerServiceEndpoints()

	if err := endpoints2.ConfigureHandlerEndpoints(router, handler, endpoints); err != nil {
		errCh <- err
	}

	router.NotFoundHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("Method not allowed, %v\n", req)
	})

	if err := svc.Server().Handle(svc.Server().NewHandler(router)); err != nil {
		errCh <- err
	}

	if err := svc.Run(); err != nil {
		errCh <- err
	}
}
