package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
		//log.Println("--> unary interceptor: ", info.FullMethod)

		if err := interceptor.authorize(ctx, info.FullMethod); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessiblePermissions, ok := interceptor.accessiblePermissions[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, userPermission := range claims.Permissions {
		for _, permission := range accessiblePermissions {
			if permission == userPermission {
				return nil
			}
		}
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
