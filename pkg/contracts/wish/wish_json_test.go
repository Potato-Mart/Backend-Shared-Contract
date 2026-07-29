package wish_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/common"
	"github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/contracts/wish"
	wishenum "github.com/Potato-Mart/Backend-Shared-Contract/v21/pkg/enums/wish"
)

func TestWishProposalJSONIsIdentityFree(t *testing.T) {
	now := time.Date(2026, 7, 20, 1, 2, 3, 0, time.UTC)
	proposal := wish.WishProposal{
		ID:                   "proposal_1",
		ProductName:          "Taiwanese pineapple cake",
		Description:          "A less-sweet option",
		ReferenceURL:         "https://example.com/product",
		State:                wishenum.WishProposalStateConverted,
		ConvertedCandidateID: "candidate_1",
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	payload := marshalWishJSON(t, proposal)
	for _, want := range []string{
		`"product_name":"Taiwanese pineapple cake"`,
		`"state":"converted"`,
		`"converted_candidate_id":"candidate_1"`,
		`"created_at":"2026-07-20T01:02:03Z"`,
	} {
		if !strings.Contains(payload, want) {
			t.Fatalf("wish proposal JSON = %s, want %s", payload, want)
		}
	}
	assertWishJSONOmits(t, payload, "customer_id", "user_id", "account_id", "created_by", "updated_by")
}

func TestWishCandidateJSONUsesApprovedLocalizedContent(t *testing.T) {
	now := time.Date(2026, 7, 20, 1, 2, 3, 0, time.UTC)
	candidate := wish.WishCandidate{
		ID: "candidate_1",
		Name: []common.LocalizedName{
			{Language: "en", Name: "Pineapple cake"},
			{Language: "zh-TW", Name: "鳳梨酥"},
		},
		Description: []common.LocalizedDescription{{Language: "en", Description: "Less sweet"}},
		ImageURLs:   []string{"https://cdn.example.com/wishes/candidate_1.jpg"},
		State:       wishenum.WishCandidateStatePublished,
		PublishedAt: &now,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	payload := marshalWishJSON(t, candidate)
	for _, want := range []string{
		`"name":[`, `"language":"zh-TW"`, `"name":"鳳梨酥"`,
		`"description":[`, `"image_urls":[`, `"state":"published"`,
		`"published_at":"2026-07-20T01:02:03Z"`,
	} {
		if !strings.Contains(payload, want) {
			t.Fatalf("wish candidate JSON = %s, want %s", payload, want)
		}
	}
	assertWishJSONOmits(t, payload,
		"customer_id", "user_id", "created_by", "updated_by", "deleted_by", "fulfilled_at")
}

func TestWishBallotRankingAndSelectionRoundTrip(t *testing.T) {
	opensAt := time.Date(2026, 7, 21, 0, 0, 0, 0, time.UTC)
	closesAt := opensAt.Add(7 * 24 * time.Hour)
	ballot := wish.WishBallot{
		ID:           "ballot_1",
		State:        wishenum.WishBallotStateOpen,
		CandidateIDs: []string{"candidate_1", "candidate_2"},
		OpensAt:      opensAt,
		ClosesAt:     closesAt,
		Revision:     4,
		AsOf:         opensAt,
	}
	selection := wish.WishSelection{
		BallotID:     ballot.ID,
		CandidateIDs: []string{"candidate_2"},
		UpdatedAt:    opensAt,
	}
	ranking := []wish.WishRankingEntry{
		{CandidateID: "candidate_2", Rank: 1, VoteCount: 23},
		{CandidateID: "candidate_1", Rank: 2, VoteCount: 17},
	}

	var decodedBallot wish.WishBallot
	roundTripWishJSON(t, ballot, &decodedBallot)
	if decodedBallot.Revision != 4 || len(decodedBallot.CandidateIDs) != 2 || decodedBallot.CandidateIDs[0] != "candidate_1" {
		t.Fatalf("wish ballot did not round-trip: %+v", decodedBallot)
	}

	var decodedSelection wish.WishSelection
	roundTripWishJSON(t, selection, &decodedSelection)
	if decodedSelection.BallotID != "ballot_1" || len(decodedSelection.CandidateIDs) != 1 || decodedSelection.CandidateIDs[0] != "candidate_2" {
		t.Fatalf("wish selection did not round-trip: %+v", decodedSelection)
	}

	var decodedRanking []wish.WishRankingEntry
	roundTripWishJSON(t, ranking, &decodedRanking)
	if len(decodedRanking) != 2 || decodedRanking[0].Rank != 1 || decodedRanking[0].VoteCount != 23 {
		t.Fatalf("wish ranking did not round-trip in order: %+v", decodedRanking)
	}

	for _, payload := range []string{marshalWishJSON(t, ballot), marshalWishJSON(t, selection), marshalWishJSON(t, ranking)} {
		assertWishJSONOmits(t, payload, "customer_id", "user_id", "account_id", "created_by", "updated_by")
	}
}

func marshalWishJSON(t *testing.T, value any) string {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal wish model: %v", err)
	}
	return string(payload)
}

func roundTripWishJSON(t *testing.T, source, target any) {
	t.Helper()
	payload, err := json.Marshal(source)
	if err != nil {
		t.Fatalf("marshal wish model: %v", err)
	}
	if err := json.Unmarshal(payload, target); err != nil {
		t.Fatalf("unmarshal wish model: %v", err)
	}
}

func assertWishJSONOmits(t *testing.T, payload string, keys ...string) {
	t.Helper()
	for _, key := range keys {
		if strings.Contains(payload, `"`+key+`"`) {
			t.Fatalf("wish JSON unexpectedly exposes %q: %s", key, payload)
		}
	}
}
