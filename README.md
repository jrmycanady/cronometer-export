# cronometer-export

cronometer-export can export user data from [Cronometer.com](https://cronometer.com). This is intended for personal
use only. All other uses should investigate the Cronometer Premium options.


## Installation

Download the appropriate executable for your operating system from the releases section.

## Basic Usage
> cronometer-export -s -3d -e 0d -u username -p password -o output_file_name.csv

## Exportable Data

* Servings
* Daily Nutrition
* Exercises
* Notes
* Biometrics

## Start/End Time Frames

The start and end times support two different formats. Provide either fromat to the start-at and end-at flags and 
the executable will handle it from there. All times will only utilize up to the day so hour/min/sec can be set to zero.

* [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) 
* Relative Time (-#d or -#m or -#y)

## Output

Ouput is provided either to a file using the -o flag or to stdout. 