package mhw

import "time"

// if command.Contains(do.RawContent(), command.Hunter) {
// 	assembly := strings.TrimSpace(do.Content())
// 	if assembly == "" {
// 		if assemblyCode == "" {
// 			return bot.PaperAirplane.ToGroup(groupOpenId).Reply(msgId, "当前没有人在开趴，真是一群杂鱼❤~")
// 		}
// 		dur := time.Since(lastUpdateTime).String()
// 		return bot.PaperAirplane.ToGroup(groupOpenId).Reply(msgId, "[距离上次更新已过去 "+dur+"] 当前集会码为："+assemblyCode)
// 	}

// 	// 更新集会码
// 	if strings.Contains(assembly, " ") || len(assembly) != 12 {
// 		return bot.PaperAirplane.ToGroup(groupOpenId).Reply(msgId, "集会码非法")
// 	}
// 	assemblyCode = assembly
// 	lastUpdateTime = time.Now()
// 	return bot.PaperAirplane.ToGroup(groupOpenId).Reply(msgId, "["+lastUpdateTime.String()+"] 集会码已更新，当前集会码为："+assemblyCode)
// }

var lastUpdateTime = time.Now()
