<?php

$test_array = array(
    array(".", ".", "."),
    array(".", "@", "."),
    array(".", ".", ".")
);

foreach ($test_array as $rowIndex => $row) {
    foreach ($row as $colIndex => $cell) {
        echo "[$rowIndex,$colIndex] = $cell ";
    }
    echo PHP_EOL;
}
