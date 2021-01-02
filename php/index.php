<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Recommender</title>
    </head>
    <body>
        <h2>Recommender</h2>

        <?php include 'check_status.php'; ?>

        <form action="rcmd.php" method="get" autocomplete="off">
            <table>
                <style> 
                    table, th, td {
                        border-collapse: collapse;
                    }
                    th, td {
                        padding: 8px;
                        text-align: center;
                    }
                </style>
                <th align="left">Please enter a movie you like:</th>
                <tr>
                    <td><input type="text" name="user_movie" size="40" list="movies_li"> <br></td>
                </tr>
                <tr>
                    <td><input type="submit" value="Click to get recommendations"></td>
                </tr>
            </table>
        </form>
    </body>
</html>

<?php include 'load_movieList.php'; ?>