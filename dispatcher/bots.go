package dispatcher

import "github.com/discless/discless/types"

var (
	Manager = make(map[string]*types.Self)
)
