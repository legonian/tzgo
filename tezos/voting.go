// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tezos

import (
	"fmt"
)

// VotingPeriodKind represents a named voting period in Tezos.
type VotingPeriodKind byte

const (
	VotingPeriodInvalid VotingPeriodKind = iota
	VotingPeriodProposal
	VotingPeriodExploration
	VotingPeriodCooldown
	VotingPeriodPromotion
	VotingPeriodAdoption
)

var VotingPeriods = []VotingPeriodKind{
	VotingPeriodProposal,
	VotingPeriodExploration,
	VotingPeriodCooldown,
	VotingPeriodPromotion,
	VotingPeriodAdoption,
}

func (v VotingPeriodKind) IsValid() bool {
	return v != VotingPeriodInvalid
}

func (v *VotingPeriodKind) UnmarshalText(data []byte) error {
	vv := ParseVotingPeriod(string(data))
	if !vv.IsValid() {
		return fmt.Errorf("tezos: invalid voting period '%s'", string(data))
	}
	*v = vv
	return nil
}

func (v VotingPeriodKind) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v VotingPeriodKind) Num() int {
	switch v {
	case VotingPeriodProposal:
		return 1
	case VotingPeriodExploration:
		return 2
	case VotingPeriodCooldown:
		return 3
	case VotingPeriodPromotion:
		return 4
	case VotingPeriodAdoption:
		return 5
	default:
		return 1
	}
}

func ToVotingPeriod(i int) VotingPeriodKind {
	switch i {
	case 2:
		return VotingPeriodExploration
	case 3:
		return VotingPeriodCooldown
	case 4:
		return VotingPeriodPromotion
	case 5:
		return VotingPeriodAdoption
	default:
		return VotingPeriodProposal
	}
}

func ParseVotingPeriod(s string) VotingPeriodKind {
	switch s {
	case "proposal":
		return VotingPeriodProposal
	case "testing_vote", "exploration":
		return VotingPeriodExploration
	case "testing", "cooldown":
		return VotingPeriodCooldown
	case "promotion_vote", "promotion":
		return VotingPeriodPromotion
	case "adoption":
		return VotingPeriodAdoption
	default:
		return VotingPeriodInvalid
	}
}

func (v VotingPeriodKind) String() string {
	switch v {
	case VotingPeriodProposal:
		return "proposal"
	case VotingPeriodExploration:
		return "exploration"
	case VotingPeriodCooldown:
		return "cooldown"
	case VotingPeriodPromotion:
		return "promotion"
	case VotingPeriodAdoption:
		return "adoption"
	default:
		return ""
	}
}

// BallotVote represents a named ballot in Tezos.
type BallotVote byte

const (
	BallotVoteInvalid BallotVote = iota
	BallotVoteYay
	BallotVoteNay
	BallotVotePass
)

func (v BallotVote) IsValid() bool {
	return v != BallotVoteInvalid
}

func (v *BallotVote) UnmarshalText(data []byte) error {
	vv := ParseBallotVote(string(data))
	if !vv.IsValid() {
		return fmt.Errorf("tezos: invalid ballot %q", string(data))
	}
	*v = vv
	return nil
}

func (v BallotVote) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v *BallotVote) UnmarshalBinary(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("tezos: short ballot data")
	}
	vv := ParseBallotTag(data[0])
	if !vv.IsValid() {
		return fmt.Errorf("tezos: invalid ballot tag %d", data[0])
	}
	*v = vv
	return nil
}

func ParseBallotVote(s string) BallotVote {
	switch s {
	case "yay":
		return BallotVoteYay
	case "nay":
		return BallotVoteNay
	case "pass":
		return BallotVotePass
	default:
		return BallotVoteInvalid
	}
}

func (v BallotVote) String() string {
	switch v {
	case BallotVoteYay:
		return "yay"
	case BallotVoteNay:
		return "nay"
	case BallotVotePass:
		return "pass"
	default:
		return ""
	}
}

func (v BallotVote) Tag() byte {
	switch v {
	case BallotVoteYay:
		return 0
	case BallotVoteNay:
		return 1
	case BallotVotePass:
		return 2
	default:
		return 255
	}
}

func ParseBallotTag(t byte) BallotVote {
	switch t {
	case 0:
		return BallotVoteYay
	case 1:
		return BallotVoteNay
	case 2:
		return BallotVotePass
	default:
		return BallotVoteInvalid
	}
}
