package io.stackdocker.ugs.apiserver;

import javax.sql.DataSource;

import at.pollux.thymeleaf.shiro.dialect.ShiroDialect;
import io.stackdocker.ugs.apiserver.model.ErrMsg;
import io.stackdocker.ugs.apiserver.model.NotFoundException;
import org.apache.shiro.authc.credential.CredentialsMatcher;
import org.apache.shiro.authc.credential.HashedCredentialsMatcher;
import org.apache.shiro.authz.AuthorizationException;
import org.apache.shiro.authz.UnauthenticatedException;
import org.apache.shiro.cache.CacheManager;
import org.apache.shiro.cache.MemoryConstrainedCacheManager;
import org.apache.shiro.crypto.hash.Sha256Hash;
import org.apache.shiro.realm.jdbc.JdbcRealm;
import org.apache.shiro.spring.LifecycleBeanPostProcessor;
import org.apache.shiro.spring.web.config.DefaultShiroFilterChainDefinition;
import org.apache.shiro.spring.web.config.ShiroFilterChainDefinition;
import org.apache.shiro.web.mgt.DefaultWebSecurityManager;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.DependsOn;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ControllerAdvice;

import io.stackdocker.ugs.apiserver.security.ShiroAuthczRealm;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;

@ControllerAdvice
@Configuration
public class ShiroSecureConf {
    private static Logger logger = LoggerFactory.getLogger(App.class);

//    @Autowired(required = false)
//    private ResourcesService resourcesService;

    @ExceptionHandler(UnauthenticatedException.class)
    @ResponseStatus(HttpStatus.UNAUTHORIZED)
    public void handleException(UnauthenticatedException e) {
        logger.debug("{} was thrown", e.getClass(), e);
    }

    @ExceptionHandler(AuthorizationException.class)
    @ResponseStatus(HttpStatus.FORBIDDEN)
    public void handleException(AuthorizationException e) {
        logger.debug("{} was thrown", e.getClass(), e);
    }

    @ExceptionHandler(NotFoundException.class)
    @ResponseStatus(HttpStatus.NOT_FOUND)
    public @ResponseBody
    ErrMsg handleException(NotFoundException e) {
        String id = e.getMessage();
        return new ErrMsg("Not Found: " + id);
    }

    @Bean(name = "shiroCredentialsMatcher")
    public CredentialsMatcher credentialsMatcher() {
        HashedCredentialsMatcher matcher = new HashedCredentialsMatcher(Sha256Hash.ALGORITHM_NAME); // alternate MD5
        matcher.setHashIterations(1024);
        matcher.setStoredCredentialsHexEncoded(false);
        return matcher;
    } 
    
    @Bean(name = "shiroAuthczRealm")
//    @DependsOn("lifecycleBeanPostProcessor")
    public ShiroAuthczRealm authczRealm(@Qualifier("shiroCredentialsMatcher")CredentialsMatcher matcher) {
        ShiroAuthczRealm realm = new ShiroAuthczRealm();
        realm.setCredentialsMatcher(matcher);
        return realm;
    }

    @Bean
    public CacheManager cacheManager() {
        // Caching isn't needed in this example, but we will use the MemoryConstrainedCacheManager for this example.
        return new MemoryConstrainedCacheManager();
        
        // EhCacheManager em = new org.apache.shiro.cache.ehcache.EhCacheManager()
        // em.setCacheManagerConfigFile("classpath:ehcache-shiro.xml");
        // return em;
    }
    
    @Bean
    public ShiroFilterChainDefinition shiroFilterChainDefinition() {
        DefaultShiroFilterChainDefinition chainDefinition = new DefaultShiroFilterChainDefinition();
        
//        // logged in users with the 'admin' role
//        chainDefinition.addPathDefinition("/admin/**", "authc, roles[admin]");

//        // logged in users with the 'document:read' permission
//        chainDefinition.addPathDefinition("/docs/**", "authc, perms[document:read]");

//        // all other paths require a logged in user
//        chainDefinition.addPathDefinition("/**", "authc");
        
//        chainDefinition.addPathDefinition("/**", "anon"); // all paths are managed via annotations
    
        // use permissive to NOT require authentication, our controller Annotations will decide that
        chainDefinition.addPathDefinition("/**", "authcBasic[permissive]");
        return chainDefinition;
    }
    
//    @Bean(name = "shiroFilter")
//    public ShiroFilterFactoryBean shiroFilterFactoryBean(@Qualifier("securityManager") SecurityManager sm) {
//        ShiroFilterFactoryBean filter = new ShiroFilterFactoryBean();
//        filter.setSecurityManager(sm);
//        Map<String, String> filterChainDefinitionMapping = new LinkedHashMap<>();
//        filterChainDefinitionMapping.put("/logout", "logout");
//        filterChainDefinitionMapping.put("/users", "authc,roles[admin]");
//        filterChainDefinitionMapping.put("/profies", "authc,roles[user]");

//        filterChainDefinitionMapping.put("/login", "anno");
//        filterChainDefinitionMapping.put("/ajaxlogin", "anno");
//        filterChainDefinitionMapping.put("/static/**", "anno");

        // Customize URL Resource permissions
//        List<Resources> resourcesList = resourcesService.queryAll();
//         for(Resources resources:resourcesList){

//            if (StringUtil.isNotEmpty(resources.getResurl())) {
//                String permission = "perms[" + resources.getResurl()+ "]";
//                filterChainDefinitionMap.put(resources.getResurl(),permission);
//            }
//        }

//        filterChainDefinitionMapping.put("/**", "authc");
        
//        filter.setFilterChainDefinitionMap(filterChainDefinitionMapping);
//        filter.setLoginUrl("/login");
//        filter.setSuccessUrl("/index");
//        filter.setUnauthorizedUrl("/403");
//        return filter;
//    }

//    @Bean(name = "shiroFilter")
//    public AbstractShiroFilter shiroFilter() throws Exception {
//        ShiroFilterFactoryBean shiroFilter = new ShiroFilterFactoryBean();
//        Map<String, String> filterChainDefinitionMapping = new HashMap<>();
//        filterChainDefinitionMapping.put("/cli/health", "authc,roles[guest],ssl[8443]");
//        filterChainDefinitionMapping.put("/login", "authc");
//        filterChainDefinitionMapping.put("/logout", "logout");
//        shiroFilter.setFilterChainDefinitionMap(filterChainDefinitionMapping);
//        shiroFilter.setSecurityManager(securityManager());
//        shiroFilter.setLoginUrl("/login");
//        Map<String, Filter> filters = new HashMap<>();
//        filters.put("anon", new AnonymousFilter());
//        filters.put("authc", new FormAuthenticationFilter());
//        LogoutFilter logoutFilter = new LogoutFilter();
//        logoutFilter.setRedirectUrl("/login?logout");
//        filters.put("logout", logoutFilter);
//        filters.put("roles", new RolesAuthorizationFilter());
//        filters.put("user", new UserFilter());
//        shiroFilter.setFilters(filters);
//        return (AbstractShiroFilter) shiroFilter.getObject();
//    }

    @Bean(name = "securityManager")
    public DefaultWebSecurityManager securityManager() {
        DefaultWebSecurityManager securityManager = new DefaultWebSecurityManager();
//        securityManager.setRealm(jdbcRealm());
        securityManager.setRealm(authczRealm(credentialsMatcher()));
//        securityManager.setCacheManager(cacheManager());
//        securityManager.setSessionManager(sessionManager());
        return securityManager;
    }

    @Autowired
    private DataSource dataSource;

    @Bean(name = "jdbcRealm")
    @DependsOn("lifecycleBeanPostProcessor")
    public JdbcRealm jdbcRealm() {
        JdbcRealm realm = new JdbcRealm();
        HashedCredentialsMatcher credentialsMatcher = new HashedCredentialsMatcher();
        credentialsMatcher.setHashAlgorithmName(Sha256Hash.ALGORITHM_NAME);
        realm.setCredentialsMatcher(credentialsMatcher);
        realm.setDataSource(dataSource);
        realm.init();
        return realm;
    }

    @Bean
    public static LifecycleBeanPostProcessor lifecycleBeanPostProcessor() {
        return new LifecycleBeanPostProcessor();
    }
    
    // Enable Spring AOP to detect Apache Shiro
//    @Bean
//    public AuthorizationAttributeSourceAdvisor authorizationAttributeSourceAdvisor() {
//        AuthorizationAttributeSourceAdvisor a = new AuthorizationAttributeSourceAdvisor();
//        a.setSecurityManager(securityManager());
//        return a;
//    }

//    @Bean
//    @ConditionalOnMissingBean
//    public DefaultAdvisorAutoProxyCreator defaultAdvisorAutoProxyCreator() {
//        final DefaultAdvisorAutoProxyCreator defaultAdvisorAutoProxyCreator = new DefaultAdvisorAutoProxyCreator();
//        defaultAdvisorAutoProxyCreator.setProxyTargetClass(true);
//        return defaultAdvisorAutoProxyCreator;
//    }
    
    // For thymeleaf
//    @Bean
//    public ShiroDialect shiroThymeleafDialect() {
//        return new ShiroDialect();
//    }
    
}