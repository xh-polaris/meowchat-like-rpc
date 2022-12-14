package errorx

import "google.golang.org/grpc/status"

var (
	ErrDataBase   = status.Error(10002, "database error")
	ErrNoThisItem = status.Error(10003, "no this item")
)
