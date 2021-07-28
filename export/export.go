package export

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jrmycanady/gocronometer"
	"regexp"
	"strconv"
	"time"
)

type ExportType string

const (
	ExportTypeServings       ExportType = "servings"
	ExportTypeDailyNutrition ExportType = "daily-nutrition"
	ExportTypeExercises      ExportType = "exercises"
	ExportTypeNotes          ExportType = "notes"
	ExportTypeBiometrics     ExportType = "biometrics"
)

// Validate will return true if the value is a valid ExportType value.
func (t ExportType) Validate() bool {
	switch t {
	case ExportTypeBiometrics, ExportTypeExercises, ExportTypeNotes, ExportTypeServings, ExportTypeDailyNutrition:
		return true
	default:
		return false
	}
}

type ExportFormat string

const (
	ExportFormatRaw  ExportFormat = "raw"
	ExportFormatJSON ExportFormat = "json"
)

// Validate will return true if the value is a valid ExportType value.
func (t ExportFormat) Validate() bool {
	switch t {
	case ExportFormatRaw, ExportFormatJSON:
		return true
	default:
		return false
	}
}

// Opts contains all the options needed to export data.
type Opts struct {
	Type     ExportType
	Start    string
	End      string
	Username string
	Password string
	Format   ExportFormat
	//OutputFile string
	//InternetMagic bool

	StartTime time.Time
	EndTime   time.Time
}

// Parse parses the opt and returns if it's a valid opt.
func (o *Opts) Parse() (err error) {
	if o.Type.Validate() == false {
		return fmt.Errorf("invalid type %s", o.Type)
	}

	//if o.Format.Validate() == false {
	//	return fmt.Errorf("invalid format %s", o.Format)
	//}

	startTime, err := parseTime(o.Start)
	if err != nil {
		return fmt.Errorf("invalid start time: %s", err)
	}
	o.StartTime = *startTime

	endTime, err := parseTime(o.End)
	if err != nil {
		return fmt.Errorf("invalid end time: %s", err)
	}
	o.EndTime = *endTime

	if endTime.Sub(*startTime).Microseconds() < 0 {
		fmt.Println(startTime)
		fmt.Println(endTime)
		return fmt.Errorf("the start must be before or the same as the end time")
	}
	o.StartTime = *startTime

	return nil
}

var shortTimeReg = regexp.MustCompile(`^-*(\d+)(y|m|d)`)

// parseTime parses the time string and returns the Time struct. Time may be in
// RFC3339 or -#(d|m|y) format.
func parseTime(s string) (*time.Time, error) {

	// Handling relative format.
	matches := shortTimeReg.FindAllStringSubmatch(s, -1)
	if len(matches) > 0 {
		now := time.Now()
		t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

		d, err := strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse digit of time")
		}
		if d != 0 {
			switch matches[0][2] {
			case "d", "D":

				t = t.AddDate(0, 0, -1*int(d))
			case "m", "M":
				t = t.AddDate(0, -1*int(d), 0)
			case "y", "Y":
				t = t.AddDate(0, 0, -1*int(d))

			default:
				return nil, fmt.Errorf("invalid unit %s", matches[0][2])
			}
		}

		return &t, nil
	}

	// Assuming time format.
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil, fmt.Errorf("invalid time format")
	}
	return &t, nil

}

// Run runs an export.
func Run(opt Opts) (string, error) {
	if err := opt.Parse(); err != nil {
		return "", err
	}

	clientOpts := gocronometer.ClientOptions{
		GWTContentType: "",
		GWTModuleBase:  "",
		GWTPermutation: "",
		GWTHeader:      "",
	}
	client := gocronometer.NewClient(&clientOpts)

	if err := client.Login(context.Background(), opt.Username, opt.Password); err != nil {
		return "", fmt.Errorf("failed to login: %s", err)
	}
	defer client.Logout(context.Background())

	switch opt.Type {
	case ExportTypeBiometrics:
		data, err := client.ExportBiometrics(context.Background(), opt.StartTime, opt.EndTime)
		if err != nil {
			return "", fmt.Errorf("failed to export biometrics: %s", err)
		}
		return data, nil
	case ExportTypeServings:
		if opt.Format == ExportFormatJSON {
			data, err := client.ExportServingsParsed(context.Background(), opt.StartTime, opt.EndTime)
			if err != nil {
				return "", fmt.Errorf("failed to export servings: %s", err)
			}

			jsonStr, err := json.Marshal(data)
			if err != nil {
				return "", fmt.Errorf("mashalling json: %s", err)
			}
			return string(jsonStr), nil
		}
		data, err := client.ExportServings(context.Background(), opt.StartTime, opt.EndTime)
		if err != nil {
			return "", fmt.Errorf("failed to export servings: %s", err)
		}
		return data, nil
	case ExportTypeNotes:
		data, err := client.ExportNotes(context.Background(), opt.StartTime, opt.EndTime)
		if err != nil {
			return "", fmt.Errorf("failed to export notes: %s", err)
		}
		return data, nil
	case ExportTypeDailyNutrition:
		data, err := client.ExportDailyNutrition(context.Background(), opt.StartTime, opt.EndTime)
		if err != nil {
			return "", fmt.Errorf("failed to export daily nutrition: %s", err)
		}
		return data, nil
	case ExportTypeExercises:
		data, err := client.ExportExercises(context.Background(), opt.StartTime, opt.EndTime)
		if err != nil {
			return "", fmt.Errorf("failed to export exercises: %s", err)
		}
		return data, nil
	default:
		return "", fmt.Errorf("unknown export type")
	}

}
