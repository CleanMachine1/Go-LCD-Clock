# Go-LCD-Clock

This repo just holds a small project for a LCD display I put to use.

To run, first compile main.go and weather/weather.go

Then setup automatic runnings of the script.

I'd recommend this in a crontab

``` sh
2,16,31 * * * * bash /path/to/speed.sh 
1,15,30,45 * * * * /path/to/weather/weather
@reboot bash /path/to/startup.sh
```

Personally I keep my ./main binary being run within a `screen` so that I can monitor it if it crashes.

Feel free to contribute to help it use less system resources!
