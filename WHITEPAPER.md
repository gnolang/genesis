# Gno: A Deterministic Multi-Realm Computing Platform

TODO: Add authors

Draft v0.1, December 2025

The current spec is on [github.com/gnolang/gno](https://github.com/gnolang/gno)

## Abstract

> At first, there was Bitcoin, out of entropy soup of the greater All.
> Then, there was Ethereum, which was created in the likeness of Bitcoin,
> but made Turing complete.
> Among these were Tendermint and Cosmos to engineer robust PoS and IBC.
> Then came Gno upon Cosmos and there spring forth Gnoland,
> simulated by the Gnomes of the Greater Resistance.

We stand at an inflection point in the evolution of decentralized systems. The
infrastructure that governs our digital lives, social networks, financial
systems, public discourse, and collective knowledge, remains captured by
centralized authorities that decide what we see, what we say, and what we own.
Blockchains promised liberation, yet their programming models remain opaque,
their governance plutocratic, and their developer experience hostile.

Gno.land is the missing piece: a platform for timeless code that stores truth
in plain sight.

Gno, the deterministic, Go-compatible programming environment powering the
Gno.land platform, provides a model where smart contracts exist as
human-readable source code: permanently on-chain, directly auditable, and
inherently forkable. Gno.land then builds on this foundation by introducing
governance driven not by token accumulation but by demonstrable contribution.
Together, they form the basis of the Logoverse: a persistent, composable,
censorship-resistant substrate for decentralized knowledge and coordination.

### Technical Summary

Problem: Current smart contract platforms exhibit: (1) bytecode opacity
preventing practical auditability, (2) manual state serialization causing
developer friction and storage-related vulnerabilities, (3) stake-weighted
governance conflating capital ownership with technical competence, and
(4) domain-specific languages fragmenting the developer ecosystem.

Solution: We introduce Gno, a deterministic Go-compatible language executed
by GnoVM, an AST-based interpreter providing automatic object persistence,
capability-based realm isolation, and source-level transparency. Governance
authority derives from Proof of Contribution (PoC), a Sybil-resistant
reputation mechanism assigning voting power based on verified contributions
rather than token holdings.

Result: Gno.land achieves: (1) source-level transparency with direct
auditability, (2) automatic state persistence eliminating serialization
overhead, (3) contribution-weighted governance resistant to plutocratic
capture, and (4) immediate developer accessibility for Go's established
developer base.

# 1. Introduction

## 1.1 Background

Bitcoin demonstrated decentralized value transfer without intermediaries[^1].
Ethereum extended this to general computation through smart contracts[^2].
Tendermint and Cosmos established practical Byzantine fault-tolerant consensus
and inter-blockchain communication[^3][^4].

Despite these advances, structural limitations persist:

- Opacity: Deployed contracts exist as bytecode, requiring specialized tools
  to interpret
- Complexity: Developers must learn domain-specific languages and manage state
  serialization manually
- Plutocracy: Governance power correlates with token holdings rather than
  expertise
- Fragmentation: Each ecosystem demands distinct languages, tooling, and
  mental models

## 1.2 Contributions

This whitepaper presents:

1. Gno Language (S4): Deterministic Go variant (~99% compatible) with
   source-level on-chain storage
2. GnoVM (S5): AST-based interpreter with automatic persistent object memory
3. Realm Model (S6): Capability-based isolation with explicit cross-realm
   semantics
4. Proof of Contribution (S9): Reputation-based consensus mechanism separating
   voting power from token ownership
5. Tendermint2 (S10): Minimal consensus engine where implementation serves as
   specification

# 2. System Model

TODO:

- Formal Foundations
- 4 Core Definitions: Realm tuple, Persistent Object Memory graph,
  Deterministic Execution, Cross-Realm Call semantics

# 3. Comparison with Existing Platforms

TODO:

- Why Gno Wins
- Bytecode problem + VM Comparison Table (EVM/WASM/MoveVM vs GnoVM)
- Solidity vs Gno code examples, e.g., State management
- Add more

## 3.1 The Bytecode Problem

Ethereum stores compiled bytecode:

```
Deployed: 0x608060405234801561001057600080fd5b50...
```

Interpreting this requires decompilation or external source verification.

Gno stores source code directly:

```go
func Transfer(to std.Address, amount int64) {
    caller := std.PreviousRealm().Address()
    balances[caller] -= amount
    balances[to] += amount
}
```

Auditing requires only the ability to read Go.

## 3.2 Virtual Machine Comparison

| Feature       | EVM        | WASM       | MoveVM     | GnoVM          |
|---------------|------------|------------|------------|----------------|
| On-chain      | Bytecode   | Bytecode   | Bytecode   | Source (AST)   |
| Auditability  | Decompiler | Decompiler | Decompiler | Direct         |
| State model   | Manual     | Manual     | Linear     | Auto-persistent|
| Safety model  | Manual     | Sandbox    | Resources  | Realm isolation|
| Determinism   | Partial    | Partial    | Strong     | Strong         |
| Language      | Solidity   | Rust       | Move       | Go             |

# 4. Gno Language

TODO:

- Language Specification
- ~99% Go compatible + Deterministic constraints + std package
- Render convention
- Interrealm specification

# 5. GnoVM Architecture

TODO:

- VM Architecture
- AST-based interpreter (not bytecode)
- Persistent Object Memory (automatic state)
- Execution lifecycle

# 6. Realm Execution Model

TODO:

- Execution Semantics
- OS Analogy Table (Process=Realm, Kernel=GnoVM)
- Package hierarchy (/p/, /r/, /e/)
- Isolation Theorem

# 7. State Commitment

In most smart contract platforms, developers manually serialize and deserialize
state using explicit storage slots or key-value mappings. Gno eliminates this
friction: global variables are automatically persistent.

```go
package counter

var total int

func Add(n int) {
    total += n
}

func Get() int {
    return total
}
```

When `Add(5)` is called, `total` changes from 0 to 5. GnoVM detects this change
and commits it to the Merkle tree automatically. No `storage.set()`, no manual
encoding, no boilerplate.

This works for any Go type: structs, slices, maps, pointers. The entire object
graph rooted at package-level variables is tracked and persisted.

TODO:

- Merkle Tree structure details
- Amino encoding (deterministic serialization)
- Gas costs for state changes

# 8. Economics

TODO:

- Token & Incentive Design
- GNOT Token + Gas model
- Storage Deposit (lock/unlock)
- Fee distribution: Validators + Contributors
- ref: https://gist.github.com/jaekwon/4d1c81ee3b82a0fc29f67d50e7f8664c

# 9. Proof of Contribution

TODO:

- Governance Innovation
- PoS limitations (capital != competence)
- Contribution based + Tier structure + DAOs
- Maybe separate section for GovDAO?
- ref: https://gist.github.com/jaekwon/918ad325c4c8f7fb5d6e022e33cb7eb3

# 10. Tendermint2

TODO:

- Consensus Layer
- Implementation = Specification + ~15K LOC
- BFT rounds (Propose -> Prevote -> Precommit)
- PoC integration
- Differences from the original Tendermint?

# 11. Safety and Liveness

TODO:

- Security Proofs
- Assumptions: network eventually delivers, <1/3 malicious, crypto works
- Theorems: all nodes compute same result, contracts can't hack each other,
  no forks, chain never halts, no censorship

# 12. The Logoverse

TODO:

- Long-term Vision
- Permanence + Composability + Transparency
- Core applications (boards, users, govdao, gnodev, etc.)
- ref: https://gno.land/r/gnoland/blog:p/gnoland-the-first-logoverse

# 13. Conclusion

TODO:

- Summary & Recap
- Go-compatible + AST transparency + auto-persistence
- PoC governance + minimal consensus

# Appendix

TODO: Add

# References

TODO: Add more

[^1]: S. Nakamoto, "Bitcoin: A peer-to-peer electronic cash system," 2008.
[^2]: V. Buterin, "Ethereum: A next-generation smart contract and decentralized
      application platform," 2014.
[^3]: J. Kwon, "Tendermint: Consensus without mining," 2014.
[^4]: J. Kwon, E. Buchman, "Cosmos: A network of distributed ledgers," 2016.

## Changelog
