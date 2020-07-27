<?php
error_reporting(E_ALL ^ E_WARNING); 
$status = get_headers("http://rcmd-go:3000/v1/movies");

if ($status[0] != "HTTP/1.1 200 OK") {
    exit("Something when wrong! :(");
}
?>