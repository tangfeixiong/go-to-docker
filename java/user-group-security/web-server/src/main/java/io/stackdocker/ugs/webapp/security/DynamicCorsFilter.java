/*
  Inspired by:
    https://gist.github.com/zeroows/80bbe076d15cb8a4f0ad
 */
package io.stackdocker.ugs.webapp.security;

import java.io.IOException;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.web.filter.OncePerRequestFilter;

/**
 * Enabling CORS support  - Access-Control-Allow-Origin
 * @author zeroows@gmail.com
 *
 * <code>
<!-- Add this to your web.xml to enable "CORS" -->
<filter>
<filter-name>cors</filter-name>
<filter-class>io.stackdocker.ugs.webapp.security.DynamicCorsFilter</filter-class>
</filter>

<filter-mapping>
<filter-name>cors</filter-name>
<url-pattern>/*</url-pattern>
</filter-mapping>
 * </code>
 */
public class DynamicCorsFilter extends OncePerRequestFilter {
    private static final Log LOG = LogFactory.getLog(DynamicCorsFilter.class);

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        String path = request.getPathInfo();
        if (path == null) {
            path = request.getRequestURI().substring(request.getContextPath().length());
        }

        // todo: verify host and path into database to show whether it can do CORS

        response.addHeader("Access-Control-Allow-Origin", "*");

        if (request.getHeader("Access-Control-Request-Method") != null && "OPTIONS".equals(request.getMethod())) {

            LOG.trace("Sending Header....");
            // CORS "pre-flight" request
            response.addHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE");
//			response.addHeader("Access-Control-Allow-Headers", "Authorization");
            response.addHeader("Access-Control-Allow-Headers", "Content-Type");
            response.addHeader("Access-Control-Max-Age", "1");
        }

        filterChain.doFilter(request, response);
    }

}
