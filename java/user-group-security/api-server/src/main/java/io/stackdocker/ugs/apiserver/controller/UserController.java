package io.stackdocker.ugs.apiserver.controller;

import io.fairymagic.core.ugs.domain.User;
import io.stackdocker.ugs.apiserver.dao.UserDao;
import io.stackdocker.ugs.apiserver.service.UserService;

import org.apache.shiro.SecurityUtils;
import org.apache.shiro.authc.AuthenticationException;
import org.apache.shiro.authc.LockedAccountException;
import org.apache.shiro.authc.UsernamePasswordToken;
import org.apache.shiro.authz.AuthorizationException;
import org.apache.shiro.authz.UnauthenticatedException;
import org.apache.shiro.subject.PrincipalCollection;
import org.apache.shiro.subject.Subject;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
public class UserController {
    private static Logger logger = LoggerFactory.getLogger(UserController.class);

    @Autowired
    UserService userService;

    @RequestMapping(value = "/v1/default", method = RequestMethod.GET)
    public ResponseEntity<Object> doSomething(@RequestHeader(name = "remote_addr")
                                                      String remoteAddress) {
        logger.debug("The Remote address added by WebFiler is :: {}", remoteAddress);
        ResponseEntity<Object> response = null;
        try {
            response = new ResponseEntity<Object>("SUCCESS", HttpStatus.OK);
        } catch (Exception ex) {
            logger.error(ex.getMessage(), ex);
            return new ResponseEntity<Object>(ex.getMessage(),
                    HttpStatus.INTERNAL_SERVER_ERROR);
        }
        return response;
    }

    @RequestMapping(value = "/v1/default/users", method = RequestMethod.POST,
            headers = "Accept=application/json", produces = "application/json")
    @ResponseBody
    public ResponseEntity<?> addOne(@RequestHeader(name = "shiro_security") String ssFlag, @RequestBody User user) {
        if (false == "disabled".equals(ssFlag.toLowerCase())) {
            UsernamePasswordToken token = new UsernamePasswordToken(user.getUsername(), user.getPassword());
            Subject currentUser = SecurityUtils.getSubject();
            PrincipalCollection principals = currentUser.getPrincipals();
            if (principals != null && !principals.isEmpty()) {
                throw new AuthorizationException("User existed: " + user.getUsername());
            }
        }

        Long id;
        try {
            id = userService.addOne(user);
        } catch (Exception e) {
            logger.warn("Failed to register: " + e.getMessage());
            return new ResponseEntity<String>("Bad Request", HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<Long>(user.getId(), HttpStatus.OK);
    }

    @RequestMapping(value = "/v1/default/credentials", method = RequestMethod.POST,
            headers = "Accept=application/json", produces = "application/json")
    @ResponseBody
    public ResponseEntity<?> verifyOne(@RequestBody User user) {
        UsernamePasswordToken token = new UsernamePasswordToken(user.getUsername(), user.getPassword());
        Subject subject = SecurityUtils.getSubject();
        User saved;
        try {
            subject.login(token);
            saved = (User) subject.getPrincipal();
        } catch (LockedAccountException e) {
            logger.warn("Account is locked: " + e.getMessage());
            token.clear();
            throw new AuthorizationException(e.getMessage());
            // return new ResponseEntity<String>("Forbidden", HttpStatus.FORBIDDEN);
        } catch (AuthenticationException e) {
            logger.warn("Failed to authenticate: " + e.getMessage());
            throw new UnauthenticatedException(e.getMessage());
            // return new ResponseEntity<String>("Unauthorized", HttpStatus.UNAUTHORIZED);
        }

        return new ResponseEntity<User>(saved, HttpStatus.OK);
    }

}
