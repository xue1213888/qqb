package event

import (
	"github.com/tencent-connect/botgo/dto"
)

// DefaultHandlers 默认的 handler 结构，管理所有支持的 handler 类型
var DefaultHandlers struct {
	Ready       ReadyHandler
	ErrorNotify ErrorNotifyHandler
	Plain       PlainEventHandler

	Guild       GuildEventHandler
	GuildMember GuildMemberEventHandler
	Channel     ChannelEventHandler

	Message             MessageEventHandler
	MessageReaction     MessageReactionEventHandler
	ATMessage           ATMessageEventHandler
	DirectMessage       DirectMessageEventHandler
	MessageAudit        MessageAuditEventHandler
	MessageDelete       MessageDeleteEventHandler
	PublicMessageDelete PublicMessageDeleteEventHandler
	DirectMessageDelete DirectMessageDeleteEventHandler

	Audio                    AudioEventHandler
	AudioOrLiveChannelMember AudioOrLiveChannelMemberHandler

	Thread     ThreadEventHandler
	OpenThread OpenThreadEventHandler
	Post       PostEventHandler
	OpenPost   OpenPostEventHandler
	Reply      ReplyEventHandler
	OpenReply  OpenReplyEventHandler
	ForumAudit ForumAuditEventHandler

	Interaction InteractionEventHandler

	FriendAdd            QQFriendAddHandler
	C2cMsgReject         QQC2cMsgRejectHandler
	C2cMsgReceive        QQC2cMsgReceiveHandler
	GroupAddRobot        QQGroupAddRobotHandler
	GroupDelRobot        QQGroupDelRobotHandler
	GroupMsgReject       QQGroupMsgRejectHandler
	GroupMsgReceive      QQGroupMsgReceiveHandler
	C2cMessageCreate     QQC2cMessageCreateHandler
	GroupAtMessageCreate QQGroupAtMessageCreateHandler

	//dto.EventFriendAdd:            qqFriendAdd,
	//dto.EventC2cMsgReject:         qqC2cMsgReject,
	//dto.EventC2cMsgReceive:        qqC2cMsgReceive,
	//dto.EventGroupAddRobot:        qqGroupAddRobot,
	//dto.EventGroupDelRobot:        qqGroupDelRobot,
	//dto.EventGroupMsgReject:       qqGroupMsgReject,
	//dto.EventGroupMsgReceive:      qqGroupMsgReceive,
	//dto.EventC2cMessageCreate:     qqC2cMessageCreate,
	//dto.EventGroupAtMessageCreate: qqGroupAtMessageCreate,
}

// ReadyHandler 可以处理 ws 的 ready 事件
type ReadyHandler func(event *dto.WSPayload, data *dto.WSReadyData)

// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
type ErrorNotifyHandler func(err error)

// PlainEventHandler 透传handler
type PlainEventHandler func(event *dto.WSPayload, message []byte) error

// GuildEventHandler 频道事件handler
type GuildEventHandler func(event *dto.WSPayload, data *dto.WSGuildData) error

// GuildMemberEventHandler 频道成员事件 handler
type GuildMemberEventHandler func(event *dto.WSPayload, data *dto.WSGuildMemberData) error

// ChannelEventHandler 子频道事件 handler
type ChannelEventHandler func(event *dto.WSPayload, data *dto.WSChannelData) error

// MessageEventHandler 消息事件 handler
type MessageEventHandler func(event *dto.WSPayload, data *dto.WSMessageData) error

// MessageDeleteEventHandler 消息事件 handler
type MessageDeleteEventHandler func(event *dto.WSPayload, data *dto.WSMessageDeleteData) error

// PublicMessageDeleteEventHandler 消息事件 handler
type PublicMessageDeleteEventHandler func(event *dto.WSPayload, data *dto.WSPublicMessageDeleteData) error

// DirectMessageDeleteEventHandler 消息事件 handler
type DirectMessageDeleteEventHandler func(event *dto.WSPayload, data *dto.WSDirectMessageDeleteData) error

// AudioOrLiveChannelMemberHandler 音视频频道成员进入退出事件 handler
type AudioOrLiveChannelMemberHandler func(event *dto.WSPayload, data *dto.WSAudioOrLiveChannelMemberData) error

// MessageReactionEventHandler 表情表态事件 handler
type MessageReactionEventHandler func(event *dto.WSPayload, data *dto.WSMessageReactionData) error

// ATMessageEventHandler at 机器人消息事件 handler
type ATMessageEventHandler func(event *dto.WSPayload, data *dto.WSATMessageData) error

// DirectMessageEventHandler 私信消息事件 handler
type DirectMessageEventHandler func(event *dto.WSPayload, data *dto.WSDirectMessageData) error

// AudioEventHandler 音频机器人事件 handler
type AudioEventHandler func(event *dto.WSPayload, data *dto.WSAudioData) error

// MessageAuditEventHandler 消息审核事件 handler
type MessageAuditEventHandler func(event *dto.WSPayload, data *dto.WSMessageAuditData) error

// ThreadEventHandler 论坛主题事件 handler
type ThreadEventHandler func(event *dto.WSPayload, data *dto.WSThreadData) error
type OpenThreadEventHandler func(event *dto.WSPayload, data *dto.WSThreadData) error

// PostEventHandler 论坛回帖事件 handler
type PostEventHandler func(event *dto.WSPayload, data *dto.WSPostData) error
type OpenPostEventHandler func(event *dto.WSPayload, data *dto.WSPostData) error

// ReplyEventHandler 论坛帖子回复事件 handler
type ReplyEventHandler func(event *dto.WSPayload, data *dto.WSReplyData) error
type OpenReplyEventHandler func(event *dto.WSPayload, data *dto.WSReplyData) error

// ForumAuditEventHandler 论坛帖子审核事件 handler
type ForumAuditEventHandler func(event *dto.WSPayload, data *dto.WSForumAuditData) error

// InteractionEventHandler 互动事件 handler
type InteractionEventHandler func(event *dto.WSPayload, data *dto.WSInteractionData) error

type QQFriendAddHandler func(event *dto.WSPayload, data *dto.WSQQFriendAddData) error
type QQC2cMsgRejectHandler func(event *dto.WSPayload, data *dto.WSQQC2cMsgRejectData) error
type QQC2cMsgReceiveHandler func(event *dto.WSPayload, data *dto.WSQQC2cMsgReceiveData) error
type QQGroupAddRobotHandler func(event *dto.WSPayload, data *dto.WSQQGroupAddRobotData) error
type QQGroupDelRobotHandler func(event *dto.WSPayload, data *dto.WSQQGroupDelRobotData) error
type QQGroupMsgRejectHandler func(event *dto.WSPayload, data *dto.WSQQGroupMsgRejectData) error
type QQGroupMsgReceiveHandler func(event *dto.WSPayload, data *dto.WSQQGroupMsgReceiveData) error
type QQC2cMessageCreateHandler func(event *dto.WSPayload, data *dto.WSQQC2cMessageCreateData) error
type QQGroupAtMessageCreateHandler func(event *dto.WSPayload, data *dto.WSMessageData) error

// RegisterHandlers 注册事件回调，并返回 intent 用于 websocket 的鉴权
func RegisterHandlers(handlers ...interface{}) dto.Intent {
	var i dto.Intent
	for _, h := range handlers {
		switch handle := h.(type) {
		case ReadyHandler:
			DefaultHandlers.Ready = handle
		case ErrorNotifyHandler:
			DefaultHandlers.ErrorNotify = handle
		case PlainEventHandler:
			DefaultHandlers.Plain = handle
		case AudioEventHandler:
			DefaultHandlers.Audio = handle
			i = i | dto.EventToIntent(
				dto.EventAudioStart, dto.EventAudioFinish,
				dto.EventAudioOnMic, dto.EventAudioOffMic,
			)
		case InteractionEventHandler:
			DefaultHandlers.Interaction = handle
			i = i | dto.EventToIntent(dto.EventInteractionCreate)
		case AudioOrLiveChannelMemberHandler:
			DefaultHandlers.AudioOrLiveChannelMember = handle
			i = i | dto.EventToIntent(dto.EventAudioOrLiveChannelMemberEnter, dto.EventAudioOrLiveChannelMemberExit)
		default:
		}
	}
	i = i | registerRelationHandlers(i, handlers...)
	i = i | registerMessageHandlers(i, handlers...)
	i = i | registerForumHandlers(i, handlers...)
	i = i | registerOpenForumHandlers(i, handlers...)
	i = i | registerQQBotHandlers(i, handlers...)
	return i
}

func registerQQBotHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case QQFriendAddHandler:
			DefaultHandlers.FriendAdd = handle
			i = i | dto.EventToIntent(dto.EventFriendAdd)
		case QQC2cMsgRejectHandler:
			DefaultHandlers.C2cMsgReject = handle
			i = i | dto.EventToIntent(dto.EventC2cMsgReject)
		case QQC2cMsgReceiveHandler:
			DefaultHandlers.C2cMsgReceive = handle
			i = i | dto.EventToIntent(dto.EventC2cMsgReceive)
		case QQGroupAddRobotHandler:
			DefaultHandlers.GroupAddRobot = handle
			i = i | dto.EventToIntent(dto.EventGroupAddRobot)
		case QQGroupDelRobotHandler:
			DefaultHandlers.GroupDelRobot = handle
			i = i | dto.EventToIntent(dto.EventGroupDelRobot)
		case QQGroupMsgRejectHandler:
			DefaultHandlers.GroupMsgReject = handle
			i = i | dto.EventToIntent(dto.EventGroupMsgReject)
		case QQGroupMsgReceiveHandler:
			DefaultHandlers.GroupMsgReceive = handle
			i = i | dto.EventToIntent(dto.EventGroupMsgReceive)
		case QQC2cMessageCreateHandler:
			DefaultHandlers.C2cMessageCreate = handle
			i = i | dto.EventToIntent(dto.EventC2cMessageCreate)
		case QQGroupAtMessageCreateHandler:
			DefaultHandlers.GroupAtMessageCreate = handle
			i = i | dto.EventToIntent(dto.EventGroupAtMessageCreate)
		}
	}
	return i
}

func registerOpenForumHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case OpenThreadEventHandler:
			DefaultHandlers.OpenThread = handle
			i = i | dto.EventToIntent(
				dto.EventOpenForumThreadCreate, dto.EventOpenForumThreadUpdate, dto.EventOpenForumThreadDelete,
			)
		case OpenPostEventHandler:
			DefaultHandlers.OpenPost = handle
			i = i | dto.EventToIntent(dto.EventOpenForumPostCreate, dto.EventOpenForumPostDelete)
		case OpenReplyEventHandler:
			DefaultHandlers.OpenReply = handle
			i = i | dto.EventToIntent(dto.EventOpenForumReplyCreate, dto.EventOpenForumReplyDelete)
		default:
		}
	}
	return i
}

func registerForumHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case ThreadEventHandler:
			DefaultHandlers.Thread = handle
			i = i | dto.EventToIntent(
				dto.EventForumThreadCreate, dto.EventForumThreadUpdate, dto.EventForumThreadDelete,
			)
		case PostEventHandler:
			DefaultHandlers.Post = handle
			i = i | dto.EventToIntent(dto.EventForumPostCreate, dto.EventForumPostDelete)
		case ReplyEventHandler:
			DefaultHandlers.Reply = handle
			i = i | dto.EventToIntent(dto.EventForumReplyCreate, dto.EventForumReplyDelete)
		case ForumAuditEventHandler:
			DefaultHandlers.ForumAudit = handle
			i = i | dto.EventToIntent(dto.EventForumAuditResult)
		default:
		}
	}
	return i
}

// registerRelationHandlers 注册频道关系链相关handlers
func registerRelationHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case GuildEventHandler:
			DefaultHandlers.Guild = handle
			i = i | dto.EventToIntent(dto.EventGuildCreate, dto.EventGuildDelete, dto.EventGuildUpdate)
		case GuildMemberEventHandler:
			DefaultHandlers.GuildMember = handle
			i = i | dto.EventToIntent(dto.EventGuildMemberAdd, dto.EventGuildMemberRemove, dto.EventGuildMemberUpdate)
		case ChannelEventHandler:
			DefaultHandlers.Channel = handle
			i = i | dto.EventToIntent(dto.EventChannelCreate, dto.EventChannelDelete, dto.EventChannelUpdate)
		default:
		}
	}
	return i
}

// registerMessageHandlers 注册消息相关的 handler
func registerMessageHandlers(i dto.Intent, handlers ...interface{}) dto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case MessageEventHandler:
			DefaultHandlers.Message = handle
			i = i | dto.EventToIntent(dto.EventMessageCreate)
		case ATMessageEventHandler:
			DefaultHandlers.ATMessage = handle
			i = i | dto.EventToIntent(dto.EventAtMessageCreate)
		case DirectMessageEventHandler:
			DefaultHandlers.DirectMessage = handle
			i = i | dto.EventToIntent(dto.EventDirectMessageCreate)
		case MessageDeleteEventHandler:
			DefaultHandlers.MessageDelete = handle
			i = i | dto.EventToIntent(dto.EventMessageDelete)
		case PublicMessageDeleteEventHandler:
			DefaultHandlers.PublicMessageDelete = handle
			i = i | dto.EventToIntent(dto.EventPublicMessageDelete)
		case DirectMessageDeleteEventHandler:
			DefaultHandlers.DirectMessageDelete = handle
			i = i | dto.EventToIntent(dto.EventDirectMessageDelete)
		case MessageReactionEventHandler:
			DefaultHandlers.MessageReaction = handle
			i = i | dto.EventToIntent(dto.EventMessageReactionAdd, dto.EventMessageReactionRemove)
		case MessageAuditEventHandler:
			DefaultHandlers.MessageAudit = handle
			i = i | dto.EventToIntent(dto.EventMessageAuditPass, dto.EventMessageAuditReject)
		default:
		}
	}
	return i
}
