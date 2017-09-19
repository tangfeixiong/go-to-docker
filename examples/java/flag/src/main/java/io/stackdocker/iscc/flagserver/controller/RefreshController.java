package io.stackdocker.iscc.flagserver.controller;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.MessageDigest;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import cn.com.isc.config.Scheduler;
import cn.com.isc.entity.Flag;
import cn.com.isc.server.FlagService;

import io.stackdocker.iscc.flagserver.api.RefreshReqResp;
import io.stackdocker.iscc.flagserver.dispatcher.RefreshScheduler;
import io.stackdocker.iscc.flagserver.domain.RefreshConfig;

@RestController
public class RefreshController {
	@Autowired
	private FlagService flagService;
	
	@Autowired
	private Scheduler scheduler;
    
    @Autowired
    private RefreshScheduler resch;
    
    public RefreshController() {
        
    }
	
	@GetMapping("/v1/open")
	public Map<String, Object> test(){
		flagService.open();
		
		Map<String, Object> result = new HashMap<String, Object>();
		result.put("time", new Date().getTime());
		result.put("status", "OK");
		scheduler.setCurrentTime(System.currentTimeMillis());
		
		return result;
	}

    // fanhonglingdeMacBook-Pro:go-to-authnz fanhongling$ curl http://172.17.4.50:8082/getMechineFlag -X POST -H "Content-Type: application/json" -d '{"env": 1, "teamNo": 1}'
    // {"flag":"3784AD50F3C6D375567CE31FC09F6D89"}
    // 比赛回合数，例如flag值30分钟一刷新，比赛时间为10个小时，那么回合数就为20
	@PostMapping("/v1/refresh-creation")
	public RefreshReqResp createRefresh( @RequestBody RefreshReqResp req ){
		System.out.println("Go to create refreshing: " + req);
        return resch.create(req);
	}
    
	@PostMapping("/v1/refresh-deletion")
	public RefreshReqResp deleteRefresh( @RequestBody RefreshReqResp req ){
		System.out.println("Go to delete refreshing: " + req);
        return resch.delete(req);
	}
    
	@PostMapping("/v1/refresh-updation")
	public RefreshReqResp updateRefresh( @RequestBody RefreshReqResp req ){
		System.out.println("Go to update refreshing: " + req);
        return resch.update(req);
	}
	
	//5）具备数据库查询接口功能，即：当靶机请求属于该靶机的Flag值时，Flag控制服务器端程序可去数据库中检索到，并response给请求发起方
	@PostMapping("/v1/Validate")
	public Map<String, String> get(@RequestBody Flag flag){
		Map<String, String> result = new HashMap<String, String>();	
		Flag value = flagService.getByFlag(flag.getMd5String());
		if(value != null && scheduler.getCount() == value.getRound()){
			result.put("Token", value.getToken());
		}else{
			result.put("Token", "");
		}
		return result;
	}
	
	//返回当前轮次
	@PostMapping("/v1/getCount")
	public Map<String, Object> get(){
		Map<String, Object> map = new HashMap<String, Object>();
		map.put("value", scheduler.getCount());
		long lCtime = System.currentTimeMillis();
		
		map.put("timestamp",  (lCtime - scheduler.getCurrentTime())/1000);
		return map;
	}
}

class Utils {
	public static String getMD5(String message) {  
        String md5 = "";  
        try {  
            MessageDigest md = MessageDigest.getInstance("MD5");  // 创建一个md5算法对象  
            byte[] messageByte = message.getBytes("UTF-8");  
            byte[] md5Byte = md.digest(messageByte);              // 获得MD5字节数组,16*8=128位  
            md5 = bytesToHex(md5Byte);                            // 转换为16进制字符串  
        } catch (Exception e) {
            e.printStackTrace();  
        }  
        return md5;  
    }  
	
	 // 二进制转十六进制  
    public static String bytesToHex(byte[] bytes) {  
        StringBuffer hexStr = new StringBuffer();  
        int num;  
        for (int i = 0; i < bytes.length; i++) {  
            num = bytes[i];  
             if(num < 0) {  
                 num += 256;  
            }  
            if(num < 16){  
                hexStr.append("0");  
            }  
            hexStr.append(Integer.toHexString(num));  
        }  
        return hexStr.toString().toUpperCase();  
    }  
}
