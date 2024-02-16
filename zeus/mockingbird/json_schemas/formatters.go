package mb_json_schemas

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
	mb_search "github.com/zeus-fyi/zeus/zeus/mockingbird/search"
)

func FormatSearchResultsV5(results []mb_search.SearchResult) string {
	if len(results) == 0 {
		return ""
	}
	var newResults []interface{}
	for _, result := range results {
		if result.WebResponse.Body != nil {
			if result.Value != "" {
				result.WebResponse.Body["msg_body"] = result.Value
			}
			newResults = append(newResults, result.WebResponse.Body)
		}
		if result.Verified != nil && *result.Verified && result.UnixTimestamp > 0 {
			msgID := fmt.Sprintf("%d", result.UnixTimestamp)
			if result.TwitterMetadata != nil && result.TwitterMetadata.TweetStrID != "" {
				msgID = result.TwitterMetadata.TweetStrID
			}
			nr := mb_search.SimplifiedSearchResultJSON{
				MessageID:   msgID,
				MessageBody: result.Value,
			}
			newResults = append(newResults, nr)
		} else {
			m := map[string]interface{}{
				"msg_id":   fmt.Sprintf("%d", result.UnixTimestamp),
				"msg_body": result.Value,
			}
			newResults = append(newResults, m)
		}
	}
	b, err := json.Marshal(newResults)
	if err != nil {
		log.Err(err).Msg("FormatSearchResultsV5: Error marshalling search results")
		return ""
	}
	return string(b)
}
