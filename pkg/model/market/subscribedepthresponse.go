package market

import (
	"github.com/cnvic/huobi_golang/v3/pkg/model/base"
)

type SubscribeDepthResponse struct {
	base.WebSocketResponseBase
	Data *Depth
	Tick *Depth
}
