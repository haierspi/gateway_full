<?php

function examplesEcho($args)
{
    $result = array("ErrorCode" => 0);
    print_r($args["Body"]);
    $result["Data"] = "Hello, " . $args["Body"];
    return array("Error" => null, "Result" => $result);
}
