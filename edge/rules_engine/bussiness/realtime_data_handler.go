package bussiness

import (
	"star-edge-cloud/edge/rules_engine/interfaces"
	"star-edge-cloud/edge/transport/coding"
)

// RealtimeDataHandler 接收到数据，调用回调方法
type RealtimeDataHandler struct {
	decoder coding.IDecoder
	encoder coding.IEncoder
	engine  interfaces.IRulesEngine
}

// SetDecoder -
func (it *RealtimeDataHandler) SetDecoder(de coding.IDecoder) {
	it.decoder = de
}

// SetEncoder -
func (it *RealtimeDataHandler) SetEncoder(en coding.IEncoder) {
	it.encoder = en
}

// SetRulesEngine -
func (it *RealtimeDataHandler) SetRulesEngine(engine interfaces.IRulesEngine) {
	it.engine = engine
}

// OnReceive -
func (it *RealtimeDataHandler) OnReceive(v interface{}) ([]byte, error) {
	it.engine.Handle(v)
	// fmt.Println(result)
	// rawYQL = `score ∩ (7,1,9,5,3)`
	// result, _ = yql.Match(rawYQL, map[string]interface{}{
	// 	"score": []int64{3, 100, 200},
	// })
	// fmt.Println(result)
	// rawYQL = `score in (7,1,9,5,3)`
	// result, _ = yql.Match(rawYQL, map[string]interface{}{
	// 	"score": []int64{3, 5, 2},
	// })
	// fmt.Println(result)
	// rawYQL = `score.sum() > 10`
	// result, _ = yql.Match(rawYQL, map[string]interface{}{
	// 	"score": []int{1, 2, 3, 4, 5},
	// })
	// fmt.Println(result)

	// i, _ := common.StringToInt(string(request.Data[:]))
	// // extlog.WriteLog("收到实时数据:" + request.ID)
	// if i%10 == 0 {
	// 	response = &models.Response{Status: "true"}
	// 	alarm := &models.Alarm{}
	// 	alarm.ID = common.GetGUID()
	// 	alarm.Data = request.Data
	// 	// for e := rh.a1.Listeners.Front(); e != nil; e = e.Next() {
	// 	// 	rh.client.PostAlarm(e.Value.(string)+"/api/transport/alarm", alarm)
	// 	// 	extlog.WriteLog("发送报警数据:" + e.Value.(string) + "/api/transport/alarm" + alarm.ID)
	// 	// }
	// 	return nil
	// }

	// response = &models.Response{Status: "false"}
	return nil, nil
}

// func (it *RealtimeDataHandler) handle(r *models.Rule, data *models.RealtimeData) {
// 	if it.client == nil {
// 		it.client = client.New()
// 	}
// 	c := it.client.GetChannel(&configuration.ChannelConfig{
// 		Timeout:  10 * time.Second,
// 		Keep:     true,
// 		DataType: constant.RealtimeData,
// 		Protocol: constant.HTTP,
// 		Cache:    false,
// 	})
// 	c.SetDecoder(&json.JSONDecoder{})
// 	c.SetEncoder(&json.JSONEncoder{})
// 	result, _ := yql.Match(r.Yql, map[string]interface{}{})
// 	if result {
// 		_request := &models.Request{
// 			ID:      common.GetGUID(),
// 			Content: data.Data,
// 		}
// 		// for e := rh.a1.Listeners.Front(); e != nil; e = e.Next() {
// 		c.Send(r.PostAddr, _request)
// 		// 	extlog.WriteLog("发送报警数据:" + e.Value.(string) + "/api/transport/alarm" + alarm.ID)
// 		// }
// 	}
// 	if r.ChildRule != nil && len(r.ChildRule) > 0 {
// 		for _, cr := range r.ChildRule {
// 			it.handle(&cr, data)
// 		}
// 	}
// }
