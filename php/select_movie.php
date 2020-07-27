<?php
error_reporting(E_ALL ^ E_WARNING);

$user_movie = str_replace('/', '', $user_movie) ;

$ch = curl_init();

curl_setopt($ch, CURLOPT_URL,"http://rcmd-go:3000/v1/movies/$user_movie");
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);


$server_output = curl_exec($ch);
curl_close ($ch);

if ($server_output == false) {
    exit("Something when wrong! :(");
}
?>