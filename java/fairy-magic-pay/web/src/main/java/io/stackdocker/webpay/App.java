package /* sample.web.staticcontent */ io.stackdocker.webpay;

import java.util.Map;
import java.util.HashMap;
import java.util.UUID;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.boot.web.servlet.support.SpringBootServletInitializer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;

import io.stackdocker.webpay.api.Payment;
import io.stackdocker.webpay.service.WeChatPayClient;

@SpringBootApplication
@ComponentScan("io.stackdocker.webpay")
public class App {
//public class App extends SpringBootServletInitializer {

	@Bean
    public Map<UUID, Payment> cachedPayments() {
        return new HashMap<UUID, Payment>();
    }

//	@Bean
//	public WeChatPayClient wxpayClient() {
//		return new WeChatPayClient();
//	}

//	@Bean
//	public AliPayService aliPayService() {
//		return new AliPayService();
//	}

//	@Bean
//	public UnionPayService unionPayService() {
//		return new UnionPayService();
//	}

	/*
	  Refer to
		https://docs.spring.io/spring-boot/docs/current/reference/html/howto-traditional-deployment.html
	*/
//	@Override
//	protected SpringApplicationBuilder configure(SpringApplicationBuilder application) {
//		return application.sources(App.class);
//	}

	public static void main(String[] args) {
		SpringApplication.run(App.class, args);
	}

}
