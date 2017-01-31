<?php

if ($_SERVER["SERVER_ADDR"]!="192.168.200.107") {

//disable_functions =exec,passthru,shell_exec,system,proc_open,popen,curl_exec,curl_multi_exec,parse_ini_file,show_source
runkit_function_rename("exec", "bear_exec");
runkit_function_rename("passthru", "bear_passthru");
runkit_function_rename("shell_exec", "bear_shell_exec");
runkit_function_rename("system", "bear_system");
runkit_function_rename("mysql_query", "bear_mysql_query");

//echo exec("dir");
//passthru("dir");
//echo shell_exec("dir");
//system("dir");

function pushELK($url, $content) {
	$curl = curl_init($url);
	curl_setopt($curl, CURLOPT_HEADER, false);
	curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
	curl_setopt($curl, CURLOPT_HTTPHEADER,
	        array("Content-type: application/json"));
	curl_setopt($curl, CURLOPT_POST, true);
	curl_setopt($curl, CURLOPT_POSTFIELDS, json_encode($content));

	$json_response = curl_exec($curl);
	$status = curl_getinfo($curl, CURLINFO_HTTP_CODE);

	if ( $status != 201 ) {
	    die("Error: call to URL $url failed with status $status, response $json_response, curl_error " . curl_error($curl) . ", curl_errno " . curl_errno($curl));
	}
	curl_close($curl);
	//$response = json_decode($json_response, true);
}

runkit_function_add("exec", '$a, $b=NULL, $c=NULL', '
	$content = array(
		"rid" => $_SERVER["HTTP_RID"],
		"cmd" => $a,
		"time" => gmdate(DATE_RFC3339)
	);
	pushELK("http://192.168.200.1:9200/cmd/bear/", $content);

	putenv("rid=".$_SERVER["HTTP_RID"]);
	if (!is_null($b)) {
		if (!is_null($c))
			return bear_exec($a, $b, $c);
		return bear_exec($a, $b);
	}
	return bear_exec($a);');


runkit_function_add('passthru', '$a, $b=NULL', '
	$content = array(
		"rid" => $_SERVER["HTTP_RID"],
		"cmd" => $a,
		"time" => gmdate(DATE_RFC3339)
	);
	pushELK("http://192.168.200.1:9200/cmd/bear/", $content);

	putenv("rid=".$_SERVER["HTTP_RID"]);
	if (!is_null($b))
		return bear_passthru($a, $b);
	return bear_passthru($a);');


runkit_function_add('shell_exec', '$a', '
	$content = array(
		"rid" => $_SERVER["HTTP_RID"],
		"cmd" => $a,
		"time" => gmdate(DATE_RFC3339)
	);
	pushELK("http://192.168.200.1:9200/cmd/bear/", $content);

	putenv("rid=".$_SERVER["HTTP_RID"]);
	return bear_shell_exec($a);');


runkit_function_add('system', '$a, $b=NULL', '
	$content = array(
		"rid" => $_SERVER["HTTP_RID"],
		"cmd" => $a,
		"time" => gmdate(DATE_RFC3339)
	);
	pushELK("http://192.168.200.1:9200/cmd/bear/", $content);

	putenv("rid=".$_SERVER["HTTP_RID"]);
	if (!is_null($b))
		return bear_system($a, $b);
	return bear_system($a);');


runkit_function_add('mysql_query', '$a, $b=NULL', '
	$content = array(
		"rid" => $_SERVER["HTTP_RID"],
		"query" => $a,
		"time" => gmdate(DATE_RFC3339)
	);

	pushELK("http://192.168.200.1:9200/query/bear/", $content);

	if (!is_null($b))
		$res = bear_mysql_query($a, $b);
	else
		$res = bear_mysql_query($a);
	return $res;');

}

?>
