package io.stackdocker.ugs.webapp.mvc;

import net.sf.json.JSONObject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Profile;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

import javax.servlet.http.HttpSession;
import java.net.MalformedURLException;
import java.net.URISyntaxException;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.atomic.AtomicLong;

import com.blackbird.web.bean.UserJsonTemplate;
import io.fairymagic.core.ugs.domain.User;

@Controller
//@Profile("dev")
public class SimpleDevController {
    private static final Logger logger = LoggerFactory.getLogger(SimpleDevController.class);
    @Autowired
    Environment env;

    private static final AtomicLong counter = new AtomicLong();

    private static final Map<String, User> cache = new HashMap<>();

    @RequestMapping(value={"/exam", "/exam/", "/exam/index.jsp"})
    public String index() {
        return "index";
    }

    /**
     * singIn
     * */
    @RequestMapping(value="/exam/user/signin", method= RequestMethod.POST)
    public String signIn(User user, HttpSession session){
        logger.info("User " + user.getUsername() + " sign in");

        if ( user.getUsername().contentEquals("ok") ) {
            logger.info("succeeded");
            long id = counter.incrementAndGet();
            user.setId(id);
            cache.put(user.getUsername(), user);
            session.setAttribute("username", user.getUsername());
            session.setAttribute("userid", Long.valueOf(id));
            //TODO 登录成功后跳转的页面
            return "login";
        }else{
            logger.warn("Failed to sign in");
            return "regist";
        }
    }

}
