package io.stackdocker.ugs.apiserver.security.simple;

import io.stackdocker.ugs.apiserver.help.State;
import io.stackdocker.ugs.apiserver.service.UserService;

import javax.servlet.*;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.util.StringTokenizer;

import org.apache.commons.codec.binary.Base64;
import org.apache.commons.lang.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;

/* @Component */ // filters are registered by default for all the URL’s.
@Order(Ordered.HIGHEST_PRECEDENCE)
public class BasicAuthFilter implements Filter {
    private static final Logger logger = LoggerFactory.getLogger(BasicAuthFilter.class);

    @Value( "${simple.basicauth.disabled:false}" )
    private boolean filterDisabled;

    private String username = "";
    private String password = "";
    private String realm = "Protected";

    @Autowired
    private UserService userService;

    @Override
    public void init(FilterConfig filterConfig) throws ServletException {
        logger.debug("Initiating simple basicauth filter");

        username = filterConfig.getInitParameter("username");
        password = filterConfig.getInitParameter("password");
        String paramRealm = filterConfig.getInitParameter("realm");
        if (StringUtils.isNotBlank(paramRealm)) {
            realm = paramRealm;
        }
    }

    @Override
    public void doFilter(ServletRequest servletRequest, ServletResponse servletResponse, FilterChain filterChain) throws IOException, ServletException {
        if (false == filterDisabled) {
            HttpServletRequest req = (HttpServletRequest) servletRequest;
            HttpServletResponse resp = (HttpServletResponse) servletResponse;

            HeaderMapRequestWrapper wrapper = new HeaderMapRequestWrapper(req);
            wrapper.addHeader("shiro_security", "disabled");
            String path = req.getPathInfo();
            if (path == null) {
                path = req.getRequestURI().substring(req.getContextPath().length());
            }
            String method = req.getMethod();
            if (false == path.toLowerCase().startsWith("/v1/default/users/")) {
                switch (method) {
                    case "POST":
                        logger.info("Skip basic auth, go to default servlet");
                        filterChain.doFilter(wrapper, servletResponse);
                        return;
                    default:
                        String remote_addr = servletRequest.getRemoteAddr();
                        wrapper.addHeader("remote_addr", remote_addr);
                }
            }

            String authHeader = req.getHeader("Authorization");
            if (authHeader != null) {
                StringTokenizer st = new StringTokenizer(authHeader);
                if (st.hasMoreTokens()) {
                    String basic = st.nextToken();

                    if (basic.equalsIgnoreCase("Basic")) {
                        try {
                            String credentials = new String(Base64.decodeBase64(st.nextToken()), "UTF-8");
                            logger.debug("Credentials: " + credentials);
                            int p = credentials.indexOf(":");
                            if (p != -1) {
                                String _username = credentials.substring(0, p).trim();
                                String _password = credentials.substring(p + 1).trim();

//                                if (!username.equals(_username) || !password.equals(_password)) {
//                                    unauthorized(resp, "Bad credentials");
//                                }
                                State state = userService.verifyWithBasicAuth(_username, _password);
                                switch (state) {
                                    case UNAUTHENTICATED_USER:
                                        logger.warn(  "Failed to authenticate user " + _username);
                                        unauthorized(resp, state.message());
                                        break;
                                    case BLOCKED_USER:
                                        logger.info(_username + " blocked");
                                        denied(resp, state.message());
                                        break;
                                        default:
                                            logger.info(_username + " authenticated");
                                            if (true == path.toLowerCase().startsWith("/v1/default/user-actions/")) {
                                                wrapper.addHeader("username", _username);
                                            }
                                            filterChain.doFilter(wrapper, servletResponse);
                                }
                                return;
                            } else {
                                unauthorized(resp, "Invalid authentication token");
                            }
                        } catch (UnsupportedEncodingException e) {
                            throw new Error("Couldn't retrieve authentication", e);
                        }
                    }
                }
                filterChain.doFilter(wrapper, servletResponse);
            } else {
                unauthorized(resp);
            }
            return;
        }
        filterChain.doFilter(servletRequest, servletResponse);
    }

    @Override
    public void destroy() {
        logger.debug("Destroying simple basicauth filter");
    }

    private void unauthorized(HttpServletResponse response, String message) throws IOException {
        response.setHeader("WWW-Authenticate", "Basic realm=\"" + realm + "\"");
        response.sendError(401, message);
    }

    private void unauthorized(HttpServletResponse response) throws IOException {
        unauthorized(response, "Unauthorized");
    }

    private void denied(HttpServletResponse response, String message) throws IOException {
        response.sendError(403, message);
    }

    private void denied(HttpServletResponse response) throws IOException {
        response.sendError(403, "Forbidden");
    }
}
