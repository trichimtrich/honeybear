<?php
require_once("hook.php");

$serverName="localhost";
$dbusername="root";
$dbpassword="chim";
$dbname="bank_db";
mysql_connect($serverName,$dbusername,$dbpassword)/* or die('the website is down for maintainance')*/;
mysql_select_db($dbname) or die(mysql_error());
?>