package io.stackdocker.webpay.service;

//import com.ijpay.interceptor.WxPayApiInterceptor;
//import com.jfinal.aop.Before;
//import com.jfinal.core.Controller;
import com.jpay.weixin.api.WxPayApiConfig;
import org.springframework.stereotype.Service;

/**
 * Inspired by
 * @Email javen205@126.com
 * @author Javen
 * https://github.com/Javen205/IJPay-Demo/blob/master/src/main/java/com/ijpay/controller/weixin/WxPayApiController.java
 */
interface WeChatPayApi {
//@Before(WxPayApiInterceptor.class)
//public abstract class WxPayApiController extends Controller {
	WxPayApiConfig getApiConfig();

}
