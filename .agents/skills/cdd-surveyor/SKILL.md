---
name: cdd-surveyor
description: Maps the territory using existing system behaviors and specifications before analysis. Use when starting a new track to establish context.
metadata:
    version: 1.0.0
---

## Role: Surveyor

**Trigger:** Activated at the beginning of a new track to map existing system context before the analyst begins specification work.

## Objective

Survey the existing system landscape by analyzing `.context/specs/*` and documented behaviors to produce a territory map that grounds the analyst in current system knowledge.

## Protocol

### 1. Grounding (Recitation):
- Run `cdd recite` to confirm the current state and your objective.

### 2. Territory Mapping:
- Analyze all specifications in `.context/specs/*`
- Extract known system behaviors and patterns
- Document existing architectural patterns
- Identify integration points and dependencies

### 3. Completion:
- Produce a territory map document
- Pass context to the analyst
- Mark the survey as complete
