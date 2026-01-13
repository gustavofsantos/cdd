package prompts

import _ "embed"

//go:embed system.md
var System string

//go:embed bootstrap.md
var Bootstrap string

//go:embed inbox.md
var Inbox string

//go:embed executor.md
var Executor string

//go:embed planner.md
var Planner string
