--TEST--
runkit_default_property_remove() remove properties with inheritance
--SKIPIF--
<?php if(!extension_loaded("runkit") || !RUNKIT_FEATURE_MANIPULATION) print "skip";
      if(array_shift(explode(".", PHP_VERSION)) < 5) print "skip";
?>
--INI--
error_reporting=E_ALL
display_errors=On
--FILE--
<?php
class RunkitClass {
    public $publicProperty = 1;
    private $privateProperty = "a";
    protected $protectedProperty = "b";
    private static $staticProperty = "s";
    public $removedProperty = "r";
}

class RunkitSubClass extends RunkitClass {
    public $publicProperty = 2;
    private $privateProperty = "aa";
    protected $protectedProperty = "bb";
    protected $staticProperty = "ss";
    function getPrivate() {return $this->privateProperty;}
}
class RunkitSubSubClass extends RunkitSubClass {
    protected $protectedProperty = "cc";
}

ini_set('display_errors', 1);
ini_set('error_reporting', E_ALL);

$className = 'RunkitClass';
$obj = new RunkitSubSubClass();

runkit_default_property_remove($className, 'publicProperty');
runkit_default_property_remove($className, 'privateProperty');
runkit_default_property_remove($className, 'protectedProperty');
runkit_default_property_remove('RunkitSubClass', 'removedProperty');
runkit_default_property_remove($className, 'removedProperty');
print_r(new RunkitClass());
print_r(new RunkitSubClass());
$out = print_r($obj, true);
$version = explode(".", PHP_VERSION);
if ((int) $version[0] == 5 && (int) $version[1] < 4) {
	$out = preg_replace("/\n\s+\[privateProperty:RunkitClass:private\] => a$/m", "", $out);
}
if ((int) $version[0] == 5 && (int) $version[1] < 3) {
	$out = preg_replace("/\n\s+\[privateProperty:private\] => a$/m", "", $out);
}
print $out;
print_r($obj->getPrivate());
?>
--EXPECTF--
Notice: runkit_default_property_remove(): Making RunkitSubSubClass::privateProperty public to remove it from class without objects overriding in %s on line %d
RunkitClass Object
(
)
RunkitSubClass Object
(
    [publicProperty] => 2
    [privateProperty%sprivate] => aa
    [protectedProperty:protected] => bb
    [staticProperty:protected] => ss
)
RunkitSubSubClass Object
(
    [protectedProperty:protected] => cc
    [publicProperty] => 2
    [privateProperty%sprivate] => aa
    [staticProperty:protected] => ss
    [removedProperty] => r
    [privateProperty] => a
)
aa
