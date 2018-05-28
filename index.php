<?php

use Uno\UnoClient;
use Uno\PBEmpty;

require_once 'vendor/autoload.php';

require_once 'uno/PBEmpty.php';
require_once 'uno/UnoClient.php';
require_once 'uno/UnoMessage.php';
// 必须引入该文件 否则会报segmentation fault
require_once 'GPBMetadata/Messages.php';

function index()
{
    $client = new UnoClient('127.0.0.1:6110', [
        'credentials' => Grpc\ChannelCredentials::createInsecure(),
    ]);
    // 空参数需要实例化PBEmpty
    list($message) = $client->Rent(new PBEmpty())->wait();
    return $message->getNo();
}

echo index();