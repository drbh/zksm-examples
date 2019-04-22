package main

import (
	"crypto/rand"
	"fmt"
	"github.com/ing-bank/zkproofs/go-ethereum/crypto/bn256"
	// "github.com/drbh/zkproofs/go-ethereum/crypto/bn256"
	// "github.com/ing-bank/zkproofs/go-ethereum/zkproofs"
	"github.com/drbh/zkproofs/go-ethereum/zkproofs"
	"math/big"
	"time"
)

func main() {
	var (
		r *big.Int
		s []int64
	)
	s = make([]int64, 9)
	s[0] = 12
	s[1] = 42
	s[2] = 61
	s[3] = 71

	s[4] = 1230
	s[5] = 121
	s[6] = 99421
	s[7] = 4210
	s[8] = 7111

	startTime := time.Now()
	p, _ := zkproofs.SetupSet(s)
	setupTime := time.Now()
	fmt.Println("Setup time:")
	fmt.Println(setupTime.Sub(startTime))

	r, _ = rand.Int(rand.Reader, bn256.Order)
	proof_out, _ := zkproofs.ProveSet(12, r, p)
	proofTime := time.Now()
	fmt.Println("Proof time:")
	fmt.Println(proofTime.Sub(setupTime))

	result, _ := zkproofs.VerifySet(&proof_out, &p)
	verifyTime := time.Now()
	fmt.Println("Verify time:")
	fmt.Println(verifyTime.Sub(proofTime))

	fmt.Println("ZK Set Membership result: ")
	fmt.Println(result)

	fmt.Println("")
}

func main() {

	fmt.Println("\n\nBASIC ZKSM\n\n\n")
	runDemo()

}
