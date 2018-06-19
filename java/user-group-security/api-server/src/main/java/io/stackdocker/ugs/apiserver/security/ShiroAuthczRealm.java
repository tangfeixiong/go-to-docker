package io.stackdocker.ugs.apiserver.security;

import io.fairymagic.core.ugs.domain.RoleBinding;
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

import io.stackdocker.ugs.apiserver.service.PermissionService;
import io.stackdocker.ugs.apiserver.service.ResourceService;
import io.stackdocker.ugs.apiserver.service.RoleService;
import io.stackdocker.ugs.apiserver.service.UserService;
import io.fairymagic.core.ugs.domain.Role;
import io.fairymagic.core.ugs.domain.User;

import java.util.Arrays;
import java.util.Collection;

@Component
public class ShiroAuthczRealm extends AuthorizingRealm {
    private static Logger logger = LoggerFactory.getLogger(ShiroAuthczRealm.class);

    protected UserService userService = null;
    
    @Autowired private ResourceService resourceService;
    
    @Autowired private RoleService roleService;
    
    @Autowired private PermissionService permissionSerivce;

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
        this.userService = userService;
    }

    // Authenticate
    protected AuthenticationInfo doGetAuthenticationInfo(AuthenticationToken authcToken) throws AuthenticationException {
        logger.debug("do Authentication");
        // String username = (String)authcToken.getPrincipal();
        UsernamePasswordToken token = (UsernamePasswordToken) authcToken;
        User user = userService.findOneByName(token.getUsername());
        if( user != null ) {
            if (true == userService.isActive(user)) {
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
        User user = userService.getOne(userId);
        if( user != null ) {
            SimpleAuthorizationInfo info = new SimpleAuthorizationInfo();
            Collection<Role> roles = roleService.getWithBindings(user.getRoleBindings());
            for( Role role : roles ) {
                info.addRole(role.getName());
            }
            for( RoleBinding binding: user.getRoleBindings()) {
                info.addStringPermissions(Arrays.asList("create", "read", "update", "delete"));
            }
            return info;
        } else {
            logger.warn("Failed to get user with name " + user.getUsername());
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

