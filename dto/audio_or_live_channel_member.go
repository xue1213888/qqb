package dto

type AudioOrLiveChannelMemberData struct {
	// 频道ID
	GuildID string `json:"guild_id"`
	// 用户ID
	UserID string `json:"user_id"`
	// 子频道ID
	ChannelID string `json:"channel_id"`
	// 频道类型 2音视频 5直播
	ChannelType int `json:"channel_type"`
}
