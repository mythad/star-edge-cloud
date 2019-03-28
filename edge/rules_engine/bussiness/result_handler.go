package bussiness

// // ResultHandler 接收到数据，调用回调方法
// type ResultHandler struct {
// 	Client interfaces.IClient
// }

// // Handle - 如果是10的倍数数据，虚拟为报警数据
// // TIME_WAIT - 可能原因是传输太快，数据存储或发送未完成，留待观察
// // 所以此处尽量减少连接，数据压缩后再传输?
// // 修改此问题，也可以通过数据缓存？
// func (rh *ResultHandler) Handle(result *models.Result, response *models.Response) error {
// 	dir := common.GetCurrentDirectory()
// 	// 加载规则配置
// 	rules := &rms.Rules{}
// 	data := common.ReadString(dir + "/rules.xml")
// 	err := xml.Unmarshal([]byte(data), rules)
// 	if err == nil {
// 		log.Println(err.Error())
// 		os.Exit(1)
// 	}

// 	for _, rule := range rules.RuleItems {
// 		//rawYQL := `name='deen' and age>=23 and (hobby in ('soccer', 'swim') or score>90))`
// 		// result, _ := yql.Match(rawYQL, map[string]interface{}{
// 		// 	"name":  "deen",
// 		// 	"age":   int64(23),
// 		// 	"hobby": "basketball",
// 		// 	"score": int64(100),
// 		// })
// 		rh.handle(rule, result)
// 	}

// 	// fmt.Println(result)
// 	// rawYQL = `score ∩ (7,1,9,5,3)`
// 	// result, _ = yql.Match(rawYQL, map[string]interface{}{
// 	// 	"score": []int64{3, 100, 200},
// 	// })
// 	// fmt.Println(result)
// 	// rawYQL = `score in (7,1,9,5,3)`
// 	// result, _ = yql.Match(rawYQL, map[string]interface{}{
// 	// 	"score": []int64{3, 5, 2},
// 	// })
// 	// fmt.Println(result)
// 	// rawYQL = `score.sum() > 10`
// 	// result, _ = yql.Match(rawYQL, map[string]interface{}{
// 	// 	"score": []int{1, 2, 3, 4, 5},
// 	// })
// 	// fmt.Println(result)

// 	// i, _ := common.StringToInt(string(request.Data[:]))
// 	// // extlog.WriteLog("收到实时数据:" + request.ID)
// 	// if i%10 == 0 {
// 	// 	response = &models.Response{Status: "true"}
// 	// 	alarm := &models.Alarm{}
// 	// 	alarm.ID = common.GetGUID()
// 	// 	alarm.Data = request.Data
// 	// 	// for e := rh.a1.Listeners.Front(); e != nil; e = e.Next() {
// 	// 	// 	rh.client.PostAlarm(e.Value.(string)+"/api/transport/alarm", alarm)
// 	// 	// 	extlog.WriteLog("发送报警数据:" + e.Value.(string) + "/api/transport/alarm" + alarm.ID)
// 	// 	// }
// 	// 	return nil
// 	// }

// 	// response = &models.Response{Status: "false"}
// 	return nil
// }

// func (rh *ResultHandler) handle(r rms.Rule, data *models.Result) {
// 	result, _ := yql.Match(r.Yql, map[string]interface{}{})
// 	if result {
// 		_request := &models.Request{}
// 		_request.ID = common.GetGUID()
// 		_request.Message = data.Data
// 		// for e := rh.a1.Listeners.Front(); e != nil; e = e.Next() {
// 		rh.Client.PostRequest(r.PostAddr, _request)
// 		// 	extlog.WriteLog("发送报警数据:" + e.Value.(string) + "/api/transport/alarm" + alarm.ID)
// 		// }
// 	}
// 	if r.ChildRule != nil && len(r.ChildRule) > 0 {
// 		for _, cr := range r.ChildRule {
// 			rh.handle(cr, data)
// 		}
// 	}
// }
