/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bulletproofs

var Helper func() Provider

type Provider interface {
	NewBulletproofs() bulletproofs
}

type bulletproofs interface {
	ProveOps
	CommitmentOps
}

type ProveOps interface {
	ProveRandomOpening(x uint64) ([]byte, []byte, []byte, error)
	ProveSpecificOpening(x uint64, opening []byte) ([]byte, []byte, error)
	Verify(proof []byte, commitment []byte) (bool, error)
	ProveAfterAddNum(x, y uint64, openingX, commitmentX []byte) ([]byte, []byte, error)
	ProveAfterAddCommitment(x, y uint64, openingX, openingY, commitmentX, commitmentY []byte) ([]byte, []byte, []byte, error)
	ProveAfterSubNum(x, y uint64, openingX, commitmentX []byte) ([]byte, []byte, error)
	ProveAfterSubCommitment(x, y uint64, openingX, openingY, commitmentX, commitmentY []byte) ([]byte, []byte, []byte, error)
	ProveAfterMulNum(x, y uint64, openingX, commitmentX []byte) ([]byte, []byte, []byte, error)
}

type CommitmentOps interface {
	PedersenRNG() ([]byte, error)
	PedersenCommitRandomOpening(x uint64) ([]byte, []byte, error)
	PedersenCommitSpecificOpening(x uint64, r []byte) ([]byte, error)
	PedersenVerify(commitment, opening []byte, value uint64) (bool, error)
	PedersenNeg(commitment []byte) ([]byte, error)
	PedersenNegOpening(opening []byte) ([]byte, error)
	PedersenAddNum(commitment []byte, value uint64) ([]byte, error)
	PedersenAddCommitment(commitment1, commitment2 []byte) ([]byte, error)
	PedersenAddOpening(opening1, opening2 []byte) ([]byte, error)
	PedersenAddCommitmentWithOpening(commitment1, commitment2, opening1, opening2 []byte) ([]byte, []byte, error)
	PedersenSubNum(commitment []byte, value uint64) ([]byte, error)
	PedersenSubCommitment(commitment1, commitment2 []byte) ([]byte, error)
	PedersenSubOpening(opening1, opening2 []byte) ([]byte, error)
	PedersenSubCommitmentWithOpening(commitment1, commitment2, opening1, opening2 []byte) ([]byte, []byte, error)
	PedersenMulNum(commitment1 []byte, value uint64) ([]byte, error)
	PedersenMulOpening(opening1 []byte, value uint64) ([]byte, error)
	PedersenMulNumWithOpening(commitment []byte, opening []byte, value uint64) ([]byte, []byte, error)
}
