#! /bin/bash

for file in $(cat yamls)
do

echo $file
# sed '$a\ \ \ \ \ \ dnsPolicy: "Default"' $file | cat > a; mv a $file
cat >> ../$file << EOF
      schedulerName: poseidon
EOF

echo "======================="
echo "***********************"
echo "======================="

done