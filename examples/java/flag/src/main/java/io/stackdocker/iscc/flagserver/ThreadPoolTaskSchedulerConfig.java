/*
  Inspired by: 
    - https://github.com/eugenp/tutorials/tree/master/spring-all/src/main/java/org/baeldung/taskscheduler
    - http://www.baeldung.com/spring-task-scheduler
*/

package io.stackdocker.iscc.flagserver;

import java.util.concurrent.TimeUnit;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.concurrent.ThreadPoolTaskScheduler;
import org.springframework.scheduling.support.CronTrigger;
import org.springframework.scheduling.support.PeriodicTrigger;

@Configuration
@EnableAsync
@EnableScheduling
// @Import(SwaggerConfig.class)
@ComponentScan( basePackages = {
    "io.stackdocker.iscc.flagserver.dispatcher",
    "cn.com.isc.config,cn.com.isc.controller",
    "cn.com.isc.entity",
    "cn.com.isc.filter",
    "cn.com.isc.repository,cn.com.isc.server"
    })
public class ThreadPoolTaskSchedulerConfig {
    @Value("${jobs.schedule}")
    private String cronExpr;

    @Bean
    public ThreadPoolTaskScheduler threadPoolTaskScheduler() {
        ThreadPoolTaskScheduler threadPoolTaskScheduler = new ThreadPoolTaskScheduler();
        threadPoolTaskScheduler.setPoolSize(5);
        threadPoolTaskScheduler.setThreadNamePrefix("ThreadPoolTaskScheduler");
        return threadPoolTaskScheduler;
    }

    @Bean
    public CronTrigger cronTrigger() {
        String env = System.getenv("JOBS_SCHEDULE");
        if ( env != null && env.trim() != "" ) 
            cronExpr = env;
        else if (cronExpr == null || cronExpr.trim() == "")
            cronExpr = "*/10 * * * * ?";
        return new CronTrigger(cronExpr);
    }

    @Bean
    public PeriodicTrigger periodicTrigger() {
        return new PeriodicTrigger(1000 * 60 * 10, TimeUnit.MICROSECONDS);
    }

    @Bean
    public PeriodicTrigger periodicFixedDelayTrigger() {
        PeriodicTrigger periodicTrigger = new PeriodicTrigger(1000 * 60 * 10, TimeUnit.MICROSECONDS);
        periodicTrigger.setFixedRate(true);
        periodicTrigger.setInitialDelay(1000 * 10);
        return periodicTrigger;
    }
}