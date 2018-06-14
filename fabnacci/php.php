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

$arr = [];
for ($i = 1; $i < 20; $i++) {
    $arr[] = fabnacci($i);
}
echo implode(',', $arr);