package s2s

import "context"

type S2SContextKey struct{}

type S2SRequestContext struct {
	TenantID string
	UserID   string
	Roles    []string
}

func FromContext(ctx context.Context) *S2SRequestContext {
	if v, ok := ctx.Value(S2SContextKey{}).(*S2SRequestContext); ok {
		return v
	}
	return nil
}
