# KZG polynomial commitment scheme 

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

$$\ xG ∈ 𝔾_1 $$\ defines a point (x.G ) on 1 curve ($$\ 𝔾_1 \$$) and $$\ xH ∈ 𝔾_2 \$$ defines a point (x.H ) on the other curve ($$\ 𝔾_2 \$$) . Notation used is: 

$$\ [x]_1 = xG ∈ 𝔾_1 \$$
$$\ [x]_2 = xH ∈ 𝔾_2 \$$

and $$\ 𝔾_1 = ⟨G⟩  \$$ and $$\ 𝔾_2 = ⟨H⟩ \$$. 

G is the generator of $$\ 𝔾_1 \$$, and H is the generator of $$\ 𝔾_2 \$$ .

## Trusted Setup 

In the first part, use a Trusted Setup, and should be computed with a Multi-Party Computation (MPC). 

The parameters are generated from random $$\ τ ∈ 𝔽_p \$$, and from this parameter we can compute $$\  [τ^i]_1 \$$ and $$\ [τ^i]_2 \$$ for i=0,...,n−1 : 

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

With this we have a polynomial of $$\ p(x)=∑^n _{i=0} p_i x^i \$$, 

and then can create a commitment with: 
$$\ c=[p(τ)]_1 \$$\ 

where $$\ c = ∑^{deg(p(x))} _{i=0} [τ^i]⋅p_i $$\ 

Prover would send this commitment c to Verifier as a proof against the polynomial. 
Verifier would choose a value $$\ z ∈ 𝔽_p \$$, 
where $$\ 𝔽_p \$$ := finite field defined by the polynomial. 

## Evaluating the proof 

Prover computes p(z)=y , and a quotient polynomial is derived as: 
$$\ q(x) = \frac{p(x)−y}{x−z} \$$. 

This polynomial proves that p(z)=y , and where p(x)−y is divisible by x−z . 
This means it has a root at z - as p(z)−y=0 . 

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

Note that we have q(x)(x−z)=p(x)−y , and which can be rearranged as $$\ q(x)=\frac{p(x)−y}{x−z} \$$. 
This is evaluated at τ for the trusted setup and not known for $$\ q(τ) = \frac{p(τ)−y}{τ−z} \$$.


"""

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
