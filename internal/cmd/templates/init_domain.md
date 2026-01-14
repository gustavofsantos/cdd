# Domain Model

## 1. Ubiquitous Language
Definition: The specific vocabulary used by the team. These terms must be used precisely in code and specs.
Format: Term: Definition. Constraints or Invariants.

* **[Term]:** ...

## 2. Domain Events (Event Storming)
**Definition:** Things that happen in the business that we care about.
Format: [Past Tense Verb], e.g., `OrderPlaced`, `PaymentFailed`.

* `[EventName]`

## 3. Bounded Contexts
**Definition:** The logical boundaries within the system.

* **Core Domain:** [The area where the business makes money/value]
* **Supporting Subdomains:** [Necessary logic that isn't the core differentiator]
* **Generic Subdomains:** [Commodity functionality, e.g., Auth, Payments]