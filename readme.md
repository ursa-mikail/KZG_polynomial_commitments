# KZG (Kate-Zaverucha-Goldberg) Polynomial Commitment Scheme

KZG polynomial commitment scheme allows a committer (Prover) to commit to a polynomial with a short string. Verifier can send Prover a challenge, and of which she creates a proof against the committed polynomial. It has applications into Credentials and Selective Disclosure of Signed Data. ZKS (zero-knowledge sets), and Veriable Secret Sharing (VSS). 
We generate an order of polynomial, and send a random Evaluation Point z: 

```
go run main.go -z 300 -deg 7
==== KZG Polynomial Commitment Proof ====
Polynomial: 6x⁷ + 5x⁶ + x⁴ + 4x³ + 9x² + 8x¹ + 0

Evaluation Point z: 300
Expected y = p(z): 1315845008208812400
Commitment: bn256.G1(25dd18a1a425edf06634351689bcfe83ce772b87b8d329d660f4205b8fee3661, 059405e251b41a316cc1446bc121301ee0e8e14157ea376e3343d5a315f1e87c)
Proof: bn256.G1(109125106987faacf54260cb97d4e1d7d29be021e6ca5cea9db1f653cec39128, 2cbde997eb486e6eb28c78cee79562f93e2afd6e3073b70e76db5d713094279a)
Verification Passed? true
```

$$\ xG ∈ 𝔾_1 \$$ defines a point (x.G ) on 1 curve ($$\ 𝔾_1 \$$) and $$\ xH ∈ 𝔾_2 \$$ defines a point (x.H ) on the other curve ($$\ 𝔾_2 \$$) . Notation used is: 

$$\ [x]_1 = xG ∈ 𝔾_1 \$$
$$\ [x]_2 = xH ∈ 𝔾_2 \$$

and $$\ 𝔾_1 = ⟨G⟩  \$$ and $$\ 𝔾_2 = ⟨H⟩ \$$. 

G is the generator of $$\ 𝔾_1 \$$, and H is the generator of $$\ 𝔾_2 \$$ .

## Trusted Setup 

In the first part, use a Trusted Setup, and should be computed with a Multi-Party Computation (MPC). 

The parameters are generated from random $$\ τ ∈ \mathbb{F} _p \$$, and from this parameter we can compute $$\  [τ^i]_1 \$$ and $$\ [τ^i]_2 \$$ for i=0,...,n−1 : 

$$\
[τ^i]_1 = ([τ^0]_1, [τ^1]_1, [τ^2]_1,...,[τ^{n−1}]_1) 
\$$

$$\
[τ^i]_2 = ([τ^0]_2, [τ^1]_2, [τ^2]_2,...,[τ^{n−1}]_2) 
\$$

Which in additive representation is: 

$$\
(G,τG,τ^2 G,...,τ^{n−1} G) ∈ 𝔾1 
\$$

$$\
(H,τH,τ^2 H,...,τ^{n−1} H) ∈ 𝔾2 
\$$

In this stage, we have taken a random value (τ ) and produced a tuple of size d+1 and where d is the polynomial degree of our target polynomial. We end up with $$\ {G,τG,τ^2 G,τ^3 G,…,τ^d G} \$$. After generating this, the random value of **τ should be deleted** .

## Commitment 

With this, we have a polynomial of $$\ p(x)=∑^n _{i=0} p_i x^i \$$, 

and we can create a commitment with: 
$$\ c=[p(τ)]_1 \$$

where $$\ c = ∑^{deg(p(x))} _{i=0} [τ^i]⋅p_i \$$

Prover would send this commitment c to Verifier as a proof against the polynomial. 
Verifier would choose a value $$\ z ∈ \mathbb{F} _p \$$, 
where $$\ \mathbb{F} _p \$$ := finite field defined by the polynomial. 

## Evaluating the proof 

Prover computes p(z)=y , and a quotient polynomial is derived as: 
$$\ q(x) = \frac{p(x)−y}{x−z} \$$. 

This polynomial proves that p(z)=y , and where p(x)−y is divisible by x−z . 
This means it has a root at z, as p(z)−y = 0 . 

The evaluation proof is: $$\ π = [q(τ)]_1 \$$ and which is determine from $$\ π = ∑^{deg(q(x))} _{i=0} [τ^i]⋅q_i \$$

Prover can send this evaluation proof π to Verifier. 

## Verifying the proof 

Verifier has the commitment of $$\ c=[p(τ)]_1 \$$ , the evaluation of y=p(z) , and the proof of $$\ π=[q(τ)]_1 \$$ . He can now check the pairing of: 

$$\ ê (π,[τ]_2−[z]_2)==ê (c−[y]_1,H) \$$

Prover provides π and c are given by prover before, and where $$\ [τ]_2 \$$ is derived in the trusted setup, $$\ [z]_2 \$$ defines the point at which the polynomial has been evaluated with, and $$\ [y]_1 \$$ defines the claimed value of p(z) . 

This works because: 

$$\ ê (π,[τ]_2−[z]_2)==ê (c−[y]_1,H) \$$

$$\ ⇒ ê ([q(τ)]_1,[τ−z]_2)==ê ([p(τ)]_1−[y]_1,H) \$$

$$\ ⇒[q(τ)⋅(τ−z)]_T==[p(τ)−y]_T \$$

Note: we have q(x)(x−z)=p(x)−y , and which can be rearranged as $$\ q(x)=\frac{p(x)−y}{x−z} \$$. 
This is evaluated at τ for the trusted setup and not known for $$\ q(τ) = \frac{p(τ)−y}{τ−z} \$$.


```
make init         # initialize module and tidy
make build        # compile the project
make run          # run with default args
make run-debug    # run with debug args (z=3, deg=4, seed=123)
make clean        # remove the binary
make fmt          # format code
make vet          # run static analysis


✅ Usage Examples
Run locally
make init
make build
make run-debug

Run tests
make test

Docker
make docker-build
make docker-run-debug

make docker-clean

🛑 Warning
Running `docker system prune -a --volumes`:
- Deletes all images not associated with a running container.
- Deletes all stopped containers.
- Deletes all unused volumes.
- Deletes all unused networks.

Don't use this on a production system or without backups.
```

✅ Key Features:
- Constant-size commitment
- Constant-size proof
- Efficient verification
- Trusted setup required (a structured reference string)


KZG is used to:
- Commit to model weights
- Prove correct evaluation of forward passes
- Ensure integrity of cryptographic hash chains over model operations


| AI/ML Scenario       | How KZG Helps                          | Benefit                            |
| -------------------- | -------------------------------------- | ---------------------------------- |
| Verifiable Inference | Commit to model + prove correct output | Offload compute securely           |
| Federated Learning   | Commit to private updates              | Privacy + auditability             |
| ZKML                 | Commit to models in ZK circuits        | Private + trustless inference      |
| Data Integrity       | Commit to training data                | Prevent tampering                  |
| On-chain AI          | Commit to models/inference proofs      | Efficient, trust-minimized compute |



## 🔹 Use Case: Data Availability Sampling in Ethereum (EIP-4844 / Proto-Danksharding)
🧩 Problem:
In Ethereum scaling (e.g., rollups), large blobs of data (like execution traces or rollup state transitions) need to be published to L1 but don’t need to be read by every node — only verifiably available.

But how can Ethereum verify large data blobs are available without every node downloading all of it?

## 🔹 Use Case: Data Availability Sampling in AI/ML
🧩 Problem:
Large blobs of data (like execution traces or data updates or builds updates state transitions) need to be published but don’t need to be read by every package and data blob — only verifiably available.

But how can AI/ML verify large data blobs and build packages are available without every computational node downloading all of it?

🔐 Why It’s Good:
- Efficient: Fast verification, constant size.
- Succinct: Short commitments and proofs.
- Scalable: Ideal for L2 scaling and ZK-rollups.
- Trusted Setup: Needs one-time setup, but reused across many applications (e.g., Ethereum’s KZG CRS).


### ✅ Solution with KZG Commitments:
1. Polynomial Representation:
- The data blob (e.g., a 4096-byte chunk) is interpreted as evaluations of a polynomial 𝑓 ( 𝑥 ) f(x) over a finite field at various points. 
- So the blob becomes values like 𝑓 ( 1 ) , 𝑓 ( 2 ) , . . . , 𝑓 ( 𝑛 ).

2. Commitment:
- A single KZG commitment to the polynomial is published on-chain.
- This commitment is a short cryptographic object (constant size, e.g., 48 bytes with BLS12-381 curve).

3. Sampling: 
- Light clients or full nodes randomly sample a few indices $$\ 𝑥_𝑖 ​ \$$ from the blob. 
- The blob producer (e.g., a rollup operator) then sends a KZG proof that $$\ 𝑓 ( 𝑥_𝑖 ) = 𝑦_𝑖 ​ \$$ for each requested $$\ 𝑥_𝑖 ​ \$$.

4. Verification:
- The KZG proof is constant size and can be verified quickly using pairing operations.
- If enough sampled points verify correctly, the blob is assumed to be available with high probability.


📦 Other Use Cases:
1. Zero-Knowledge Proofs (e.g., PLONK, zk-SNARKs):
KZG is used to commit to witness polynomials and prove that computations were done correctly.

2. Verifiable Computation:
You can outsource computation over polynomials and verify results efficiently.

3. Verifiable Secret Sharing (VSS):
Share a secret via Shamir’s scheme, commit to the polynomial, and let receivers verify their shares.

## ✅ AI/ML Use Cases for KZG Polynomial Commitments
1. Verifiable Machine Learning (VML)
Use case: Ensuring the correctness of ML inference or training outsourced to an untrusted party.

Example:
A model owner outsources inference (or part of training) to a third party (e.g., edge devices, cloud compute).

The model is expressed via polynomials (common in ZKML pipelines — e.g., ReLU approximated with polynomials).

The prover commits to the polynomial representation of the model and inputs using KZG.

They return:
- The inference result
- KZG proofs showing correct evaluation at specific points

✅ Benefit: A verifier (e.g., a client or auditor) checks that the model output was computed faithfully, without needing to run the full model themselves.

2. Privacy-Preserving Federated Learning
Use case: Multiple parties collaboratively train a model, but want privacy and integrity guarantees.

Scenario:
Each participant shares model updates.

Instead of sharing raw data or even full gradient vectors, participants:
- Encode their updates as polynomials
- Commit to them using KZG
- Optionally prove integrity of the data (e.g., constraints were followed)

✅ Benefit: The aggregator can verify that participants are behaving honestly without seeing the private data.


3. Zero-Knowledge ML Inference
Use case: Prove that a model made a decision correctly without revealing the model or the input.

Example:
A loan approval system runs an ML model.

The system wants to prove that a user was denied based on fair and agreed-upon logic, without leaking sensitive financial data.

Using ZK-SNARKs + KZG:
- Encode the model as polynomials (e.g., arithmetic circuits, quantized networks).
- Commit to model and input with KZG.
- Publish constant-size proof of correct inference.

✅ Benefit: Enables auditable AI with privacy.


4. Data Provenance & Integrity
Use case: AI models are trained on datasets that must be verified for integrity (e.g., medical data, sensitive research).

Each data provider encodes their dataset into a polynomial.

KZG commitment ensures:
- That the data has not been tampered with.
- That any claimed data was actually used.

✅ Benefit: Tamper-proof audits of data used in ML models.

5. On-Chain or Verifiable AI Inference (Web3 + AI)
Use case: Models served on decentralized platforms or L2 chains need trustless inference guarantees.

KZG enables committing to the model weights (or even quantized model architectures).

Inference done off-chain returns a proof using KZG + zk-SNARK that the model evaluated properly.

Only commitment and verification on-chain.

✅ Benefit: Verifiable off-chain compute with low cost on-chain commitments.




| Feature           | KZG Commitment                                |
| ----------------- | --------------------------------------------- |
| Commitment Size   | Constant                                      |
| Proof Size        | Constant                                      |
| Verification Time | Fast (pairing)                                |
| Trusted Setup     | Yes                                           |
| Used in           | Ethereum (EIP-4844), PLONK, zkEVM, ZK-Rollups |


