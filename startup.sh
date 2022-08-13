#!/bin/bash
sleep 15

touch /tmp/speedtest.txt
touch /tmp/weather.txt
touch /tmp/status.txt

bash /home/$USER/clock-go/speed.sh
python3.9 /home/$USER/clock-go/weather.py

bash /home/$USER/clock-go/status.sh
sleep 3
/home/$USER/clock-go/main &
