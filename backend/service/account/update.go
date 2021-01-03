package account

import (
	"context"
)

func (Namespace) DoChangePasswordNormally(ctx context.Context, newPassword []byte) {

}

func (Namespace) DoChangePasswordBySecret(ctx context.Context, name, secret, newPassword []byte) {

}

func (Namespace) DoChangeName(ctx context.Context, name []byte) {
}

func (Namespace) DoChangeAlias(ctx context.Context, alias []byte) {
}
