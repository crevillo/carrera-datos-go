<?php

$start = microtime(true);
$csv = fopen('n.csv', 'r');
while (($currRow = fgetcsv($csv, 256)) !== FALSE) {
}
fclose($csv);
$time1 = microtime(true) - $start;

print "PHP Takes  $time1 seconds \n";

