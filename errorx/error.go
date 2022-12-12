package errorx

import "google.golang.org/grpc/status"

var (
	ErrExample = status.Error(10001, "example error")
)
