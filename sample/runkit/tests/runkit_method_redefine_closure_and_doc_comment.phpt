--TEST--
runkit_method_redefine() function with closure and doc_comment
--SKIPIF--
<?php
	if(!extension_loaded("runkit") || !RUNKIT_FEATURE_MANIPULATION) print "skip";
	if(version_compare(PHP_VERSION, '5.3.0', '<')) print "skip";
?>
--INI--
display_errors=on
--FILE--
<?php
class runkit_class {
	public function runkit_method() {}
	public function runkitMethod() {}
}
runkit_method_redefine('runkit_class','runkit_method',function() {}, NULL, 'new doc_comment1');
runkit_method_redefine('runkit_class','runkitMethod',function() {}, RUNKIT_ACC_PRIVATE, 'new doc_comment2');
$r1 = new ReflectionMethod('runkit_class', 'runkit_method');
echo $r1->getDocComment(), "\n";
$r2 = new ReflectionMethod('runkit_class', 'runkitMethod');
echo $r2->getDocComment(), "\n";
echo $r2->isPrivate();
?>
--EXPECT--
new doc_comment1
new doc_comment2
1
