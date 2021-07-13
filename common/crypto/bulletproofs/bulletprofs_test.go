//+build bulletproofs

/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bulletproofs

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBulletproofs(t *testing.T) {
	bulletproofs := Helper().NewBulletproofs()
	commitment, opening, err := bulletproofs.PedersenCommitRandomOpening(10)
	if err != nil {
		panic(err)
	}
	commitment2, err := bulletproofs.PedersenCommitSpecificOpening(10, opening)
	if err != nil {
		panic(err)
	}
	commitment3, opening3, err := bulletproofs.PedersenCommitRandomOpening(100)
	if err != nil {
		panic(err)
	}
	commitment4, err := bulletproofs.PedersenAddNum(commitment, 5)
	if err != nil {
		panic(err)
	}
	commitment5, opening5, err := bulletproofs.PedersenAddCommitmentWithOpening(commitment, commitment3, opening, opening3)
	if err != nil {
		panic(err)
	}
	commitment6, err := bulletproofs.PedersenSubNum(commitment3, 20)
	if err != nil {
		panic(err)
	}
	commitment7, opening7, err := bulletproofs.PedersenSubCommitmentWithOpening(commitment3, commitment, opening3, opening)
	if err != nil {
		panic(err)
	}
	commitment8, opening8, err := bulletproofs.PedersenSubCommitmentWithOpening(commitment3, commitment6, opening3, opening3)
	if err != nil {
		panic(err)
	}
	commitment9, opening9, err := bulletproofs.PedersenSubCommitmentWithOpening(commitment3, commitment3, opening3, opening3)
	if err != nil {
		panic(err)
	}

	commitmentMul, openingMul, err := bulletproofs.PedersenMulNumWithOpening(commitment3, opening3, 20)
	if err != nil {
		panic(err)
	}

	ret, err := bulletproofs.PedersenVerify(commitment, opening, 10)
	fmt.Println("1: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitment2, opening, 10)
	fmt.Println("2: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitment3, opening3, 100)
	fmt.Println("3: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitment, opening, 100)
	fmt.Println("4: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitment3, opening, 10)
	fmt.Println("5: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitment4, opening, 15)
	fmt.Println("6: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitment5, opening5, 110)
	fmt.Println("7: ", ret)

	ret, err = bulletproofs.PedersenVerify(commitment6, opening3, 80)
	fmt.Println("sub1: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitment7, opening7, 90)
	fmt.Println("sub2: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitment8, opening8, 20)
	fmt.Println("sub3: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitment9, opening9, 0)
	fmt.Println("sub4: ", ret)

	ret, err = bulletproofs.PedersenVerify(commitmentMul, openingMul, 2000)
	fmt.Println("Mul: ", ret)

	proof, commitmentf, openingf, err := bulletproofs.ProveRandomOpening(10)
	if err != nil {
		panic(err)
	}
	proof2, commitmentf2, err := bulletproofs.ProveSpecificOpening(10, openingf)
	if err != nil {
		panic(err)
	}
	proof3, commitmentf3, _, err := bulletproofs.ProveRandomOpening(100)
	if err != nil {
		panic(err)
	}

	proof1Base64 := base64.StdEncoding.EncodeToString(proof)
	proof2Base64 := base64.StdEncoding.EncodeToString(proof2)

	fmt.Println("proof1: " + proof1Base64)
	fmt.Println("proof2: " + proof2Base64)

	ret, err = bulletproofs.Verify(proof, commitmentf)
	fmt.Println("1: ", ret)
	ret, err = bulletproofs.Verify(proof2, commitmentf2)
	fmt.Println("2: ", ret)
	ret, err = bulletproofs.Verify(proof2, commitmentf)
	fmt.Println("3: ", ret)
	ret, err = bulletproofs.Verify(proof, commitmentf2)
	fmt.Println("4: ", ret)
	ret, err = bulletproofs.Verify(proof3, commitmentf3)
	fmt.Println("5: ", ret)
	ret, err = bulletproofs.Verify(proof3, commitmentf)
	fmt.Println("6: ", ret)

	proofAdd, commitmentAdd, err := bulletproofs.ProveAfterAddNum(10, 30, opening, commitment)
	if err != nil {
		panic(err)
	}
	proofAddC, commitmentAddC, openingAddC, err := bulletproofs.ProveAfterAddCommitment(100, 10, opening3, opening, commitment3, commitment)
	if err != nil {
		panic(err)
	}

	ret, err = bulletproofs.Verify(proofAdd, commitmentAdd)
	fmt.Println("Add num proof: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitmentAdd, opening, 40)
	fmt.Println("Add num commit: ", ret)
	ret, err = bulletproofs.Verify(proofAddC, commitmentAddC)
	fmt.Println("Add commit proof: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitmentAddC, openingAddC, 110)
	fmt.Println("Add commit commit: ", ret)

	proofSub, commitmentSub, err := bulletproofs.ProveAfterSubNum(100, 10, opening3, commitment3)
	if err != nil {
		panic(err)
	}
	proofSubC, commitmentSubC, openingSubC, err := bulletproofs.ProveAfterSubCommitment(100, 10, opening3, opening, commitment3, commitment)
	if err != nil {
		panic(err)
	}

	ret, err = bulletproofs.Verify(proofSub, commitmentSub)
	fmt.Println("Add num proof: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitmentSub, opening3, 90)
	fmt.Println("Add num commit: ", ret)
	ret, err = bulletproofs.Verify(proofSubC, commitmentSubC)
	fmt.Println("Add commit proof: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitmentSubC, openingSubC, 90)
	fmt.Println("Add commit commit: ", ret)

	proofMult, commitmentMult, openingMult, err := bulletproofs.ProveAfterMulNum(100, 10, opening3, commitment3)
	if err != nil {
		panic(err)
	}

	ret, err = bulletproofs.Verify(proofMult, commitmentMult)
	fmt.Println("Mul num proof: ", ret)
	ret, err = bulletproofs.PedersenVerify(commitmentMult, openingMult, 1000)
	fmt.Println("Mul num commit: ", ret)

	randomOpening, err := bulletproofs.PedersenRNG()
	if err != nil {
		panic(err)
	}
	commitmentRandom, err := bulletproofs.PedersenCommitSpecificOpening(100, randomOpening)
	if err != nil {
		panic(err)
	}
	ret, err = bulletproofs.PedersenVerify(commitmentRandom, randomOpening, 100)
	fmt.Println("random commitment: ", ret)

	commitment0, opening0, err := bulletproofs.PedersenCommitRandomOpening(0)
	if err != nil {
		panic(err)
	}
	commitment0Neg, err := bulletproofs.PedersenNeg(commitment0)
	if err != nil {
		panic(err)
	}
	opening0Neg, err := bulletproofs.PedersenNegOpening(opening0)
	if err != nil {
		panic(err)
	}
	ret, err = bulletproofs.PedersenVerify(commitment0Neg, opening0Neg, 0)
	fmt.Println("Negated commitment: ", ret)
}
