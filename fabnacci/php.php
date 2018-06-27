<?php
function fabnacci(int $pos)
{
    $pos = intval($pos);
    if ($pos <= 0) {
              throw new Exception('must be a int bigger than 0');
    } elseif ($pos <= 2) {
        return 1;
    } else {
        $pre = 1;
        $ppre = 1;
        $sum = 0;
        for ($i = 2; $i < $pos; $i++) {
            $sum = $ppre + $pre;
            $ppre = $pre;
            $pre = $sum;
        }
        return $sum;
    }
}

//$arr = [];
//for ($i = 1; $i < 20; $i++) {
//    $arr[] = fabnacci($i);
//}
//echo implode(',', $arr);

$t=microtime(true);
echo fabnacci(30),PHP_EOL;
useTime($t);

function useTime($preTime){
    $t= microtime(true)-$preTime;
    $t=$t*1000;
    echo 'use time of: ',PHP_EOL;
    echo number_format($t,6),' ms',PHP_EOL;
//    echo round($t,6),PHP_EOL;
}

echo PHP_EOL,'========== use recurse way',PHP_EOL;
function fibonacci($n) {
    if ($n < 2) {
        return $n;
    }
    return fibonacci($n - 2) + fibonacci($n - 1);
}
$t=microtime(true);
echo fibonacci(30),PHP_EOL;
useTime($t);
