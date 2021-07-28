# cronometer-export

cronometer-export can export user data from [Cronometer.com](https://cronometer.com). This is intended for personal
use only. All other uses should investigate the Cronometer Premium options.


## Installation

Download the appropriate executable for your operating system from the releases section.

## Basic Usage
> cronometer-export -s -3d -e 0d -u username -p password -o output_file_name.csv

## Help
```
Usage:
  cronometer-export [flags]

Flags:
  -e, --end-at string     The end date in either RFC3339 or -d/w/m/y shorthand.
  -f, --format string     The output format. (raw | json) (Only available on the servings type.) (default "raw")
  -h, --help              help for cronometer-export
  -o, --out-file string   The file to output the data to. If not provided stdout will be used.
  -p, --password string   
  -s, --start-at string   The start date in either RFC3339 or -d/w/m/y shorthand.
  -t, --type string       The type of data to export. (servings | daily-nutrition | exercises | notes | biometrics (default "servings")
  -u, --username string   The username of the user to export data from.

```

## Exportable Data

cronometer-export supports the 5 major export types the web application supports. Each type can be specifed by the -t parameter.

` -t, --type string       The type of data to export. (servings | daily-nutrition | exercises | notes | biometrics (default "servings")`


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

## Output Location

Ouput is provided either to a file using the -o flag or to stdout. 

## Output Format

By default, all output is the raw CSV format provided by the API. Some types support a json output that can be enabled via the -f flag.

`-f, --format string     The output format. (raw | json) (Only available on the servings type.) (default "raw")`
