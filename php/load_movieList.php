<datalist id="movies_li">
    <?php
    $json = file_get_contents("/rest/titles.json");
    $obj = json_decode($json);
    $movies = $obj->{'movies'};
    foreach ($movies as $movie) {
        echo "<option>$movie</option>";
    }
    ?>
</datalist>