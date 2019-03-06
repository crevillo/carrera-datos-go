<?php
/**
 * Created by PhpStorm.
 * User: carlos
 * Date: 2019-03-06
 * Time: 20:53
 */


$start = microtime(true);
$data = array_map("str_getcsv", file('n.csv'));
$time2 = microtime(true) - $start;

$start = microtime(true);
$csv = fopen('n.csv', 'r');
while (($currRow = fgetcsv($csv, 256)) !== FALSE) {
}
fclose($csv);
$time1 = microtime(true) - $start;




print "While $time1 \n";
print "Map $time2 \n";
