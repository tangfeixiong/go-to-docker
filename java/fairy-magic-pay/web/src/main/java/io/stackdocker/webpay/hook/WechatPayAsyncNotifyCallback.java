package io.stackdocker.webpay.hook;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

@Component
public class WechatPayAsyncNotifyCallback {
    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    @RequestMapping(value = "/wechatpay/anh",
                    consumes = "application/xml", produces = "application/xml")
    public String reapNotification(@RequestParam(value="name", required=false, defaultValue="World") String name, Model model) {
        model.addAttribute("name", name);
        return "greeting";
    }

}
