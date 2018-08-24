#!/bin/bash

mc --quiet --no-color mb node1/mybucket

function put() {
    n=$1
    mc --quiet --no-color cp 100mb/100mb node1/mybucket/100mb-$n >/tmp/100mb-$n 2>&1
    if [ $? -ne 0 ]; then
	echo "Upload node1/mybucket/100mb-$n failed"
	cat /tmp/100mb-$n
    else
	rm -f /tmp/100mb-$n
    fi
}

for (( i = 0; i < 1000/100; i++ )); do
    for (( j = 1; j <= 100; j++ )); do
	(( n = j + i * 100 ))
	put $n &
    done

    wait
    echo "Waiting for GC to kick at server"
    sleep 4m
done
