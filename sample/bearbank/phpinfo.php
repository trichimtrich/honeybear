<html>
<body>
<?php

echo $_SERVER["HTTP_RID"]."<br>";

echo "Server IP: ".$_SERVER["SERVER_ADDR"];
if ($_SERVER["SERVER_ADDR"]=="192.168.200.107") 
	echo "<h1>I'M PRODUCT</h1><br>";
else 
	echo "<h1>I'M FARM</h1><br>";

session_start();

if (!isset($_SESSION['remember']) && isset($_COOKIE['PHPSESSID'])) {
	$_SESSION['remember'] = 1;
	setcookie("PHPSESSID", $_COOKIE["PHPSESSID"], time() + 365*24*3600);
}

$value = 'xxx';

setcookie("TestCookie1", rand());
setcookie("TestCookie2", rand(), time()+3600);  /* expire in 1 hour */
setcookie("TestCookie3", "hihi=;hoho");

if (isset($_GET['action']) && $_GET['action']==='login') {
	$_SESSION['login'] = 1;
}

if (isset($_GET['action']) && $_GET['action']==='logout') {
	session_unset();
	session_destroy();
	setcookie("PHPSESSID", "");
}


if (isset($_SESSION['login'])) {
?>
	Hi boss!
	<form method="GET" action=""><input type="submit" name="action" value="logout"></form>
<?php
}
else
{
?>
	You must login
	<form method="GET" action=""><input type="submit" name="action" value="login"></form>
<?php
}

var_dump($_POST);
echo "<br/><br/>";
var_dump($_GET);
echo "<br/><br/>";
var_dump($_SESSION);
echo "<br/><br/>";
var_dump($_COOKIE);
echo "<br/><br/>";

phpinfo();

?>
</body>
</html>