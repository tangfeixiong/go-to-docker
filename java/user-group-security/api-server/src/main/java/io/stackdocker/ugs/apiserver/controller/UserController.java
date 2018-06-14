package io.stackdocker.ugs.apiserver.controller;

import java.util.List;

import net.sf.json.JSONObject;
import org.apache.shiro.SecurityUtils;
import org.apache.shiro.Subject;
import org.apache.shiro.authc.AuthenticationException;
import org.apache.shiro.authc.LockedAccountException;
import org.apache.shiro.authc.UsernamePasswordToken;
import org.apache.shiro.authz.AuthorizationException;
import org.apache.shiro.authz.UnauthenticatedException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import io.stackdocker.ugs.apiserver.service.UserService;
import io.fairymagic.exam.api.dao.User;
import io.fairymagic.exam.api.dao.UserJsonTemplate;

@Controller
public class UserController {
    private static Logger logger = LoggerFactory.getLogger(UserController.class);
	
	@Autowired
	private UserService userService;
	
	
	@RequestMapping("/v1/test01/{id}")
	@ResponseBody
	public List<User> test01(@PathVariable("id") String id){
		
		List<User> userList = userService.findAllUser();
		System.out.println(id);
		
		return userList;
	}
	
	@RequestMapping(value = "/v1/default/users", method = POST,
            headers = "Accept=application/json", produces = "application/json")
	@ResponseBody
	public ResponseEntity<?> addOne(@RequestBody User user) {
		UsernamePasswordToken token = new UsernamePasswordToken(user.getUsername(), user.getPassword());
        Subject subject = SecurityUtils.getSubject();
        User saved;
        try {
            saved = userService.registerUser(user);
        } catch (Exception e) {
		    logger.Warn("Failed to register: " + e.getMessage())
            return new ResponseEntity<String>("Bad Request", HttpStatus.BAD_REQUEST);
        }
		
		return new ResponseEntity<Integer>(Integer.valueOf(user.getId()), HttpStatus.OK);
	}
    
	@RequestMapping(value = "/v1/default/credentials", method = POST,
            headers = "Accept=application/json", produces = "application/json")
	@ResponseBody
	public ResponseEntity<?> signIn(@RequestBody User user) {
		UsernamePasswordToken token = new UsernamePasswordToken(user.getUsername(), user.getPassword());
        Subject subject = SecurityUtils.getSubject();
        User saved;
        try {
            subject.login(token);
            saved = (User) subject.getPrincipal();
        } catch (LockedAccountException e) {
            logger.Warn("Account is locked: " + e.getMessage())
            token.clear();
            throw new AuthorizationException(e.getMessage());
            // return new ResponseEntity<String>("Forbidden", HttpStatus.FORBIDDEN);
        } catch (AuthenticationException e) {
		    logger.Warn("Failed to authenticate: " + e.getMessage())
            throw new UnauthenticatedException(e.getMessage());
            // return new ResponseEntity<String>("Unauthorized", HttpStatus.UNAUTHORIZED);
        }
		
		return new ResponseEntity<User>(saved, HttpStatus.OK);	
    }
    
	/**
	 * 判断用户名是否存在
	 * @param userName
	 * @return true==用户名已存在       false==用户名不存在
	 */
	@RequestMapping("/v1/ctf/user/findByName/{userName}")
	@ResponseBody
	public boolean findByName(@PathVariable("userName") String userName){
		User user = userService.findByName(userName);
		if(user!=null){
			return true;
		}else{
			return false;
		}
	}
	
	/**
	 * 判断手机号是否注册
	 * @param userMobile
	 * @return true==手机号已被注册   false==未被注册
	 */
	@RequestMapping("/v1/ctf/user/findByMobile/{userMobile}")
	@ResponseBody
	public boolean findByMobile(@PathVariable("userMobile")String userMobile){
		User user=userService.findByMobile(userMobile);
		if(user!=null){
			return true;
		}else{
			return false;
		}
	}
	
	/**
	 * 注册新用户
	 * @param user
	 * @return true==注册成功  false==注册失败
	 */
	@RequestMapping("/v1/ctf/user/register/{user}")
	@ResponseBody
	public boolean register(@PathVariable("user")String user){
		JSONObject jsonObject = JSONObject.fromObject(user);
		User registerUser = (User) JSONObject.toBean(jsonObject,User.class);
		int num = userService.registerUser(registerUser);
		if(num==1){
			return true;
		}else{
			return false;
		}
	}
	
	/**
	 * 用户登录验证
	 * @param user
	 * @return
	 */
	@RequestMapping("/v1/ctf/user/login/{user}")
	@ResponseBody
	public UserJsonTemplate login(@PathVariable("user")String user){
		JSONObject jsonObject = JSONObject.fromObject(user);
		User loginUser = (User) JSONObject.toBean(jsonObject,User.class);
		String userName = loginUser.getUsername();
		User u = userService.findByName(userName);
		String loginResult = "loginResult";
		if(u != null ){
			if(u.getPassword().equals(loginUser.getPassword())){
				return new UserJsonTemplate(u.getId(), u.getUsername(), loginResult,true);
			}else{
				return new UserJsonTemplate(null, null, loginResult,false);
			}
		}else{
			return new UserJsonTemplate( null, null, loginResult, false);
		}
	}
	
	/**
	 * 修改用户信息
	 * @param user
	 * @return true==修改成功  false==修改失败
	 */
	@RequestMapping("/v1/ctf/user/updateUser/{user}")
	@ResponseBody
	public boolean updateUser(@PathVariable("user")String user){
		JSONObject jsonObject=JSONObject.fromObject(user);
		User updateUser=(User) JSONObject.toBean(jsonObject,User.class);
		int num=userService.updateUser(updateUser);
		if(num==1){
			return true;
		}else{
			return false;
		}
	}
	
	/**
	 * 删除用户(逻辑删除，修改用户状态)
	 * @param id
	 * @return true==修改成功  false==修改失败
	 */
	@RequestMapping("/v1/ctf/user/deleteUser/{id}")
	@ResponseBody
	public boolean deleteUser(@PathVariable("id")String id){
		int num=userService.deleteUser(Integer.valueOf(id));
		if(num==1){
			return true;
		}else{
			return false;
		}
	}
	
	/**
	 * 通过id查找team的成员列表
	 * @param id
	 * @return
	 */
	@RequestMapping("/v1/ctf/team/findUsersByTeamId/{id}")
	@ResponseBody
	public List<User> findUsersByTeamId(@PathVariable("id")String id){
		return userService.findUsersByTeamId(Integer.valueOf(id));
	}
	
	/**
	 * 通过名字查找team成员列表
	 * @param name
	 * @return
	 */
	@RequestMapping("/v1/ctf/team/findUsersByTeamName/{name}")
	@ResponseBody
	public List<User> findUsersByTeamName(@PathVariable("name")String name){
		return userService.findUsersByTeamName(name);
	}
		
	
	
}
