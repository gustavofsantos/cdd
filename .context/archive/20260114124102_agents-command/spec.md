# Track: agents-command

## 1. User Intent
The product is pivoting to work mainly as two things: a CLI to coordinate work and Agent Skills to orchestrate AI.
Remove the `prompts` command and add a new command called `agents` with the flag `--install` that performs the skill installation.

## 2. Relevant Context
- `internal/cmd/prompts.go`: Contains the current implementation of the `prompts` command, including the `--install` logic.
- `internal/cmd/prompts_test.go`: Tests for the `prompts` command.
- `internal/cmd/prompts_install_test.go`: Tests for the `--install` flag of the `prompts` command.
- `internal/cmd/init.go`: References `cdd prompts --bootstrap`.

## 4. Scenarios
- Scenario: Remove prompts command
  Given the cdd CLI
  When I run `cdd prompts`
  Then it should return a command not found error

- Scenario: Install agent skill
  Given the cdd CLI
  When I run `cdd agents --install`
  Then it should install the CDD System Prompt as an Agent Skill in `.agent/skills/cdd/SKILL.md`

## 5. Relevant Files
- `internal/cmd/prompts.go`
- `internal/cmd/prompts_test.go`
- `internal/cmd/prompts_install_test.go`
- `internal/cmd/agents.go`
- `internal/cmd/agents_test.go`
- `internal/cmd/init.go`
- `README.md`
