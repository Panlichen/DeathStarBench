#! /bin/bash

for file in $(cat yamls)
do
    echo $file
    sed '$a\ \ \ \ \ \ dnsPolicy: "Default"' $file | cat > a; mv a $file

    echo "======================="
    echo "***********************"
    echo "======================="
done