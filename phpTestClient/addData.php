<?php
$ch = curl_init();
$host =  "http://127.0.0.1:13000";
$inputUrl = $host . "/api/v1/datapoints";
/*
$input[0]["timestamp"] = 1389024000989;
$input[0]["value"] = 1.0;
*/
$input[0]["name"] = "wyatt_new";
$input[0]["tags"]["host"] = "server1";
$input[0]["tags"]["speed"] = "10";
$input[0]["tags"]["type"] = "tp2";
$dp = array();

for ($i = 0 ; $i < 3000 ; $i++ ) {
    $ele[0] = time() * 1000 + $i;
    $ele[1] = $i;
    $dp[$i] = $ele;
}

$input[0]["datapoints"] = $dp;
$input[1]["name"] = "wyatt_new";
$input[1]["tags"]["host"] = "server11";
$input[1]["tags"]["speed"] = "11";
$input[1]["tags"]["type"] = "tp1";
for ($i = 0 ; $i < 3000 ; $i++ ) {
    $ele[0] = time() * 1000 + $i * 100 + 3;
    $ele[1] = $i;
    $dp[$i] = $ele;
}
$input[1]["datapoints"] = $dp;

$post = json_encode($input);
echo $post;
curl_setopt($ch, CURLOPT_URL ,$inputUrl);
curl_setopt($ch, CURLOPT_HEADER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query(array("data" => $post)));
$result = curl_exec($ch);
curl_close($ch);




