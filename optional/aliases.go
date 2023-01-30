package optional

import "time"

type (
	Bool     = Value[bool]
	Duration = Value[time.Duration]
	String   = Value[string]
)
