package events

import (
	"context"
	"fmt"
	"glass/dao/model"
)

type _AccountNamespace struct{}

var Account _AccountNamespace

func (_AccountNamespace) AfterRegister(ctx context.Context, uid int64) {
	fmt.Printf("user registered: %d\n", uid)
}

func (_AccountNamespace) BeforeUnregister(ctx context.Context, user *model.User) {}

func (_AccountNamespace) AfterLogin(ctx context.Context, uif int64) {}
