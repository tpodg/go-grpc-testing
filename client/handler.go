package main

import (
	"context"
	"crypto/tls"
	"github.com/tpodg/go-grpc-testing/client/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"time"
)

type handler struct {
	router     *http.ServeMux
	grpcClient *grpcClient
	restClient *restClient
}

func newHandler(cfg cfg) handler {
	var opts []grpc.DialOption
	if cfg.Grpc.Tls {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.Dial(cfg.Grpc.Target, opts...)
	if err != nil {
		log.Fatal("failed to dial", err)
	}

	h := handler{
		router: http.NewServeMux(),
		grpcClient: &grpcClient{
			pb.NewDemoServiceClient(conn),
		},
		restClient: &restClient{
			Client: &http.Client{
				Timeout: 3 * time.Second,
			},
			url: cfg.Rest.Target,
		},
	}
	h.router.HandleFunc("/grpc", h.sendGrpc)
	h.router.HandleFunc("/rest", h.sendRest)

	return h
}

func (h handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.router.ServeHTTP(rw, req)
}

func (h handler) sendRest(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving %s /restCfg", r.Method)

	res, err := h.restClient.Get(h.restClient.url + "/rest")
	if err != nil {
		log.Println("error occurred:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("error reading body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		log.Println("response:", string(body))
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func (h handler) sendGrpc(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving %s /grpcCfg", r.Method)

	res, err := h.grpcClient.Send(context.Background(), &pb.Request{Value: r.URL.Query().Get("value")})

	if err != nil {
		log.Println("error occurred:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	} else {
		log.Println("response:", res)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res.String()))
	}
}
