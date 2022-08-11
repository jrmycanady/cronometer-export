#!/bin/bash

./cronometer-export -u <USERNAME> -p <PASSWORD> -e -0d -s -100m > cronometer.csv
./gdrive upload -p <FOLDER-ID> cronometer.csv
