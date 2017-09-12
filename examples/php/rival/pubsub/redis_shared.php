<?php

if (PHP_SAPI !== 'cli') {
    die("Example scripts are meant to be executed locally via CLI.");
}

require __DIR__.'/vendor/autoload.php';

function redis_version($info)
{
    if (isset($info['Server']['redis_version'])) {
        return $info['Server']['redis_version'];
    } elseif (isset($info['redis_version'])) {
        return $info['redis_version'];
    } else {
        return 'unknown version';
    }
}

$single_server = array(
    'host' => '172.18.0.3',
    'port' => 6379,
    'database' => 15,
);

$multiple_servers = array(
    array(
       'host' => '127.0.0.1',
       'port' => 6379,
       'database' => 15,
       'alias' => 'first',
    ),
    array(
       'host' => '127.0.0.1',
       'port' => 6380,
       'database' => 15,
       'alias' => 'second',
    ),
);
