package middleware

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func logRequest(ctx context.Context, method string, err error) {
	clientIP := "unknown"
	if p, ok := peer.FromContext(ctx); ok {

		clientIP = p.Addr.String()
	}

	logEntry := fmt.Sprintf("IP: %s, Method: %s,", clientIP, method)

	if err != nil {
		st, _ := status.FromError(err)
		log.Println(logEntry, "Code:", st.Code(), st.Message())
	} else {
		log.Println(logEntry, "Status: Success")
	}
}

func UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)
	logRequest(ctx, info.FullMethod, err)
	return res, err
}
