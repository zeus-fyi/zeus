package mb_search

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type RetrievalItem struct {
	RetrievalStrID           *string                           `json:"retrievalStrID"`
	RetrievalID              *int                              `json:"retrievalID,omitempty"` // ID of the retrieval
	RetrievalName            string                            `json:"retrievalName"`         // Name of the retrieval
	RetrievalGroup           string                            `json:"retrievalGroup"`        // Group of the retrieval
	RetrievalItemInstruction `json:"retrievalItemInstruction"` // Instructions for the retrieval
}

type RetrievalItemInstruction struct {
	RetrievalPlatform         string      `json:"retrievalPlatform"`
	RetrievalPrompt           *string     `json:"retrievalPrompt,omitempty"`           // Prompt for the retrieval
	RetrievalPlatformGroups   *string     `json:"retrievalPlatformGroups,omitempty"`   // Platform groups for the retrieval
	RetrievalKeywords         *string     `json:"retrievalKeywords,omitempty"`         // Keywords for the retrieval
	RetrievalNegativeKeywords *string     `json:"retrievalNegativeKeywords,omitempty"` // Keywords for the retrieval
	RetrievalUsernames        *string     `json:"retrievalUsernames,omitempty"`        // Usernames for the retrieval
	WebFilters                *WebFilters `json:"webFilters,omitempty"`                // Web filters for the retrieval

	Instructions json.RawMessage `json:"instructions,omitempty"` // Instructions for the retrieval
}

func SetInstructions(r *RetrievalItem) error {
	b, err := json.Marshal(r.RetrievalItemInstruction)
	if err != nil {
		log.Err(err).Msg("failed to marshal retrieval instructions")
		return err
	}
	r.Instructions = b
	return nil
}
