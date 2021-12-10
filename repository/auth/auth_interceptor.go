package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

// AuthInterceptor -
type AuthInterceptor struct {
	jwtManager            *JWTManager
	accessiblePermissions map[string][]string
}

// NewAuthInterceptor -
func NewAuthInterceptor(jwtManager *JWTManager, accessiblePermissions map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessiblePermissions}
}

// Unary -
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		if err := interceptor.authorize(ctx, info.FullMethod); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	log.Println("Starting authorize...")
	accessiblePermissions, ok := interceptor.accessiblePermissions[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("metadata is not provided")
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		log.Println("authorization token is not provided")
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		log.Println("access token is invalid")
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, userPermission := range claims.Permissions {
		for _, permission := range accessiblePermissions {
			if permission == userPermission {
				return nil
			}
		}
	}

	log.Println("no permission to access this RPC")
	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
