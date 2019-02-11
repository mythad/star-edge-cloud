package com.starcloud.extension;

import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@EnableAutoConfiguration
public class HomeController {

    @RequestMapping(value = "/data", method = RequestMethod.POST)
    public String data(@RequestParam String req) {

        return "success";
    }

    @RequestMapping(value = "/list", method = RequestMethod.POST)
    public String list(@RequestParam int start, @RequestParam int end) {

        return "success";
    }

    @RequestMapping(value = "/count")
    public String count() {

        return "success";
    }
}