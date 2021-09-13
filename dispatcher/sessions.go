package dispatcher

type Cookie string

// Sessions returns the bots that the user is authorized to manage
var (
	Sessions = make(map[string]string)
)
