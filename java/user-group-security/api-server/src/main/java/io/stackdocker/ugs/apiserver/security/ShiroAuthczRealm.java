package io.stackdocker.ugs.apiserver.security;

import javax.annotation.Resource;

import org.apache.shiro.authc.*;
import org.apache.shiro.authc.credential.Sha256CredentialsMatcher;
import org.apache.shiro.authz.AuthorizationInfo;
import org.apache.shiro.authz.SimpleAuthorizationInfo;
import org.apache.shiro.realm.AuthorizingRealm;
import org.apache.shiro.subject.PrincipalCollection;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import io.stackdocker.ugs.apiserver.service.UserService;
import io.fairymagic.auth.api.Role;
import io.fairymagic.exam.api.dao.User;

@Component
public class ShiroAuthczRealm extends AuthorizingRealm {
    private static Logger logger = LoggerFactory.getLogger(ShiroAuthczRealm.class);

    protected UserService userDAO = null;
    
    @Resource private ResourceService resourceService;
    
    @Autowired private RoleDao roleService;
    
    @Autowired private PermissionService permissionSerivce 

    public ShiroAuthczRealm() {
        setName("ShiroAuthczRealm"); //This name must match the name in the User class's getPrincipals() method
        setCredentialsMatcher(new Sha256CredentialsMatcher());
        
//        HashedCredentialsMatcher matcher = new HashedCredentialsMatcher();
//        matcher.setHashAlgorithmName(Sha256Hash.ALGORITHM_NAME);
//        matcher.setHashIterations(1024);
//        matcher.setStoredCredentialsHexEncoded(false);        
    }

    @Autowired
    public void setUserDAO(UserService userService) {
        this.userDAO = userService;
    }

    // Authenticate
    protected AuthenticationInfo doGetAuthenticationInfo(AuthenticationToken authcToken) throws AuthenticationException {
        logger.debug("do Authentication");
        // String username = (String)authcToken.getPrincipal();
        UsernamePasswordToken token = (UsernamePasswordToken) authcToken;
        User user = userDAO.findByName(token.getUsername());
        if( user != null ) {
            if (0 == user.getEnabled()) {
                throw new LockedAccountException();
            }
            // SimpleAuthenticationInfo info = new SimpleAuthenticationInfo(
            //         user, user.getPassword(), ByteSource.Util.bytes(username), getName());
            // Session session = SecurityUtils.getSubject().getSession();
            // session.setAttribute("userSession", user);
            // session.setAttribute("userSessionId", user.getId());
            // return info;    
            return new SimpleAuthenticationInfo(user.getId(), user.getPassword(), getName());
        } else {
            logger.warn("Failed to get user with name " + token.getUsername());
            return null;
        }

//        UsernamePasswordToken upToken = (UsernamePasswordToken) token;

//        String username = upToken.getUsername();
//        checkNotNull(username, "Null usernames are not allowed by this realm.");

//        String password = userDAO.getPassword(username);
//        checkNotNull(password, "No account found for user [" + username + "]");

//        return new SimpleAuthenticationInfo(username, password.toCharArray(), getName());
    }

    // Authorize
    protected AuthorizationInfo doGetAuthorizationInfo(PrincipalCollection principals) {
        logger.debug("do Authorization");
        // User user = (User) SecurityUtils.getSubject().getPrincipal();
        // Map<String, Object> map = new HashMap<String, Object>();
        // map.put("userId", user.getId());
        // List<Resource> resources = resourceService.loadResourcesByUser(map);
        // SimpleAuthorizationInfo info = new SimpleAuthorizationInfo();
        // for (Resource item: resources) {
        //     info.AddStringPermission(item.getResUrl());
        // }
        Long userId = (Long) principals.fromRealm(getName()).iterator().next();
        User user = userDAO.getUser(userId);
        if( user != null ) {
            SimpleAuthorizationInfo info = new SimpleAuthorizationInfo();
            for( Role role : user.getRoles() ) {
                info.addRole(role.getName());
                info.addStringPermissions( role.getPermissions() );
            }
            return info;
        } else {
            logger.warn("Failed to get user with name " + token.getUsername());
            return null;
        }

//        checkNotNull(principals, "PrincipalCollection method argument cannot be null.");

//        String username = (String) principals.getPrimaryPrincipal();
//        SimpleAuthorizationInfo info = new SimpleAuthorizationInfo(userDAO.getRoles(username));
//        info.setStringPermissions(userDAO.getPermissions(username));
//        return info;
    }

    private void checkNotNull(Object reference, String message) {
        if (reference == null) {
            throw new AuthenticationException(message);
        }
    }

}

