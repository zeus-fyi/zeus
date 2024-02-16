package mb_search

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type SearchResult struct {
	UnixTimestamp   int              `json:"unixTimestamp"`
	Source          string           `json:"source"`
	Value           string           `json:"value"`
	Group           string           `json:"group"`
	Verified        *bool            `json:"verified,omitempty"`
	RedditMetadata  RedditMetadata   `json:"redditMetadata"`
	TwitterMetadata *TwitterMetadata `json:"twitterMetadata,omitempty"`
	WebResponse     WebResponse      `json:"webResponses,omitempty"`
}

type TwitterMetadata struct {
	TweetStrID string `json:"in_reply_to_tweet_id"`
	Text       string `json:"text"`
}

type RedditMetadata struct {
	PostID           string `json:"postID"`
	FullPostID       string `json:"fullPostID"`
	NumberOfComments int    `json:"numberOfComments"`
	Url              string `json:"url"`
	Title            string `json:"title"`
	Body             string `json:"body"`
	Permalink        string
	Author           string          `json:"author"`
	AuthorID         string          `json:"authorID"`
	Subreddit        string          `json:"subreddit"`
	Score            int             `json:"score"`
	UpvoteRatio      float64         `json:"upvoteRatio"`
	Metadata         json.RawMessage `json:"metadata"`
}

type WebResponse struct {
	WebFilters *WebFilters `json:"webFilters,omitempty"`
	Body       echo.Map    `json:"body"`
	RawMessage []byte      `json:"rawMessage"`
}

type WebFilters struct {
	RoutingGroup       *string  `json:"routingGroup,omitempty"`
	LbStrategy         *string  `json:"lbStrategy,omitempty"`
	MaxRetries         *int     `json:"maxRetries,omitempty"`
	BackoffCoefficient *float64 `json:"backoffCoefficient,omitempty"`
	EndpointRoutePath  *string  `json:"endpointRoutePath,omitempty"`
	EndpointREST       *string  `json:"endpointREST,omitempty"`
}

type SimplifiedSearchResultJSON struct {
	MessageID   string `json:"msg_id"`
	MessageBody string `json:"msg_body"`
}
