#!/bin/sh

curl -v -X POST "http://192.168.49.2:30007/telemetry" -d "{\"device_id\": $(shuf -i 0-1000 -n 1),\"health\": \"HEALTHTY\",\"gps_level\": $(shuf -i 0-1000 -n 1)}"