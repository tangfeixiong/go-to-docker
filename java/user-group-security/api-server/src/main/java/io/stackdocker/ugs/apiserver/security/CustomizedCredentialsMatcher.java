package io.stackdocker.ugs.apiserver.security;

import org.apache.shiro.authc.credential.SimpleCredentialsMatcher;
import org.apache.shiro.authc.AuthenticationInfo; 
import org.apache.shiro.authc.AuthenticationToken; 
import org.apache.shiro.authc.UsernamePasswordToken;


public class CustomizedCredentialsMatcher extends SimpleCredentialsMatcher {
    @Override
    public boolean doCredentialsMatch(AuthenticationToken token, Authentication info) {
        UsernamePasswordToken upt = (UsernamePasswordToken) token;
        // add salt if required
        String submitted = new String(upt.getPassword());
        // get from db
        String saved = (String)info.getgetCredentials();
        return saved.equals(submitted);

//        String tokenCredentials = charArrayToString(tok.getCredentials());
//        String reverseToken = StringUtils.reverse(tokenCredentials);
//        String encryptedToken = new Sha256Hash(reverseToken).toString();

//        String accountCredentials = charArrayToString(info.getCredentials());
//        return accountCredentials.equals(encryptedToken);
    }

    private String charArrayToString(Object credentials) {
        return new String((char[]) credentials);
    }

}