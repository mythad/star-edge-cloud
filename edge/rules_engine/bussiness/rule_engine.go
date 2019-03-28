package bussiness

import (
	"encoding/xml"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/transport/client"
	"star-edge-cloud/edge/transport/client/configuration"
	"star-edge-cloud/edge/transport/coding/json"
	"star-edge-cloud/edge/utils/common"
	"time"

	"github.com/caibirdme/yql"
)

// YQLRulesEngine -
type YQLRulesEngine struct {
	client *client.Client
	rules  *models.Rules
}

// LoadRules -
func (it *YQLRulesEngine) LoadRules(data []byte) {
	// 加载规则配置
	xml.Unmarshal([]byte(data), it.rules)

	// todo:根据规则要求，创建指定协议客户端，发送处理结果
	if it.client == nil {
		it.client = client.New()
	}
}

// Handle -
func (it *YQLRulesEngine) Handle(v interface{}) {
	for _, r := range it.rules.RuleItems {
		data := make(map[string]interface{})
		data["realtime_data"] = v
		it.handle(&r, data)
	}
}

// handle -
func (it *YQLRulesEngine) handle(r *models.Rule, data map[string]interface{}) {
	c := it.client.GetChannel(&configuration.ChannelConfig{
		Timeout:  10 * time.Second,
		Keep:     true,
		DataType: constant.RealtimeData,
		Protocol: constant.HTTP,
		Cache:    false,
	})
	c.SetDecoder(&json.JSONDecoder{})
	c.SetEncoder(&json.JSONEncoder{})
	if result, _ := yql.Match(r.Yql, data); result {
		_request := &models.Request{
			ID: common.GetGUID(),
			// Content: data,
		}
		// for e := rh.a1.Listeners.Front(); e != nil; e = e.Next() {
		c.Send(r.PostAddr, _request, map[string]interface{}{"headerType": "text/plain"})
	}
	if r.ChildRule != nil && len(r.ChildRule) > 0 {
		for _, cr := range r.ChildRule {
			it.handle(&cr, data)
		}
	}
}
