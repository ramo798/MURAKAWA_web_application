<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>

<%
	String inputstr = request.getParameter("id");
	String inputpass = request.getParameter("password");
	String inputradio = request.getParameter("anonymous");
	String inputmizaru = request.getParameter("mizaru");
	String inputkikazaru = request.getParameter("kikazaru");
	String inputsyaberazaru = request.getParameter("syaberazaru");

%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>receive</title>
</head>
<body>
	入力されたID:<%=inputstr %><br>
	入力されたパスワード:<%= inputpass %><br>
	入力されたラジオボタン(匿名ログイン):<%=inputradio %><br>
	入力チェックボックス(見ざる):<%=inputmizaru %><br>
	入力チェックボックス(聞かざる):<%=inputkikazaru %><br>
	入力チェックボックス(喋らざる):<%=inputsyaberazaru %><br>
	※全ての入力は文字列となっている<br>
	<a href = "index.html">入力ページに戻る</a>
</body>
</html>