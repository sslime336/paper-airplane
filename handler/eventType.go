package handler

const (
	EventGuildCreate           = "GUILD_CREATE"
	EventGuildUpdate           = "GUILD_UPDATE"
	EventGuildDelete           = "GUILD_DELETE"
	EventChannelCreate         = "CHANNEL_CREATE"
	EventChannelUpdate         = "CHANNEL_UPDATE"
	EventChannelDelete         = "CHANNEL_DELETE"
	EventGuildMemberAdd        = "GUILD_MEMBER_ADD"
	EventGuildMemberUpdate     = "GUILD_MEMBER_UPDATE"
	EventGuildMemberRemove     = "GUILD_MEMBER_REMOVE"
	EventMessageCreate         = "MESSAGE_CREATE"
	EventMessageReactionAdd    = "MESSAGE_REACTION_ADD"
	EventMessageReactionRemove = "MESSAGE_REACTION_REMOVE"
	EventAtMessageCreate       = "AT_MESSAGE_CREATE"
	EventPublicMessageDelete   = "PUBLIC_MESSAGE_DELETE"
	EventDirectMessageCreate   = "DIRECT_MESSAGE_CREATE"
	EventDirectMessageDelete   = "DIRECT_MESSAGE_DELETE"
	EventAudioStart            = "AUDIO_START"
	EventAudioFinish           = "AUDIO_FINISH"
	EventAudioOnMic            = "AUDIO_ON_MIC"
	EventAudioOffMic           = "AUDIO_OFF_MIC"
	EventMessageAuditPass      = "MESSAGE_AUDIT_PASS"
	EventMessageAuditReject    = "MESSAGE_AUDIT_REJECT"
	EventMessageDelete         = "MESSAGE_DELETE"
	EventForumThreadCreate     = "FORUM_THREAD_CREATE"
	EventForumThreadUpdate     = "FORUM_THREAD_UPDATE"
	EventForumThreadDelete     = "FORUM_THREAD_DELETE"
	EventForumPostCreate       = "FORUM_POST_CREATE"
	EventForumPostDelete       = "FORUM_POST_DELETE"
	EventForumReplyCreate      = "FORUM_REPLY_CREATE"
	EventForumReplyDelete      = "FORUM_REPLY_DELETE"
	EventForumAuditResult      = "FORUM_PUBLISH_AUDIT_RESULT"
	EventInteractionCreate     = "INTERACTION_CREATE"

	EventC2CMessageCreate     = "C2C_MESSAGE_CREATE"      // 用户单聊发消息给机器人时候
	EventGroupAtMessageCreate = "GROUP_AT_MESSAGE_CREATE" // 用户在群里@机器人时收到的消息
)
