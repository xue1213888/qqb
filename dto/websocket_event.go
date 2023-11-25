package dto

func init() {
	eventIntentMap = transposeIntentEventMap(intentEventMap)
}

// 事件类型
const (
	// 频道事件 1 << 0
	EventGuildCreate EventType = "GUILD_CREATE"
	EventGuildUpdate EventType = "GUILD_UPDATE"
	EventGuildDelete EventType = "GUILD_DELETE"

	// 子频道事件 1 << 0
	EventChannelCreate EventType = "CHANNEL_CREATE"
	EventChannelUpdate EventType = "CHANNEL_UPDATE"
	EventChannelDelete EventType = "CHANNEL_DELETE"

	// 频道成员事件 1 << 1
	EventGuildMemberAdd    EventType = "GUILD_MEMBER_ADD"
	EventGuildMemberUpdate EventType = "GUILD_MEMBER_UPDATE"
	EventGuildMemberRemove EventType = "GUILD_MEMBER_REMOVE"

	// 私欲消息事件 1 << 9
	EventMessageCreate EventType = "MESSAGE_CREATE"
	EventMessageDelete EventType = "MESSAGE_DELETE"

	// 表情表态事件 1 << 10
	EventMessageReactionAdd    EventType = "MESSAGE_REACTION_ADD"
	EventMessageReactionRemove EventType = "MESSAGE_REACTION_REMOVE"

	// 私信事件 1 << 12
	EventDirectMessageCreate EventType = "DIRECT_MESSAGE_CREATE"
	EventDirectMessageDelete EventType = "DIRECT_MESSAGE_DELETE"

	// 公域消息时间 1 << 30
	EventAtMessageCreate     EventType = "AT_MESSAGE_CREATE"
	EventPublicMessageDelete EventType = "PUBLIC_MESSAGE_DELETE"

	// 消息审核事件 1 << 27
	EventMessageAuditPass   EventType = "MESSAGE_AUDIT_PASS"
	EventMessageAuditReject EventType = "MESSAGE_AUDIT_REJECT"

	// 音频事件 1 << 29
	EventAudioStart  EventType = "AUDIO_START"
	EventAudioFinish EventType = "AUDIO_FINISH"
	EventAudioOnMic  EventType = "AUDIO_ON_MIC"
	EventAudioOffMic EventType = "AUDIO_OFF_MIC"

	// 论坛事件 1 << 28 私域
	EventForumThreadCreate EventType = "FORUM_THREAD_CREATE"
	EventForumThreadUpdate EventType = "FORUM_THREAD_UPDATE"
	EventForumThreadDelete EventType = "FORUM_THREAD_DELETE"
	// 帖子事件 1 << 28
	EventForumPostCreate EventType = "FORUM_POST_CREATE"
	EventForumPostDelete EventType = "FORUM_POST_DELETE"
	// 帖子回复事件 1 << 28
	EventForumReplyCreate EventType = "FORUM_REPLY_CREATE"
	EventForumReplyDelete EventType = "FORUM_REPLY_DELETE"
	// 论坛审核事件 1 << 28
	EventForumAuditResult EventType = "FORUM_PUBLISH_AUDIT_RESULT"

	// 论坛事件 1 << 18 公域
	EventOpenForumThreadCreate EventType = "OPEN_FORUM_THREAD_CREATE"
	EventOpenForumThreadUpdate EventType = "OPEN_FORUM_THREAD_UPDATE"
	EventOpenForumThreadDelete EventType = "OPEN_FORUM_THREAD_DELETE"
	// 帖子事件 1 << 18
	EventOpenForumPostCreate EventType = "OPEN_FORUM_POST_CREATE"
	EventOpenForumPostDelete EventType = "OPEN_FORUM_POST_DELETE"
	// 帖子回复事件 1 << 18
	EventOpenForumReplyCreate EventType = "OPEN_FORUM_REPLY_CREATE"
	EventOpenForumReplyDelete EventType = "OPEN_FORUM_REPLY_DELETE"

	// 互动事件创建，就是消息按钮之类的 1 << 26
	EventInteractionCreate EventType = "INTERACTION_CREATE"

	// 音频用户进入退出直播间 1 << 19
	EventAudioOrLiveChannelMemberEnter = "AUDIO_OR_LIVE_CHANNEL_MEMBER_ENTER"
	EventAudioOrLiveChannelMemberExit  = "AUDIO_OR_LIVE_CHANNEL_MEMBER_EXIT"

	// 群相关事件 1 << 25
	EventFriendAdd            EventType = "FRIEND_ADD"              // 好友添加
	EventC2cMsgReject         EventType = "C2C_MSG_REJECT"          // 拒绝私聊主动消息
	EventC2cMsgReceive        EventType = "C2C_MSG_RECEIVE"         // 同意私聊主动消息
	EventGroupAddRobot        EventType = "GROUP_ADD_ROBOT"         // 群内加入机器人
	EventGroupDelRobot        EventType = "GROUP_DEL_ROBOT"         // 群内踢出机器人
	EventGroupMsgReject       EventType = "GROUP_MSG_REJECT"        // 群内拒绝主动消息
	EventGroupMsgReceive      EventType = "GROUP_MSG_RECEIVE"       // 群内同意主动消息
	EventC2cMessageCreate     EventType = "C2C_MESSAGE_CREATE"      // 私聊消息
	EventGroupAtMessageCreate EventType = "GROUP_AT_MESSAGE_CREATE" // 群内@消息
)

// intentEventMap 不同 intent 对应的事件定义
var intentEventMap = map[Intent][]EventType{
	IntentGuilds: {
		EventGuildCreate, EventGuildUpdate, EventGuildDelete,
		EventChannelCreate, EventChannelUpdate, EventChannelDelete,
	},
	IntentGuildMembers:          {EventGuildMemberAdd, EventGuildMemberUpdate, EventGuildMemberRemove},
	IntentGuildMessages:         {EventMessageCreate, EventMessageDelete},
	IntentGuildMessageReactions: {EventMessageReactionAdd, EventMessageReactionRemove},
	IntentGuildAtMessage:        {EventAtMessageCreate, EventPublicMessageDelete},
	IntentDirectMessages:        {EventDirectMessageCreate, EventDirectMessageDelete},
	IntentAudio:                 {EventAudioStart, EventAudioFinish, EventAudioOnMic, EventAudioOffMic},
	IntentAudit:                 {EventMessageAuditPass, EventMessageAuditReject},
	IntentForum: {
		EventForumThreadCreate, EventForumThreadUpdate, EventForumThreadDelete, EventForumPostCreate,
		EventForumPostDelete, EventForumReplyCreate, EventForumReplyDelete, EventForumAuditResult,
	},
	IntentInteraction:              {EventInteractionCreate},
	IntentQQBot:                    {EventFriendAdd, EventC2cMsgReject, EventC2cMsgReceive, EventGroupAddRobot, EventGroupDelRobot, EventGroupMsgReject, EventGroupMsgReceive, EventC2cMessageCreate, EventGroupAtMessageCreate},
	IntentAudioOrLiveChannelMember: {EventAudioOrLiveChannelMemberEnter, EventAudioOrLiveChannelMemberExit},
	IntentOpenForumEvent:           {EventOpenForumThreadCreate, EventOpenForumThreadUpdate, EventOpenForumThreadDelete, EventOpenForumPostCreate, EventOpenForumPostDelete, EventOpenForumReplyCreate, EventOpenForumReplyDelete},
}

var eventIntentMap = transposeIntentEventMap(intentEventMap)

// transposeIntentEventMap 转置 intent 与 event 的关系，用于根据 event 找到 intent
func transposeIntentEventMap(input map[Intent][]EventType) map[EventType]Intent {
	result := make(map[EventType]Intent)
	for i, eventTypes := range input {
		for _, s := range eventTypes {
			result[s] = i
		}
	}
	return result
}

// EventToIntent 事件转换对应的Intent
func EventToIntent(events ...EventType) Intent {
	var i Intent
	for _, event := range events {
		i = i | eventIntentMap[event]
	}
	return i
}
