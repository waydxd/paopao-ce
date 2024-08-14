// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cs

const (
	MetricActionCreateTweet uint8 = iota
	MetricActionDeleteTweet
)

type TweetMetric struct {
	PostId          int64
	CommentCount    int64
	UpvoteCount     int64
	CollectionCount int64
	ShareCount      int64
	ThumbsUpCount   int64
	ThumbsDownCount int64
}

type CommentMetric struct {
	CommentId       int64
	ReplyCount      int32
	ThumbsUpCount   int32
	ThumbsDownCount int32
}
type CommunityMetric struct {
	CommunityId int64
	PostCount   int64
	MemberCount int64
}

func (m *TweetMetric) RankScore(motivationFactor int) int64 {
	if motivationFactor == 0 {
		motivationFactor = 1
	}
	return (m.CommentCount + m.UpvoteCount*2 + m.CollectionCount*4 + m.ShareCount*8) * int64(motivationFactor)
}

func (m *CommentMetric) RankScore(motivationFactor int) int64 {
	if motivationFactor == 0 {
		motivationFactor = 1
	}
	return int64(m.ReplyCount*2+m.ThumbsUpCount*4-m.ThumbsDownCount) * int64(motivationFactor)
}
