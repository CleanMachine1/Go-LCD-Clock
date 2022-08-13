#!/bin/bash
# This could probably be made better but it works for now
# You must install speedtest-cli


speedtest-cli --simple > /tmp/temp1 || exit 1

sed '2!d' /tmp/temp1 > /tmp/temp2
sed -i -e 's/load//g' /tmp/temp2
tr -d '\n' < /tmp/temp2 > /tmp/temp3


cp /tmp/temp3 /tmp/speedtest.txt



