package openai

import (
	"github.com/songquanpeng/one-api/relay/adaptor/ai360"
	"github.com/songquanpeng/one-api/relay/adaptor/baichuan"
	"github.com/songquanpeng/one-api/relay/adaptor/groq"
	"github.com/songquanpeng/one-api/relay/adaptor/lingyiwanwu"
	"github.com/songquanpeng/one-api/relay/adaptor/minimax"
	"github.com/songquanpeng/one-api/relay/adaptor/mistral"
	"github.com/songquanpeng/one-api/relay/adaptor/moonshot"
	"github.com/songquanpeng/one-api/relay/adaptor/stepfun"
	"github.com/songquanpeng/one-api/relay/channeltype"
)

var CompatibleChannels = []int{
	channeltype.Azure,
	channeltype.AI360,
	channeltype.Moonshot,
	channeltype.Baichuan,
	channeltype.Minimax,
	channeltype.Mistral,
	channeltype.Groq,
	channeltype.LingYiWanWu,
	channeltype.StepFun,
}

func GetCompatibleChannelMeta(channelType int) (string, []string) {
	switch channelType {
	case channeltype.Azure:
		return "azure", ModelList
	case channeltype.AI360:
		return "360", ai360.ModelList
	case channeltype.Moonshot:
		return "moonshot", moonshot.ModelList
	case channeltype.Baichuan:
		return "baichuan", baichuan.ModelList
	case channeltype.Minimax:
		return "minimax", minimax.ModelList
	case channeltype.Mistral:
		return "mistralai", mistral.ModelList
	case channeltype.Groq:
		return "groq", groq.ModelList
	case channeltype.LingYiWanWu:
		return "lingyiwanwu", lingyiwanwu.ModelList
	case channeltype.StepFun:
		return "stepfun", stepfun.ModelList
	default:
		return "openai", ModelList
	}
}
