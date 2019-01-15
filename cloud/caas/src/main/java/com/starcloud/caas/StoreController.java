package com.starcloud.caas;

import java.io.IOException;

import com.alibaba.fastjson.JSON;
import com.starcloud.hbase.FamilyManager;
import com.starcloud.model.Alarm;
import com.starcloud.model.Family;
import com.starcloud.model.RealtimeData;
import com.starcloud.model.Result;

import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@EnableAutoConfiguration
public class StoreController {

    @RequestMapping(value = "/api/transport/realtimedata", method = RequestMethod.POST)
    @ResponseBody
    public String pub_realtime_data(@RequestBody String req) {
        try {
            RealtimeData rd = JSON.parseObject(req, RealtimeData.class);
            FamilyManager.putRealtimeData(rd);
            return "success";
        } catch (Exception ex) {
            return ex.getMessage();
        }

    }

    @RequestMapping(value = "/api/transport/result", method = RequestMethod.POST)
    @ResponseBody
    public String pub_result_data(@RequestBody String req) {
        try {
            Result r = JSON.parseObject(req, Result.class);
            FamilyManager.putResult(r);
            return "success";
        } catch (Exception ex) {
            return ex.getMessage();
        }

    }

    @RequestMapping(value = "/api/transport/alarm", method = RequestMethod.POST)
    @ResponseBody
    public String pub_alarm_data(@RequestBody String req) {
        try {
            Alarm a = JSON.parseObject(req, Alarm.class);
            FamilyManager.putAlarm(a);
            return "success";
        } catch (Exception ex) {
            return ex.getMessage();
        }

    }

    @RequestMapping(value = "/api/transport/count")
    public String count() {
        long r = 0L;
        try {
            r = FamilyManager.RowCountWithScanAndFilter();
        } catch (IOException e) {
            e.printStackTrace();
		}
        return String.valueOf(r);
    }

    @RequestMapping(value = "/api/transport/get_family")
    @ResponseBody
    public Family getFamily(@RequestParam String familyId) {
        try {
            Family f = FamilyManager.getFamily(familyId);
            return f;
        } catch (Exception ex) {
            return null;
        }
    }

    // @RequestMapping(value = "/api/list", method = RequestMethod.POST)
    // public String list(@RequestParam int start, @RequestParam int end) {

    //     return "success";
    // }

    // /**
    // * 创建日期:2018年4月6日<br/>
    // * 代码创建:黄聪<br/>
    // * 功能描述:通过HttpServletRequest 的方式来获取到json的数据<br/>
    // *
    // * @param request
    // * @return
    // */
    // @ResponseBody
    // @RequestMapping(value = "/request/data", method = RequestMethod.POST,
    // produces = "application/json;charset=UTF-8")
    // public String getByRequest(HttpServletRequest request) {
    // // 获取到JSONObject
    // JSONObject jsonParam = this.getJSONParam(request);
    // // 将获取的json数据封装一层，然后在给返回
    // JSONObject result = new JSONObject();
    // result.put("msg", "ok");
    // result.put("method", "request");
    // result.put("data", jsonParam);
    // return result.toJSONString();
    // }

    // /**
    // * 创建日期:2018年4月6日<br/>
    // * 代码创建:黄聪<br/>
    // * 功能描述:通过request来获取到json数据<br/>
    // *
    // * @param request
    // * @return
    // */
    // public JSONObject getJSONParam(HttpServletRequest request) {
    // JSONObject jsonParam = null;
    // try { // 获取输入流
    // BufferedReader streamReader = new BufferedReader(new
    // InputStreamReader(request.getInputStream(), "UTF-8"));
    // // 写入数据到Stringbuilder
    // StringBuilder sb = new StringBuilder();
    // String line = null;
    // while ((line = streamReader.readLine()) != null) {
    // sb.append(line);
    // }
    // jsonParam = JSONObject.parseObject(sb.toString());
    // // 直接将json信息打印出来
    // System.out.println(jsonParam.toJSONString());
    // } catch (Exception e) {
    // e.printStackTrace();
    // }
    // return jsonParam;
    // }

}