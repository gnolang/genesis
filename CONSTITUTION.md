# [DRAFT] gno.land Constitution

>---
>
>**DOCUMENT NOTES**
>
>- This is a v0 working document. It is very early and rough around the edges. The goal here is to introduce an early working prototype that can live in [gnolang/genesis](https://github.com/gnolang/genesis) as a v0 document.
>- This prototype takes heavy influence from the [ATOMONE CONSTITUTION](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md).
>- Public discussions on this topic can be found int the github issue [gno.land Constitution #3080](https://github.com/gnolang/gno/issues/3080)
>
> **GOALS**
>
>- citizens, gnomes, should feel protected by this constitution
>- Constitution is about protecting long-term value of the chain
>- Quality over quantity
>- We are creating a GovDAO to protect the chain and keep it running for the very long term. Building a team anyone can trust and rely on.
>- Ensure that the skeleton is in place and placeholders are clear
>- Its ok to remove things as needed for v1.
>
>---

## Preamble

We, the people of gno.land, in order to create a more transparent, innovative, and decentralized world, do ordain and establish this Constitution for gno.land.

### Section 1: Declaration of Intent

>TODO: Outline the intent of this, we should be able to utilize prior work from those who contributed to Why Gno and other blog posts.

The vision of gno.land is...

### Section 2: Fundamental Principles

This Constitution of gno.land, hereinafter “Constitution”, hereby establishes the foundations of the rights, responsibilities, amd governance structure of those involved with of gno.land.

All subsequent governance proposals shall align with the provisions of this Constitution, and proponents of each proposal, along with all active governance voters, are required to ensure consistency between such proposals and this Constitution, and cannot violate any explicit restriction.

Amendments are permitted to innovate and adapt this Constitution, but they must still respect and adhere to the fundamental principles of this Constitution and are subject to specific procedures as described in Article 5: Amendments.

### Section 3: General Mission and Objectives

> TODO: Add mission and objectives

The mission of gno.land is to establish a platform focused on transparent governance, user agency and authentic content.

The objectives of gno.land are:

- Develop a governance model based on contributions
- Encourage forking as a tool for decentralization and innovation
- Reward authentic content and content moderation
- Build spam and manipulation resilience

## Article 1: The GovDAO

>TODO: Defines the GovDAO as the overarching governance body responsible for setting and upholding shared principles, ensuring transparency, and guiding the vision of the community.

## Article 2: Sub-DAOs

> TODO: Allows for the creation and recognition of specialized sub-DAOs that operate within the GovDAO’s framework, each managing its own domain while remaining accountable to the GovDAO’s foundational principles.

## Article 3: Citizen Rights and Responsibilities

> TODO: Affirms that each recognized citizen holds equal rights to participation, proposal, and vote, while maintaining responsibilities to uphold the constitution, act ethically, and foster a collaborative, merit-based environment.

### Citizen rights

- The freedom to fork and innovate
- The freedom to participate in transparent governance
- The right to access the decentralized, censorship resistant platform

### Citizen obligations

- To contribute to the community development and moderation
- To uphold transparency and integrity

## Article 4: Governance

> TODO: Establishes clear, minimal procedures for proposing initiatives, reaching consensus, and executing decisions, ensuring that all actions are transparent, auditable, and in alignment with the GovDAO’s mission.

>---
>**TODO**
>- This is modified from AtomOne, but needs to be checked for accuracy.
>- There may be governance items currently missing.
>- Outline relationship in GovDAOs (worksDAO, GovDAO, OversightDAO, ConstitutionDAO, etc)
>---

### Section 1: Procedures

The working language of gno.lang Governance is International English.

The governance process shall be defined through law, but will strictly adhere to all restrictions and provisions of this Constitution.

The vote types are YES, NO, and ABSTAIN.

To prevent spam and to ensure the quality of proposals, proposals must have a sufficient amount of burn deposit before voting can begin. The minimum amount of burn deposit needed shall self-adjust to target on average 1 proposal per 2 week period.

The quorum necessary for a proposal to be valid is 40%. The denominator shall be the number of bonded GNOT tokens.

A Simple Majority in gno.land Governance is defined to be exactly “more than half”. No proposal is considered to have passed in gno.lad with a Simple Majority; a Supermajority or Constitutional Majority is required. A Supermajority in gno.land Governance is defined to be exactly "more than two thirds". This is distinct from a Supermajority in DAO council governance which is defined to be exactly “two thirds or more”. For clarity, 3 out of 4 votes is necessary for a Supermajority in gno.land Governance, but 2 out of 3 is sufficient for a Supermajority in DAO council voting. The Constitutional Majority threshold is defined to be exactly “more than 90%”. These definitions of majority cannot change even by a Constitutional Majority.

The minimum voting period for a proposal is 3 weeks. To ensure adequate representation and participation, the governance process must extend the voting deadline to guarantee a minimum of 2 weeks of voting after the minimum quorum has been met.

### Section 2: Laws

>---
>**TODO**
>
>- Largely carried over from AtomOne
>- Goal: explain that GovDAO should have things like super majority for decisions.
>- Goal: OK to request schism with 40% of vote. (maybe TBD)
>- Goal: chain upgrade
>- Goal: parameter tuning
>- Goal: List Cases and set of requirements. 
>
>---

This Constitution allows the creation and enactment of laws as means to complement and enhance its content, restrictions, provisions and established processes.

Laws shall be enacted to address specific governance, economic, and operational aspects, and must align with the principles and framework established by this Constitution. In cases where laws conflict with this Constitution, this Constitution shall always prevail. The gno.land Governance structure will include several Decentralized Autonomous Organizations (DAOs) and sub-DAOs. Each DAO may create its own bylaws to govern its operations, provided they do not conflict with this Constitution, overarching laws or any mandates issued by the gno.land Governance. Mandates issued by an ancestor of a DAO are binding and cannot be amended or repealed by those DAOs or their sub-DAOs and remain in effect unless modified or repealed. Any DAO bylaws that conflict with its Governing Documents shall be deemed invalid and abrogated.

Laws may be proposed by any Citizen of gno.land through the established governance process. The approval of new laws, amendment of existing laws require a Constitutional Majority vote. However, an endorsement from the Steering DAO shall lower the passing threshold to a Supermajority vote for all proposals except for Constitutional amendment proposals, which shall require a Constitutional Majority.

Once a law is approved by the required majority, it is implemented according to the specified procedures. The implementation process includes updating this Constitution, communicating the changes to the community, and ensuring compliance with the new provisions.

Laws must be named plaintext files or folders of plaintext files, or folders of folders.

### Section 3: The gno.land Chain

>---
>**NOTES**
>- This should cover the unique characteristics gno.land such as how Proof of Contribution, Validators, and DAOs work related to this.
>---

The gno.land chain is the foundational infrastructure enabling decentralized applications and transparent governance. It establishes the framework for Proof of Contribution, validator selection, and DAOs. This section defines the unique characteristics of the gno.land chain and its core principles:  

#### Proof of Contribution

- Governance is driven by meaningful contributions rather than financial influence.
- Contributions include creating content, moderating platforms, developing software, and maintaining the chain's security. (TO BE CONFIRMED)

#### Validator Selection

- Validators are chosen based on technical capability and alignment with gno.land’s vision, not purely on capital stake.
- Validators must uphold the principles of decentralization, transparency, and security.

#### GovDAO and Sub-DAOs

- The chain supports GovDAO, the principal governing DAO, and any number of DAOs & sub-DAOs for governance and operations.
- Each DAO operates under the principles of transparency and collective decision-making.

### Section 4: Updates to gno.land

>---
>**TODO**
>- largely [ported from ATOMONE](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md#section-6-special-purpose-daos)
>- Review for correctness and add gno.land specific rules.
>---

All updates to gno.land must be proposed as distinct, independent components. Each component must be discussed and voted on separately to ensure thorough deliberation and transparency. Proposals shall not bundle different changes into a single submission. Each proposal must address a single, clearly defined update or change. The only exception is made to regular code maintenance and bug fixes, and critical security updates.

Adequate time must be provided between the consideration of each proposal to allow for comprehensive review and community input. The minimum duration for discussion and voting on each proposal shall be two weeks.

The passing of Software Update Proposals shall require the passage of law.

[TODO: name stakeholders from governance structure] are responsible and may be held accountable for voting YES only on software update proposals that make correct changes toward the living constitution and laws (any substantial changes to the software must first be reflected in the constitution and laws). This responsibility exists independently of the Oversight DAOs’ regulatory function.

## Article 5: Amendments

> TODO: Provides a simple mechanism for introducing, debating, and ratifying amendments to this constitution, ensuring it can evolve thoughtfully and remain relevant over time.

## Defined Terms

>---
>**TODO**
>
>- Add any important defined terms as outlined in [ATOMONE CONSTITUTION](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md#defined-terms)
>- Defined terms should be added to remove ambiguity
>
>---

TBD

## TODOs

- [ ] GovDAO requirements (not specifically for v1 implementation, but generally for any implementation)  
- [ ] Airdrop distribution?
- [ ] ...