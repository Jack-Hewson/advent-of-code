<?php
$lines = file('input.txt');

$grid = [];

$accessible = 0;

foreach ($lines as $rowIndex => $line) {
    $chars = str_split($line);
    foreach ($chars as $charIndex => $char) {
        if ($char == "@") {
            $boundary_array = array(
                $lines[$rowIndex - 1][$charIndex - 1] ?? null,
                $lines[$rowIndex - 1][$charIndex] ?? null,
                $lines[$rowIndex - 1][$charIndex + 1] ?? null,
                $lines[$rowIndex][$charIndex - 1] ?? null,
                $lines[$rowIndex][$charIndex + 1] ?? null,
                $lines[$rowIndex + 1][$charIndex - 1] ?? null,
                $lines[$rowIndex + 1][$charIndex] ?? null,
                $lines[$rowIndex + 1][$charIndex + 1] ?? null,
            );

            $boundary_counter = 0;

            foreach ($boundary_array as $boundary_cell) {
                if ($boundary_cell == "@") {
                    $boundary_counter++;
                }
            }

            if ($boundary_counter < 4) {
                $accessible++;
            }
        }
    }
}

echo $accessible;
