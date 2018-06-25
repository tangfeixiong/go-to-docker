<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core"%>
<%@ taglib prefix="spring" uri="http://www.springframework.org/tags"%>

<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>index</title>
</head>
<body>
<form action="${pageContext.request.contextPath}/exam/user/signin" method="post">
    用户名：<input type="text" name="username"/><br>
    <br>
    密码：<input type="password" name="password"/><br>
    <br>
    <input type="submit" value="signIn" id="submit" name"submit">
</form>
</body>
</html>