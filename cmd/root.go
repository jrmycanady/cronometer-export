package cmd

import (
	"fmt"
	"github.com/jrmycanady/cronometer-export/export"
	"github.com/spf13/cobra"
	"os"
)

type ExportType string

const (
	ExportTypeServings       ExportType = "servings"
	ExportTypeDailyNutrition ExportType = "daily-nutrition"
	ExportTypeExercises      ExportType = "exercises"
	ExportTypeNotes          ExportType = "notes"
	ExportTypeBiometrics     ExportType = "biometrics"
)

type CLIOpts struct {
	ExportOpts export.Opts
	Type       string

	OutputFile string
}

var cliOpts CLIOpts

const version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "cronometer-export",
	Short: "Exports user data such as nutrition information from Cronometer.\nVersion: " + version,
	Run:   run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringVarP(&cliOpts.Type, "type", "t", "servings", "The type of data to export. (servings | daily-nutrition | exercises | notes | biometrics")
	rootCmd.Flags().StringVarP(&cliOpts.ExportOpts.Start, "start-at", "s", "", "The start date in either RFC3339 or -d/w/m/y shorthand.")
	rootCmd.MarkFlagRequired("start-at")
	rootCmd.Flags().StringVarP(&cliOpts.ExportOpts.End, "end-at", "e", "", "The end date in either RFC3339 or -d/w/m/y shorthand.")
	rootCmd.MarkFlagRequired("end-at")
	rootCmd.Flags().StringVarP(&cliOpts.ExportOpts.Username, "username", "u", "", "The username of the user to export data from.")
	rootCmd.MarkFlagRequired("username")
	rootCmd.Flags().StringVarP(&cliOpts.ExportOpts.Password, "password", "p", "", "")
	rootCmd.MarkFlagRequired("password")
	//rootCmd.Flags().StringVarP(&cliOpts.ExportOpts.Format, "format", "f", "raw", "The output format. (raw | json)")
	//rootCmd.Flags().BoolVarP(&cliOpts.ExportOpts.InternetMagic, "internet-magic", "i", false, "Denotes if the magic values should be pulled from the internet.")
	rootCmd.Flags().StringVarP(&cliOpts.OutputFile, "out-file", "o", "", "The file to output the data to. If not provided stdout will be used.")
}

func run(cmd *cobra.Command, args []string) {

	cliOpts.ExportOpts.Type = export.ExportType(cliOpts.Type)
	data, err := export.Run(cliOpts.ExportOpts)
	if err != nil {
		fmt.Println(err)
		//os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	if cliOpts.OutputFile != "" {
		if err := os.WriteFile(cliOpts.OutputFile, []byte(data), 0660); err != nil {
			//os.Stderr.WriteString(err.Error())
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	fmt.Println(string(data))

}
