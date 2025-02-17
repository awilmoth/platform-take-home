package server

import (
	"context"
	"errors"

	"github.com/rs/cors"
	"github.com/skip-mev/platform-take-home/observability/logging"
	"go.uber.org/zap"

	"fmt"
	"net"
	"net/http"

	"github.com/skip-mev/platform-take-home/api/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func StartGRPCGateway(ctx context.Context, host string, port int) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
		}))

	if err := types.RegisterTakeHomeServiceHandlerFromEndpoint(ctx, mux, "localhost:9008", opts); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		return fmt.Errorf("error creating listener: %v", err)
	}

	// Create a new mux for both health check and gRPC gateway
	rootMux := http.NewServeMux()

	// Add health check endpoint
	rootMux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
			logging.FromContext(ctx).Error("failed to write health check response", zap.Error(err))
		}
	})

	// Add gRPC gateway endpoints
	corsMiddleware := cors.New(cors.Options{})
	grpcHandler := corsMiddleware.Handler(mux)
	rootMux.Handle("/", grpcHandler)

	server := http.Server{Handler: rootMux}

	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			logging.FromContext(ctx).Fatal("error shutting down http server", zap.Error(err))
		}
	}()

	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error serving http: %v", err)
	}

	return nil
}
