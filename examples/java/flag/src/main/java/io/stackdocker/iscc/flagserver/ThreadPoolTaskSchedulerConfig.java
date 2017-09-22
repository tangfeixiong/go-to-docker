/*
  Inspired by: 
    - https://github.com/eugenp/tutorials/tree/master/spring-all/src/main/java/org/baeldung/taskscheduler
    - http://www.baeldung.com/spring-task-scheduler
    - https://github.com/johnlabarge/spring-boot-redis-k8s/tree/master/src/main/java/com/google/cloudjlb/Application.java
    - https://dzone.com/articles/integrate-redis-to-your-spring-project
    - http://www.baeldung.com/spring-cache-tutorial
    - https://dzone.com/articles/enabling-caching-in-mongodb-database-with-redis-us
    - http://caseyscarborough.com/blog/2014/12/18/caching-data-in-spring-using-redis/
*/

package io.stackdocker.iscc.flagserver;

import java.util.concurrent.TimeUnit;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cache.CacheManager;
import org.springframework.cache.annotation.CachingConfigurerSupport;
import org.springframework.cache.annotation.EnableCaching;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.core.env.Environment;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;
import org.springframework.data.redis.cache.RedisCacheManager;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.connection.jedis.JedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.RedisSerializer;
import org.springframework.data.redis.serializer.StringRedisSerializer;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.concurrent.ThreadPoolTaskScheduler;
import org.springframework.scheduling.support.CronTrigger;
import org.springframework.scheduling.support.PeriodicTrigger;

@Configuration
@EnableAsync
@EnableScheduling
@EnableCaching
@EnableJpaRepositories (basePackages = {
    "io.stackdocker.iscc.flagserver.cache",
    "cn.com.isc.repository"
})
// @Import(SwaggerConfig.class)
@ComponentScan( basePackages = {
    "io.stackdocker.iscc.flagserver.api",
    "io.stackdocker.iscc.flagserver.cache",
    "io.stackdocker.iscc.flagserver.controller",
    "io.stackdocker.iscc.flagserver.dispatcher",
    "io.stackdocker.iscc.flagserver.domain",
    "cn.com.isc.config",
    "cn.com.isc.controller",
    "cn.com.isc.entity",
    "cn.com.isc.filter",
    "cn.com.isc.repository",
    "cn.com.isc.server"
})
public class ThreadPoolTaskSchedulerConfig extends CachingConfigurerSupport {
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


    @Autowired
    private Environment env;

    @Bean
    public JedisConnectionFactory redisConnectionFactory() {
	    String redisHost = env.getProperty("spring.redis.host");
        int redisPort = 6379;
        JedisConnectionFactory jedisConnectionFactory = new JedisConnectionFactory();
        jedisConnectionFactory.setHostName(redisHost);
        jedisConnectionFactory.setPort(redisPort);
        jedisConnectionFactory.setUsePool(true);
        return jedisConnectionFactory;
    }
    @Bean
    public RedisTemplate<Object, Object> redisTemplate()
    {
        RedisTemplate<Object, Object> redisTemplate = new RedisTemplate<Object, Object>();
        redisTemplate.setConnectionFactory(redisConnectionFactory());
        redisTemplate.setExposeConnection(true);
        return redisTemplate;
    }
    @Bean
    public RedisCacheManager cacheManager()
    {
        RedisCacheManager redisCacheManager = new RedisCacheManager(redisTemplate());
        redisCacheManager.setTransactionAware(true);
        redisCacheManager.setLoadRemoteCachesOnStartup(false);
        redisCacheManager.setUsePrefix(true);
        return redisCacheManager;
    }
    @Bean
    public RedisSerializer redisStringSerializer() {
        StringRedisSerializer stringRedisSerializer = new StringRedisSerializer();
        return stringRedisSerializer;
    }
//    @Bean(name="redisTemplate")
//    public RedisTemplate<String, String> redisTemplate(RedisConnectionFactory cf,RedisSerializer redisSerializer) {
//        RedisTemplate<String, String> redisTemplate = new RedisTemplate<String, String>();
//        redisTemplate.setConnectionFactory(cf);
//        redisTemplate.setDefaultSerializer(redisSerializer);
//        return redisTemplate;
//    }
//    @Bean
//    public CacheManager cacheManager() {
//        RedisCacheManager cacheManager = new RedisCacheManager(redisTemplate(redisConnectionFactory(),redisStringSerializer()));
//        cacheManager.setDefaultExpiration(3000);
//        return cacheManager;
//    }

}