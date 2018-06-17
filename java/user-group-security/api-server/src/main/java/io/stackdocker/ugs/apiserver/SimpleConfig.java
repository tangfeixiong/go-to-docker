package io.stackdocker.ugs.apiserver;

import javax.servlet.Filter;

import io.stackdocker.ugs.apiserver.security.simple.BasicAuthFilter;
import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.ArrayList;
import java.util.List;

@Configuration
public class SimpleConfig {

    @Bean
    public FilterRegistrationBean<BasicAuthFilter> theBasicAuthFilter(){
        FilterRegistrationBean<BasicAuthFilter> registrationBean
                = new FilterRegistrationBean<>();

        registrationBean.setFilter(new BasicAuthFilter());
//        List<String> ups = new ArrayList<>();
//        ups.add("/v1/*");
//        registrationBean.setUrlPatterns(ups);
        registrationBean.addUrlPatterns("/v1/default/*");
        registrationBean.addInitParameter("paramName", "paramValue");

        return registrationBean;
    }

//    @Bean
//    public FilterRegistrationBean filterRegistration() {
//
//        FilterRegistrationBean registration = new FilterRegistrationBean();
//        registration.setFilter(basicauthFilter());
//        registration.addUrlPatterns("/v1/*");
//        registration.addInitParameter("paramName", "paramValue");
//        registration.setName("simpleBasicAuthFilter");
//        registration.setOrder(1);
//        return registration;
//    }
//
//    public Filter basicauthFilter() {
//        return new BasicAuthFilter();
//    }
}
