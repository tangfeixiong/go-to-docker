
/*
  refer to
    http://www.mkyong.com/spring3/spring-aop-aspectj-annotation-example/
 */
package io.stackdocker.webpay.service;

import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.After;
import org.aspectj.lang.annotation.AfterReturning;
import org.aspectj.lang.annotation.AfterThrowing;
import org.aspectj.lang.annotation.Around;

import com.ijpay.controller.weixin.WxPayApiController;
//import com.jfinal.aop.Interceptor;
//import com.jfinal.aop.Invocation;
//import com.jfinal.core.Controller;
import com.jpay.weixin.api.WxPayApiConfigKit;
/**
 * Inspired by 
 * @Email javen205@126.com
 * @author Javen
 * https://github.com/Javen205/IJPay-Demo/blob/master/src/main/java/com/ijpay/interceptor/WxPayApiInterceptor.java
 */
@Aspect
public class WeChatPayApiInterceptor {
//public class WxPayApiInterceptor implements Interceptor {

//	@Override
//	public void intercept(Invocation inv) {
//		Controller controller = inv.getController();
//		if (controller instanceof WxPayApiController == false)
//			throw new RuntimeException("控制器需要继承 WxPayApiController");
		
//		try {
//			WxPayApiConfigKit.setThreadLocalWxPayApiConfig(((WxPayApiController)controller).getApiConfig());
//			inv.invoke();
//		}
//		finally {
//		}
//	}
   @AfterReturning(
      pointcut = "execution(* io.stackdocker.service.WeChatPayClient.getApiConfig(..))",
      returning= "result")
   public void logAfterReturning(JoinPoint joinPoint, Object result) {

	System.out.println("logAfterReturning() is running!");
	System.out.println("hijacked : " + joinPoint.getSignature().getName());
	System.out.println("Method returned value is : " + result);

   }
}