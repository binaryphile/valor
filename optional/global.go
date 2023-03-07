package optional

import "time"

var (
	NoBool     = Bool{}
	NoDuration = Duration{}
	NoString   = String{}
)

type (
	Bool     = Value[bool]
	Duration = Value[time.Duration]
	String   = Value[string]
)
