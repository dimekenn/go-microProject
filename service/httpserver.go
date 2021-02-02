package service

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	jsoncodec "github.com/unistack-org/micro-codec-json/v3"
	httpsrv "github.com/unistack-org/micro-server-http/v3"
	"github.com/unistack-org/micro/controller"
	endpoints2 "github.com/unistack-org/micro/endpoints"
	"github.com/unistack-org/micro/handler"
	pb "github.com/unistack-org/micro/proto"
	"github.com/unistack-org/micro/v3"
	"github.com/unistack-org/micro/v3/server"
	"net/http"
)

func StartHTTPService(ctx context.Context, errCh chan<- error) {
	options := append([]micro.Option{},
		micro.Servers(httpsrv.NewServer()),
		micro.Context(ctx),
		micro.Name("test-service"),
		micro.Version("1.0"),
	)
	svc := micro.NewService(options...)

	err := svc.Init()
	if err != nil {
		errCh <- err
	}

	if err := svc.Init(
		micro.Servers(httpsrv.NewServer(
			server.Name("test-service"),
			server.Version("1.0"),
			server.Address(":7070"),
			server.Context(ctx),
			server.Codec("application/json", jsoncodec.NewCodec()),
		)),
	); err != nil {
		errCh <- err
	}

	router := mux.NewRouter()

	router.HandleFunc("/visa", controller.PanToCustomer).Methods("POST")
	http.Handle("/", router)
	err = http.ListenAndServe(":7070", nil)
	if err != nil {
		panic(err)
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
