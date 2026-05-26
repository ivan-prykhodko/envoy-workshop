package currency

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ivan-prykhodko/envoy-workshop/protogen/go/currency"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	currency.ExchangeRateServiceServer

	port                int
	server              *grpc.Server
	exchangeRateService *ExchangeRateService
}

func NewGrpcServer(port int, exchangeRateService *ExchangeRateService) *GrpcServer {
	return &GrpcServer{
		port:                port,
		exchangeRateService: exchangeRateService,
	}
}

func (s *GrpcServer) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d : %v\n", s.port, err)
	}
	log.Printf("Server listening on port %d\n", s.port)

	grpcServer := grpc.NewServer()
	s.server = grpcServer

	currency.RegisterExchangeRateServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC on port %d : %v\n", s.port, err)
	}
}

func (s *GrpcServer) Stop() {
	s.server.Stop()
}

func (s *GrpcServer) GetExchangeRate(ctx context.Context, req *currency.ExchangeRateRequest) (*currency.ExchangeRateResponse, error) {
	fromCurrency, toCurrency := req.GetFromCurrency(), req.GetToCurrency()

	rate, err := s.exchangeRateService.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &currency.ExchangeRateResponse{
		ExchangeRate: &currency.ExchangeRate{
			FromCurrency: fromCurrency,
			ToCurrency:   toCurrency,
			Rate:         rate,
			UpdatedAt:    timestamppb.Now(),
		},
	}, nil
}
