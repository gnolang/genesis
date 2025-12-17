# Gno: A Deterministic Multi-Realm Computing Platform

## Abstract
TODO: Concise summary of what Gno is, the problem it addresses, and its core differentiators.

## Introduction
TODO: Motivation for Gno, limitations of existing blockchain platforms, and high-level goals.

## Gno

### The Gno Language

Gno is approximately 99% compatible with Go 1.17. Developers write standard Go
code with minimal modifications. The language enforces determinism through
careful restrictions on non-deterministic operations.

Determinism constraints:

- `time.Now()` returns block time, not system time
- No goroutines, channels, or concurrency primitives
- No OS, network, or filesystem access
- No reflection or unsafe operations
- TODO: continue this list

All standard Go types are supported: bool, int, uint, float32, float64, string,
arrays, slices, maps, structs, pointers, and functions.

TODO: Introduced keywords (address, etc).
TODO: chain/* stdlibs.
TODO: why Go is the chosen abstraction.  
TODO: Document security implications of language-level constraints.

### The Gno Virtual Machine

Unlike traditional VMs that execute bytecode, GnoVM interprets the Abstract
Syntax Tree (AST) directly. Source code is parsed and stored on-chain as AST
nodes, making the execution model transparent and auditable.

When code is deployed:

1. Source is parsed into an AST
2. AST is type-checked and validated
3. AST nodes are stored on-chain (not compiled bytecode)
4. Execution walks the AST directly

This approach eliminates the compilation step that obscures intent in other
platforms. What you read is what executes.

TODO: Explain how GnoVM executes Gno code, its interpreter model, and how it differs from bytecode-based VMs.  
TODO: Describe how determinism, isolation, and safety are enforced by the VM.

### Execution Environment
TODO: Describe how transactions are executed, lifecycle of execution, and runtime guarantees.  
TODO: Clarify how failures, reverts, and safety boundaries are handled.

### State and Persistence

Global variables are automatically persistent:

```go
package counter

var total int

func Add(n int) {
    total += n
}
```

When `Add(5)` is called, `total` changes from 0 to 5. GnoVM detects this change
and commits it to the Merkle tree automatically. No `storage.set()`, no manual
encoding, no boilerplate.

This works for any Go type: structs, slices, maps, pointers. The entire object
graph rooted at package-level variables is tracked and persisted.

TODO: Explain how state is represented, how persistent objects are tracked automatically, and how changes are committed.  
TODO: Describe persistence guarantees and security implications of automatic state tracking.

### Gas and Resource Accounting
TODO: Define how computation and storage are metered, what resources are limited, and how costs are calculated.  
TODO: Describe protections against resource exhaustion and abuse.

### Governance and Economics
TODO: Describe the governance model (including Proof of Contribution) and economic incentives.  
TODO: Explain how governance mitigates capture, abuse, and long-term risk.

### Consensus
TODO: Explain how nodes reach agreement on execution results and state, including the role of Tendermint2.  
TODO: Document assumptions and safety guarantees of the consensus layer.

### Interoperability
TODO: Describe how Gno interacts with external chains and systems, including IBC.  
TODO: Describe trust boundaries and security considerations for cross-chain interactions.

## Applications and Vision
TODO: Describe the long-term vision for Gno, including the Logoverse and core applications.  
TODO: Explain how the design choices enable safe, composable, long-lived applications.

## Conclusion
TODO: Summarize the system and restate its core principles and goals.
