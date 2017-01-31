--TEST--
runkit_default_property_remove() and runkit_default_property_add() functions on classes having dynamic properties (without overriding objects)
--SKIPIF--
<?php if(!extension_loaded("runkit") || !RUNKIT_FEATURE_MANIPULATION) print "skip";
      if(array_shift(explode('.', PHP_VERSION)) < 5) print "skip";
?>
--FILE--
<?php
class A {
	private $c = 1;
	public $d = 'd';
	public function getC() {return $this->c;}
}
class B extends A {}

$o = new B;
echo $o->getC(), "\n";
echo $o->b, $o->d, "\n";
$o->b = 'b';
echo $o->getC(), "\n";
echo $o->b, $o->d, "\n";
echo "remove\n";
runkit_default_property_remove('A', 'c');
echo $o->getC(), "\n";
echo $o->b, $o->d, "\n";
echo "add public\n";
runkit_default_property_add('A', 'c', 2);
echo $o->getC(), "\n";
echo $o->b, $o->d, "\n";
echo "remove\n";
runkit_default_property_remove('A', 'c');
$o1 = new B;
echo $o1->getC(), "\n";
echo $o->getC(), "\n";
echo $o->b, $o->d, "\n";
echo "add private\n";
runkit_default_property_add('A', 'c', 3, RUNKIT_ACC_PRIVATE);
echo $o->getC(), "\n";
echo $o->b, $o->d, "\n";
echo "remove\n";
runkit_default_property_remove('A', 'c');
echo $o->getC(), "\n";
echo $o->b, $o->d, "\n";
echo "add public to B\n";
runkit_default_property_add('B', 'c', 2);
echo $o->getC(), "\n";
echo $o->b, $o->d, "\n";
--EXPECTF--
1

Notice: Undefined property: B::$b in %s on line %d
d
1
bd
remove

Notice: runkit_default_property_remove(): Making B::c public to remove it from class without objects overriding in %s on line %d
1
bd
add public
1
bd
remove

Notice: Undefined property: B::$c in %s on line %d

1
bd
add private
1
bd
remove

Notice: runkit_default_property_remove(): Making B::c public to remove it from class without objects overriding in %s on line %d
1
bd
add public to B
1
bd
