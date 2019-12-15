package ls

import "context"

type RequestContext struct {
	Server  Server
	Context context.Context
}
