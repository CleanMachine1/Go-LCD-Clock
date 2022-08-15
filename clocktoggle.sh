#!/bin/bash
#
# Run to toggle the clock

content=$(cat /tmp/status.txt)

if [ $content == "1" ];
then
echo -n 0 > /tmp/status.txt
echo "off"
exit 0
fi

echo -n 1 > /tmp/status.txt
echo "on"
