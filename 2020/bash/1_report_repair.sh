#!/bin/bash

in="../1_input"
nos=$(sort -n < $in)

for x in "$nos"; do
  for y in "$nos"; do
    echo $y
    if (( $x + $y == 2020 )); then
      #echo "x: $x, y: $x"
      echo -n
    fi
  done
done
