package prompts

import _ "embed"

//go:embed system.md
var System string

//go:embed analyst.md
var Analyst string

//go:embed architect.md
var Architect string

//go:embed executor.md
var Executor string

//go:embed integrator.md
var Integrator string
