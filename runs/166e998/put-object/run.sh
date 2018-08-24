#!/bin/bash

mc --quiet --no-color mb node1/mybucket

for (( i = 1; i <= 135; i++ )); do
    mc --quiet --no-color cp 2gb/2gb node1/mybucket/2gb-$i &
done

wait
