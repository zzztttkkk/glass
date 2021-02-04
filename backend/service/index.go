package service

import (
	"context"
	"glass/internal"
	"glass/service/account"
)

var Account account.Namespace

func BuiltTime(_ context.Context) string {
	return internal.BuiltTime
}
