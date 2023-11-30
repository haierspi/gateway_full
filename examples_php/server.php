<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/services/echo.php';


use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

$connection = new AMQPStreamConnection('127.0.0.1', 5672, 'test', 'test');
$channel = $connection->channel();
$channel->queue_declare('examples_1.0', false, false, false, false);

function on_request($msg)
{
    echo($msg . '\n');
    $msg_obj = json_decode($msg, true);

    $result = null;
    switch ($msg_obj["Method"]) {
        case "Examples.Echo":
            $result = examplesEcho($msg_obj["Params"]);
            break;
        
        default:
            $result = array("Error"=>"Call server fail.", "Result"=>null);
            break;
    }

    return json_encode($result);
}

$callback = function ($req) {
    $msg = new AMQPMessage(
        (string) on_request($req->body),
        array('correlation_id' => $req->get('correlation_id'))
    );

    $req->delivery_info['channel']->basic_publish(
        $msg,
        '',
        $req->get('reply_to')
    );
    $req->delivery_info['channel']->basic_ack(
        $req->delivery_info['delivery_tag']
    );
};

$channel->basic_qos(null, 1, null);
$channel->basic_consume('examples_1.0', '', false, false, false, false, $callback);

while ($channel->is_consuming()) {
    $channel->wait();
}

$channel->close();
$connection->close();
