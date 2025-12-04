# Gno.land Litepaper

Draft v0.1, December 2025

## What is Gno?

Gno.land is a platform for timeless code that stores truth in plain sight.

Gno is a deterministic, Go-compatible programming environment. Smart contracts
exist as human-readable source code: permanently on-chain, directly auditable,
and inherently forkable. Governance is driven by contribution, not tokens.

## The Problem

- **Bytecode opacity**: contracts are compiled, requiring tools to interpret
- **Manual state management**: developers serialize state by hand
- **Plutocratic governance**: voting power follows token holdings
- **Ecosystem fragmentation**: each platform has its own language and tools

## The Solution

**Gno Language**: a deterministic Go variant (~99% compatible) with source-level
on-chain storage. Write in a familiar language; audit exactly what executes.

**GnoVM**: an AST-based interpreter with automatic persistent object memory.
State management is transparent; developers focus on logic.

**Proof of Contribution**: a reputation-based governance mechanism. Voting power
comes from verified contributions, not capital.

## Source Transparency

Ethereum stores compiled bytecode:

```
0x608060405234801561001057600080fd5b50...
```

Gno stores source code directly:

```go
func Transfer(to std.Address, amount int64) {
    caller := std.PreviousRealm().Address()
    balances[caller] -= amount
    balances[to] += amount
}
```

## Platform Comparison

| Feature       | EVM        | WASM       | MoveVM     | GnoVM        |
|---------------|------------|------------|------------|--------------|
| On-chain      | Bytecode   | Bytecode   | Bytecode   | Source (AST) |
| Auditability  | Decompiler | Decompiler | Decompiler | Direct       |
| State model   | Manual     | Manual     | Linear     | Auto         |
| Determinism   | Partial    | Partial    | Strong     | Strong       |
| Language      | Solidity   | Rust       | Move       | Go           |

## Learn More

- https://docs.gno.land
- https://github.com/gnolang/gno
- https://gno.land
