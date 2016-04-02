<?php
$ch = curl_init();
$host =  "http://127.0.0.1:13000";
$inputUrl = $host . "/api/v1/query";

//$tags["type"] = "tp0";                                                                                    
//$tags["speed"] = "11";



//$metric["groupby"][0] = "host";
//$metric["groupby"][1] = "type";
//$metric["groupby"][2] = "speed";

$metric["aggregator"]["name"] = "sum";
$metric["aggregator"]["sampling"]["unit"] = "h";
$metric["aggregator"]["sampling"]["value"] = 1;

$metric["tags"] = $tags;
$metric["name"] = "wyatt_test";

 $input["startabsolute"]= 1389024000000;
  $input["endabsolute"] = 1389024000000 +  60 * 60 * 4 *  5 *1000 + 10500;
$input["metric"] = $metric;

$post = json_encode($input);
echo $post;
curl_setopt($ch, CURLOPT_URL ,$inputUrl);
curl_setopt($ch, CURLOPT_HEADER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query(array("data" => $post)));
$response = curl_exec($ch);
$header_size = curl_getinfo($ch, CURLINFO_HEADER_SIZE);
$header = substr($response, 0, $header_size);
$body = substr($response, $header_size);
var_dump($response);
var_dump(json_decode($body));
curl_close($ch);




