#!/bin/bash

for (( i = 1; i <= 26; i++ )); do
    echo "Downloading node1/mybucket/100mb-$i"
    mc --quiet --no-color cat node1/mybucket/100mb-$i >/dev/null &
done

wait
