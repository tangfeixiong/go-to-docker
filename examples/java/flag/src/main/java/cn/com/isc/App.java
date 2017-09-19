package cn.com.isc;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.web.servlet.ServletComponentScan;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.scheduling.annotation.EnableScheduling;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

import cn.com.isc.config.IpFilterUpdateFlag;
import cn.com.isc.config.OpenFlag;
import cn.com.isc.config.TokenUpdate;

@SpringBootApplication
@ServletComponentScan
@EnableScheduling
//@EnableAutoConfiguration
@ComponentScan( basePackages = {
                  "cn.com.isc",
                  "io.stackdocker.iscc.flagserver"
                })
public class App {
	
	public static void main(String[] args) {
		SpringApplication.run(App.class, args);
	}
		
	@Bean
	public IpFilterUpdateFlag ipFilterUpdateFlag(){
		return new IpFilterUpdateFlag();
	}
	@Bean
	public OpenFlag openFlag(){
		return new OpenFlag();
	}

	@Bean
	public TokenUpdate tokenUpdate(){
		return new TokenUpdate();
	}
}
