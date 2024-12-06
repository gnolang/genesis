# [DRAFT] gno.land Constitution


>---
>**DOCUMENT NOTES**
>- This is a v0 working document. It is very early and rough around the edges. The goal here is to introduce an early working prototype that can live in [gnolang/genesis](https://github.com/gnolang/genesis) as a v0 document.
>- This prototype takes heavy influence from the [ATOMONE CONSTITUTION](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md).
>- Public discussions on this topic can be found int the github issue [gno.land Constitution #3080](https://github.com/gnolang/gno/issues/3080)
> ### Goals
>- citizens, gnomes, should feel protected by this constitution
>- Constitution is about protecting long-term value of the chain
>- Quality over quantity
>- we are creating a DAO to protect the chain and keep it running for the very long term. Building a team anyone can trust and rely on.
>- Ensure that the skeleton is in place and placeholders are clear
>- Its ok to remove things as needed for v1.

>---

## Preamble

>TODO: hone this list for this preamble to ensure it meets a succinct set of goals for the constitution.

We, the people of gno.land, in order to create a more decentralized world, open internet, transparent governance, and citizen-powered innovation, do ordain and establish this Constitution for gno.land.

### Declaration of Intent

>TODO: Outline the intent of this, we should be able to utilize prior work from those who contributed to Why Gno and other blog posts.

The vision of gno.land is...

## Article 1: General Provisions

### Section 1: Fundamental Principles

This Constitution of gno.land, hereinafter “Constitution”, hereby establishes the foundations of the governance model, economical model, and operating system of gno.land.

All subsequent governance proposals shall align with the provisions of this Constitution, and proponents of each proposal, along with all active governance voters, are required to ensure consistency between such proposals and this Constitution, and cannot violate any explicit restriction.

Amendments are permitted to innovate and adapt this Constitution, but they must still respect and adhere to the fundamental principles upon which this Constitution is founded, and are subject to specific procedures as described in Governance Article 2.

### Section 2: Sovereignty of gno.land

> TODO: TBD

### Section 3: General Mission and Objectives

> TODO: Add mission and objectives

The mission of gno.land is to establish a platform focused on transparent governance, user agency and authentic content.

The objectives of gno.land are:

- Develop a governance model based on contributions
- Encourage forking as a tool for decentralization and innovation
- Reward authentic content and content moderation
- Build spam and manipulation resilience

### Section 4: Rights, Liberties and Obligations in gno.land

> TODO: Need to outline the rights of citizens (gnomes). This should outline the citizens rights, how one might become a citizen, and the obligations one has as a citizen. 

Citizen rights
- The freedom to fork and innovate
- The freedom to participate in transparent governance
- The right to access the decentralized, censorship resistant platform

Citizen obligations
- To contribute to the community development and moderation
- To uphold transparency and integrity

## Article 2: Governance

### Section 1: The gno.land Chain

>---
>**NOTES**
>- This should cover the unique characteristics gno.land such as how Proof of Contribution, Validators, and DAOs work related to this.
>---

The gno.land chain is the foundational infrastructure enabling decentralized applications and transparent governance. It establishes the framework for Proof of Contribution, validator selection, and DAOs. This section defines the unique characteristics of the gno.land chain and its core principles:  

Proof of Contribution

- Governance is driven by meaningful contributions rather than financial influence.
- Contributions include creating content, moderating platforms, developing software, and maintaining the chain's security. (TO BE CONFIRMED)

Validator Selection

- Validators are chosen based on technical capability and alignment with gno.land’s vision, not purely on capital stake.
- Validators must uphold the principles of decentralization, transparency, and security.

DAOs and Sub-DAOs

- The chain supports GovDAO, the principal governing DAO, and any number of DAOs & sub-DAOs for governance and operations.
- Each DAO operates under the principles of transparency and collective decision-making.


### Section 2: Laws and Amendments

>---
>**TODO**
>- Largely carried over from AtomOne
>- (n2p5): I replaced "AtomOne Governance Hub" with "gno.land Governance", but there needs to be a review on proper structure
>- Goal: explain that GovDAO should have things like super majority for decisions.
>- Goal: OK to request schism with 40% of vote. (maybe TBD)
>- Goal: chain upgrade
>- Goal: parameter tuning
>- Goal: List Cases and set of requirements. 
>---

This Constitution allows the creation and enactment of laws as means to complement and enhance its content, restrictions, provisions and established processes.

Laws shall be enacted to address specific governance, economic, and operational aspects, and must align with the principles and framework established by this Constitution. In cases where laws conflict with this Constitution, this Constitution shall always prevail. The gno.land Governance structure will include several Decentralized Autonomous Organizations (DAOs) and sub-DAOs. Each DAO may create its own bylaws to govern its operations, provided they do not conflict with this Constitution, overarching laws or any mandates issued by the gno.land Governance. Mandates issued by an ancestor of a DAO are binding and cannot be amended or repealed by those DAOs or their sub-DAOs and remain in effect unless modified or repealed. Any DAO bylaws that conflict with its Governing Documents shall be deemed invalid and abrogated.

Laws may be proposed by any Citizen of gno.land through the established governance process. The approval of new laws, amendment of existing laws or amendment of this Constitution require a Constitutional Majority vote. However, an endorsement from the Steering DAO shall lower the passing threshold to a Supermajority vote for all proposals except for Constitutional amendment proposals, which shall require a Constitutional Majority.

Once an amendment is approved by the required majority, it is implemented according to the specified procedures. The implementation process includes updating this Constitution, communicating the changes to the community, and ensuring compliance with the new provisions.

Laws must be named plaintext files or folders of plaintext files, or folders of folders.

### Section 3: gno.land Governance

>---
>**TODO**
>- This is modified from AtomOne, but needs to be checked for accuracy.
>- There may be governance items currently missing.
>- Outline relationship in GovDAOs (worksDAO, GovDAO, OversightDAO, ConstitutionDAO, etc)
>---

The working language of gno.lang Governance is International English.

The governance process shall be defined through law, but will strictly adhere to all restrictions and provisions of this Constitution.

The vote types are YES, NO, and ABSTAIN.

To prevent spam and to ensure the quality of proposals, proposals must have a sufficient amount of burn deposit before voting can begin. The minimum amount of burn deposit needed shall self-adjust to target on average 1 proposal per 2 week period.

The quorum necessary for a proposal to be valid is 40%. The denominator shall be the number of bonded GNOT tokens.

A Simple Majority in gno.land Governance is defined to be exactly “more than half”. No proposal is considered to have passed in gno.lad with a Simple Majority; a Supermajority or Constitutional Majority is required. A Supermajority in gno.land Governance is defined to be exactly "more than two thirds". This is distinct from a Supermajority in DAO council governance which is defined to be exactly “two thirds or more”. For clarity, 3 out of 4 votes is necessary for a Supermajority in gno.land Governance, but 2 out of 3 is sufficient for a Supermajority in DAO council voting. The Constitutional Majority threshold is defined to be exactly “more than 90%”. These definitions of majority cannot change even by a Constitutional Majority.

The minimum voting period for a proposal is 3 weeks. To ensure adequate representation and participation, the governance process must extend the voting deadline to guarantee a minimum of 2 weeks of voting after the minimum quorum has been met.

### Section 4.a: Common DAO Spec

>---
>- TODO: Review. This is largely ported over from ATOMONE
>---

These common DAO specifications shall apply for all Core DAOs unless otherwise specified, Special Purpose DAOs, and all sub-DAOs of these DAOs. Other DAOs that are not Core DAOs or Special Purpose DAOs or Descendants of these DAOs need not implement these specifications.

All sub-DAOs have parent DAOs. The parent DAO of the Core DAOs and Special Purpose DAOs are the governance of gno.land itself. Therefore all Core DAOs and Special Purpose DAOs as well as their sub-DAOs and gno.land Governance itself altogether form a tree structure. The parent DAO and the parent DAO’s parent DAO and so on, all the way up to Hub Governance are altogether called the Ancestors of a DAO. The sub-DAOs and their sub-DAOs and so on are called the Descendants of a DAO.

Every DAO, upon creation, must have a Charter (which is composed of Purpose and Description), an initial set of Council members (which may be empty) and may also have Bylaws and Mandates, The Purpose and Description must be plaintext files. The Bylaws and Mandates must be named plaintext files or folders of plaintext files, or folders of folders.

A DAO’s Charter, Bylaws, and Mandates may be changed by a Simple Majority vote from any of the DAO’s ancestors, except from AtomOne Hub Governance which shall require a Supermajority vote.

A DAO’s Bylaws, Mandates, and the Bylaws and Mandates of its ancestor DAOs, the relevant Laws, and this Constitution, altogether are called the Governing Documents of the DAO.

A DAO has a Council composed of zero or more members, with no maximum number of members unless otherwise specified in its Governing Documents.

The Council of a DAO may change the Bylaws of the DAO, and otherwise make Decisions on behalf of the DAO by passing Proposals.

A DAO may establish any number of sub-DAOs through the DAO Council’s Simple Majority vote, with their own defined Charters and specific bylaws and mandates, as necessity may arise and in accordance with the parent DAO Charter and bylaws. Sub-DAOs are owned by and can be controlled by the parent DAO, and members are also subject to the ancestor DAOs’ bylaws and mandates.

A Simple Majority in DAO governance is defined to be exactly “more than half”. A Supermajority in DAO governance is defined to be exactly "two thirds or more". This is distinct from a Supermajority in gno.land Governance.

By default, unless specified otherwise in its Governing Documents, the following rules shall apply for Council voting:

- Each member shall have equal voting power (no member may occupy multiple seats).
- A Council member may resign and thereby remove themselves from the Council.
- The tally denominator is the number of voters minus ABSTAINs (no quorum requirement).
- Proposals are open until they are decided by sufficient majority, or dismissed, or expired.
- proposals are immediately dismissed by a Simple Majority vote of NO.

By default, unless specified otherwise in its Governing Documents, the following rules shall apply for Council membership election:

- The Council may elect one or more new members, and/or remove one or more members, by Super Majority vote. (self mutating).
- The DAO’s Ancestors may modify the Council membership with a Super Majority vote.
- Each DAO shall have an associated crypto address which can hold any number of tokens.

DAOs may operate with logic on core shards, or, represented as a m-of-n multisig account on the AtomOne hub where the signers are each members of the DAO’s council, where m is more than ½ n and also m is 3 or more. In all cases financial transactions from the DAO’s treasury must follow the passage of governance proposals on the DAO.

### Section 4.b: Core DAOs with Special Powers

>---
>**TODO**
>- We need feedback on these sections and gno.land vision on initial DAO structure.
>- [AtomOne uses this to  cover Steering DAO, Oversight DAO, etc](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md#section-4b-core-daos-with-special-powers)
>---

TBD

### Section 5: Airdrops and forks

>---
>**TODO**
>- How much of this [should we port over from ATOMONE](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md#section-5-airdrops-and-forks)?
>---

TBD

### Section 6: Special Purpose DAOs

>---
>**TODO**
>- How much of this should be [ported over from ATOMONE](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md#section-6-special-purpose-daos)?
>---

TBD

## Article 3: Economics

### Section 1: Economic Model

>---
>- TODO: TBD
>---


### Section 2: The GNOT Token

>---
>- TODO: TBD
>---


### Section 3: Inflation

>---
>- TODO: TBD
>---

### Section 4: The Community Pool

>---
>- TODO: TBD
>---

### Section 5: Validators

>---
>- TODO: TBD
>---

## Article 4: Implementation

### Section 1: Updates to gno.land

>---
>**TODO**
>- largely [ported from ATOMONE](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md#section-6-special-purpose-daos)
>- Review for correctness and add gno.land specific rules.
>---

All updates to gno.land must be proposed as distinct, independent components. Each component must be discussed and voted on separately to ensure thorough deliberation and transparency. Proposals shall not bundle different changes into a single submission. Each proposal must address a single, clearly defined update or change. The only exception is made to regular code maintenance and bug fixes, and critical security updates.

Adequate time must be provided between the consideration of each proposal to allow for comprehensive review and community input. The minimum duration for discussion and voting on each proposal shall be two weeks.

The passing of Software Update Proposals shall require the passage of law.

[TODO: name stakeholders from governance structure] are responsible and may be held accountable for voting YES only on software update proposals that make correct changes toward the living constitution and laws (any substantial changes to the software must first be reflected in the constitution and laws). This responsibility exists independently of the Oversight DAOs’ regulatory function.

### Section 2: Protocols

>---
>- TODO: TBD
>---

### Section 3: Compute/storage/memory limitations

>---
>- TODO: TBD
>---

## Defined Terms

>---
>**TODO**
>- Add any important defined terms as outlined in [ATOMONE CONSTITUTION](https://github.com/atomone-hub/genesis/blob/main/CONSTITUTION.md#defined-terms)
>- Defined terms should be added to remove ambiguity
>---

TBD

## TODOs

- [ ] GovDAO requirements (not specifically for v1 implementation, but generally for any implementation)  
- [ ] Airdrop distribution?
- [ ] ...