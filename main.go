package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"time"

	zzz "github.com/arnaucube/kzg-commitments-study"
)

type Output struct {
	Z          string `json:"z"`
	Y          string `json:"y"`
	Polynomial string `json:"polynomial"`
	Commitment string `json:"commitment"`
	Proof      string `json:"proof"`
	Verified   bool   `json:"verified"`
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func evaluatePolynomial(p []*big.Int, x int) int {
	result := 0
	for i, coeff := range p {
		result += int(coeff.Int64()) * powInt(x, i)
	}
	return result
}

func generateRandomPolynomial(degree int, seed int64) []*big.Int {
	rng := rand.New(rand.NewSource(seed))
	p := make([]*big.Int, degree+1)
	for i := range p {
		p[i] = big.NewInt(rng.Int63n(10)) // random coeffs 0–9
	}
	return p
}

func runProof(ts *zzz.TrustedSetup, p []*big.Int, z *big.Int, y *big.Int) (string, bool) {
	proof, err := zzz.EvaluationProof(ts, p, z, y)
	if err != nil {
		log.Fatalf("Failed to generate proof: %v", err)
	}
	c := zzz.Commit(ts, p)
	ok := zzz.Verify(ts, c, proof, z, y)
	return proof.String(), ok
}

func main() {
	// === Flags ===
	zFlag := flag.Int("z", 3, "Evaluation point z")
	degreeFlag := flag.Int("deg", 3, "Degree of the polynomial")
	seedFlag := flag.Int64("seed", time.Now().Unix(), "Random seed for polynomial generation")
	jsonFlag := flag.Bool("json", false, "Output as JSON")

	flag.Parse()

	// === Polynomial setup ===
	p := generateRandomPolynomial(*degreeFlag, *seedFlag)

	// === Evaluate at z ===
	z := big.NewInt(int64(*zFlag))

	yVal := evaluatePolynomial(p, *zFlag)
	y := big.NewInt(int64(yVal))

	// === Trusted setup ===
	ts, err := zzz.NewTrustedSetup(len(p))
	if err != nil {
		log.Fatalf("Failed to create trusted setup: %v", err)
	}

	// === Run proof and verification ===
	proofStr, verified := runProof(ts, p, z, y)

	// === Output ===
	out := Output{
		Z:          z.String(),
		Y:          y.String(),
		Polynomial: zzz.PolynomialToString(p),
		Commitment: zzz.Commit(ts, p).String(),
		Proof:      proofStr,
		Verified:   verified,
	}

	if *jsonFlag {
		prettyJSON, _ := json.MarshalIndent(out, "", "  ")
		fmt.Println(string(prettyJSON))
	} else {
		fmt.Println("==== KZG Polynomial Commitment Proof ====")
		fmt.Printf("Polynomial: %s\n", out.Polynomial)
		fmt.Printf("Evaluation Point z: %s\n", out.Z)
		fmt.Printf("Expected y = p(z): %s\n", out.Y)
		fmt.Printf("Commitment: %s\n", out.Commitment)
		fmt.Printf("Proof: %s\n", out.Proof)
		fmt.Printf("Verification Passed? %v\n", out.Verified)
	}
}

/*
go mod init trial_kzg
go mod tidy
go run main.go

go run main.go -z 4 -deg 5 -seed 123 -json

🔐 For Real Randomness (e.g., from Verifier)
In an interactive version of the protocol:
- The verifier picks a fresh random z each time.
- This makes the challenge truly random, not just unpredictable.

But that brings in:
- The need for communication (round-trips),
- State management (to prevent reuse),
- A way to verify z was chosen after the commitment.

🛡️ Hybrid Approach: Domain Separation to Make z Unique per Session
For per-session randomness, you can combine the commitment with a random nonce or session ID, like so:

```
sessionID := []byte("session-abc-123") // can be a timestamp, UUID, or verifier random
input := append(sessionID, commitmentBytes...) // domain-separated Fiat-Shamir
hash := sha256.Sum256(input)
z := new(big.Int).SetBytes(hash[:])
```

``` deterministic Fiat-Shamir within a session:
// different sessions lead to different z values — a form of randomness without relying on the prover:
sessionFlag := flag.String("session", "default", "Session/domain separation tag")
...
input := append([]byte(*sessionFlag), commitmentBytes...)
hash := sha256.Sum256(input)
z := new(big.Int).SetBytes(hash[:])
````
`
This:
- Still uses Fiat-Shamir,
- Still avoids interaction,
- Makes z session-unique, solving your randomness concern.

You could even expose a -session flag in the CLI.


trial_kzg % go run main.go -z 4 -deg 7
==== KZG Polynomial Commitment Proof ====
Polynomial: 8x⁷ + 2x⁶ + 3x⁵ + 9x⁴ + 7x³ + 5x² + 5x¹ + 5
Evaluation Point z: 4
Expected y = p(z): 145193
Commitment: bn256.G1(178d052e4fc294f4da4d961ec274c95ce3146f81e282a96dc820b5fafbe989f9, 2d7f0eab5893eab4321824b9367563473db9757c7250e40f0389d2168428d178)
Proof: bn256.G1(16cd1bb71cc959c38e643fa8cf59c872a8a76cca20573d1442292153dc7a0eaf, 0f966974d7a0a1fd057da10291fe5eff4c811e38483aeb6aafecdd7cc7e3dd2d)
Verification Passed? true
*/
