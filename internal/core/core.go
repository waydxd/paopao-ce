// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

// DataService 数据服务集成
type DataService interface {
	// 钱包服务
	WalletService

	// 消息服务
	MessageService

	// 话题服务
	TopicService

	// 推文服务
	TweetService
	TweetManageService
	TweetHelpService

	// 推文指标服务
	UserMetricServantA
	TweetMetricServantA
	CommentMetricServantA

	// 动态信息相关服务
	TrendsManageServantA

	// 评论服务
	CommentService
	CommentManageService

	// 用户服务
	UserManageService
	ContactManageService
	FollowingManageService
	UserRelationService

	// 安全服务
	SecurityService
	AttachmentCheckService

	// 社区服务
	CommunityService
	CommunityManageService
}

// WebDataServantA Web数据服务集成(版本A)
type WebDataServantA interface {
	// 话题服务
	TopicServantA

	// 推文服务
	TweetServantA
	TweetManageServantA
	TweetHelpServantA
}
