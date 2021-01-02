<?php
$user_movie = $_GET['user_movie'];
include 'select_movie.php';
?>

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Recommender</title>
    </head>
    <body>
        <h2>Recommender</h2>
        <table>
            <style>
                table, th, td {
                    /* border: 1px solid rgb(71, 141, 207); */
                    border-collapse: collapse;
                }
                th, td {
                    padding: 8px;
                    text-align: left;
                }
            </style>
            <th>The movie you choose is:</th>
            <tr>
                <?php
                $obj = json_decode($server_output);
                $movie = $obj->{'movie'};
                echo  "<td>$movie</td>";
                ?>
            </tr>
            <tr>
                <td style="color: white;">_</td>
            </tr>
            <tr>
                <?php
                if ($movie != "The movie is not in our database. Please try another one.") {
                    echo "<th>Recommendations:</th>";
                }
                ?>
            </tr>
        </table>
        <ul>
            <?php
            $rcmd_movies = $obj->{'recommend'};
            foreach ($rcmd_movies as $rcmd) {
                echo "<li style='padding-bottom: 5px;'>$rcmd</li>";
            }
            ?>
        </ul>
        <br><br>
        <p style="padding-left: 15em;">Back to <a href="./index.php" style="color: blue;"><b><i>home</i></b></a> </p>
    </body>
</html>