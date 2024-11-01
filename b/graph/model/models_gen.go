// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Element interface {
	IsElement()
	GetID() string
}

type Post interface {
	IsPost()
	GetID() string
}

type CardPostCommunityRecommendation struct {
	Posts *PostConnection `json:"posts,omitempty"`
}

type CardPostCommunityRecommendationsFeedUnit struct {
	ID                       string                             `json:"id"`
	CommunityRecommendations []*CardPostCommunityRecommendation `json:"communityRecommendations"`
}

func (CardPostCommunityRecommendationsFeedUnit) IsElement()         {}
func (this CardPostCommunityRecommendationsFeedUnit) GetID() string { return this.ID }

type CompactPostCommunityRecommendation struct {
	Posts *PostConnection `json:"posts,omitempty"`
}

type CompactPostCommunityRecommendationsFeedUnit struct {
	ID                       string                                `json:"id"`
	CommunityRecommendations []*CompactPostCommunityRecommendation `json:"communityRecommendations"`
}

func (CompactPostCommunityRecommendationsFeedUnit) IsElement()         {}
func (this CompactPostCommunityRecommendationsFeedUnit) GetID() string { return this.ID }

type ElementConnection struct {
	Edges []*ElementEdge `json:"edges"`
}

type ElementEdge struct {
	Node Element `json:"node,omitempty"`
}

type MediaSource struct {
	ID string `json:"id"`
}

type PostConnection struct {
	Edges []*PostEdge `json:"edges"`
}

type PostEdge struct {
	Node Post `json:"node,omitempty"`
}

type Query struct {
}

type SDWatchFeed struct {
	Elements *ElementConnection `json:"elements,omitempty"`
}

type StillMedia struct {
	Content *MediaSource `json:"content,omitempty"`
}

type SubredditPost struct {
	ID string `json:"id"`
}

func (SubredditPost) IsPost()            {}
func (this SubredditPost) GetID() string { return this.ID }

func (SubredditPost) IsEntity() {}
