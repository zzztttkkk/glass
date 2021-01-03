package captcha

import "context"

func (Namespace) Rand(ctx context.Context) string {
	var ret string
	_ = op.RowColumns(ctx, "txt", "where id>=FLOOR(RAND()*MAX(id)) limit 1", &ret)
	return ret
}
