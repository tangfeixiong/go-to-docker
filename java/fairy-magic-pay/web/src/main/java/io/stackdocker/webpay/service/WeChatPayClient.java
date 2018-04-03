package io.stackdocker.webpay.service;

import java.util.Map;
import java.util.UUID;
import java.util.concurrent.atomic.AtomicInteger;

import com.jpay.ext.kit.IpKit;
import com.jpay.ext.kit.PaymentKit;
import com.jpay.ext.kit.ZxingKit;
import com.jpay.secure.RSAUtils;
//import com.jpay.vo.AjaxResult;
import com.jpay.weixin.api.WxPayApiConfig;
import com.jpay.weixin.api.WxPayApiConfigKit;
import com.jpay.weixin.api.WxPayApiConfig.PayModel;
import com.jpay.weixin.api.WxPayApi;
import com.jpay.weixin.api.WxPayApi.TradeType;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import io.stackdocker.webpay.api.Payment;
import io.stackdocker.webpay.help.WeChatUnifiedOrderApiException;

//@Service
public class WeChatPayClient {
    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    private final AtomicInteger counter = new AtomicInteger();
    
    @Autowired Map<UUID, Payment> cachedPayments; 
    
	// tecent open platform
    //   https://pay.weixin.qq.com/wechatpay_guide/help_docs.shtml
    
    // application identity
    @Value("${wechat.pay.app_id}") private String appId;
	
    // vender id
    @Value("${wechat.pay.mch_id") private String mchId;
    
    @Value("${wxpay.partner_key}") private String partnerKey;
    
    @Value("${wxpay.cert_path") private String certPath;
    
    @Value("${wxpay.notify_url") private String notifyUrl;

    public WeChatPayClient() {
        logger.info("props=" + appId + ";" + mchId + ";" + partnerKey + ";" + certPath + ";" + notifyUrl);
    }

    public String getAppId() {
    	return appId;
	}

	public String getMchId() {
    	return mchId;
	}

	public String getPartnerKey() {
    	return partnerKey;
	}

	public String getCertPath() {
    	return certPath;
	}

	public String getNotifyUrl() {
    	return notifyUrl;
	}

	
	public WxPayApiConfig getApiConfig() {
		return WxPayApiConfig.New()
				.setAppId(appId)
				.setMchId(mchId)
				.setPaternerKey(partnerKey)
				.setPayModel(PayModel.BUSINESSMODEL);
	}

	
	/**
     * refer
     *   WeChat Payment API(OverseasServiceProviderModeV1.3.1.2).pdf
     *
     * inspired by 
     *   https://github.com/Javen205/IJPay-Demo/blob/master/src/main/java/com/ijpay/controller/weixin/WxPayController.java#scanCode2
	 * 扫码支付模式二
	 * 已测试
     * https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=6_4
     * https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_1
	 */
	public Map<String, String> prePay(UUID billing, int booking, String attachedArg, String chargeSub, String ip) throws WeChatUnifiedOrderApiException {
        logger.info("Naitve payment via mode 2");
		
		String openId="oUpF8uMuAJO_M2pxb1Q9zNjWeS6o";
//		String openId = (String) getSession().getAttribute("openId");
				
//		if (StrKit.isBlank(openId)) {
//			ajax.addError("openId is null");
//			renderJson(ajax);
//			return;
//		}
		
        // if not cached billing
        int count = counter.incrementAndGet();
        String prePayId = Integer.toString(count);
        
		Map<String, String> params = WxPayApiConfigKit.getWxPayApiConfig()
				.setAttach(attachedArg)
				.setBody(chargeSub)
				.setOpenId(openId)
				.setSpbillCreateIp(ip)
				.setTotalFee(Integer.toString(booking))
				.setTradeType(TradeType.NATIVE)
				.setNotifyUrl(notifyUrl)
				.setOutTradeNo(/*String.valueOf(System.currentTimeMillis())*/prePayId)
				.build();
		
		String xmlResult = WxPayApi.pushOrder(false, params);
		logger.debug(xmlResult);
        
        // Step 1: Prepay
		Map<String, String> result = PaymentKit.xmlToMap(xmlResult);
		
		String return_code = result.get("return_code");
		String return_msg = result.get("return_msg");
		if (!PaymentKit.codeIsOK(return_code)) {
//			System.out.println(return_msg);
//			renderText(xmlResult);
//			return;
            throw new WeChatUnifiedOrderApiException(return_msg);
		}
		String result_code = result.get("result_code");
		if (!PaymentKit.codeIsOK(result_code)) {
//			System.out.println(return_msg);
//			renderText(xmlResult);
//			return;
            throw new WeChatUnifiedOrderApiException(return_msg);
		}
//		//生成预付订单success
		
		String qrCodeUrl = result.get("code_url");
//		String name = "payQRCode2.png";
		
//		Boolean encode = ZxingKit.encode(qrCodeUrl, BarcodeFormat.QR_CODE, 3, ErrorCorrectionLevel.H, "png", 200, 200,
//				PathKit.getWebRootPath()+File.separator+name);
//		if (encode) {
////			renderQrCode(qrCodeUrl, 200, 200);
//			//在页面上显示
//			ajax.success(name);
//			renderJson(ajax);
//		}

       cachedPayments.put(billing, new Payment((long)count, billing, booking, qrCodeUrl));
       return result;
	}
    
}