package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var db *string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spinme",
	Short: "Spin and manage common services",
	Long: `SpinMe is a wrapper around docker to run common applications.
  Use this to easily create dependent services such as databases.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	hd, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbFile := filepath.Join(hd, ".spinme")
	db = rootCmd.PersistentFlags().String("db", dbFile, "Database for local storage")
	initDB(*db)
}

func initDB(dbPath string) {
	// See if the file already exists, if not, create it
	db, err := filepath.Abs(dbPath)
	if err != nil {
		fmt.Printf("Could not parse file path: %s", dbPath)
		os.Exit(1)
	}
	if _, err := os.Stat(db); err == nil {
	} else if os.IsNotExist(err) {
		f, err := os.Create(db)
		if err != nil {
			fmt.Printf("Could not create a local store in %s\n", err.Error())
			os.Exit(1)
		}
		defer f.Close()
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
