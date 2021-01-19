package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// viper unmarshal, Must start with a Upper letter
type getOptions struct {
	CalendarId     string
	SecretFile     string
	Prefix         string
	CacheTokenFile string
	CacheToken     bool
	Strict         bool
	Debug          bool
}

var options getOptions

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go963",
	Short: "google calendar tool",
	Long: `google calendar tool
`,
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
	cobra.OnInitialize(initConfig)

	godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go963.yaml)")
	rootCmd.PersistentFlags().StringVarP(&options.CalendarId, "calendarid", "c", "", "Google Calendar Id.")
	rootCmd.PersistentFlags().StringVarP(&options.SecretFile, "secretfile", "s", "client_secret.json", "OAuth 2.0 Client Secret JSON. Default client_secret.json.")
	rootCmd.PersistentFlags().StringVarP(&options.Prefix, "prefix", "p", "[Go963]", "go963 google calendar event prefix")
	rootCmd.PersistentFlags().StringVarP(&options.CacheTokenFile, "cacheTokenFile", "t", "", "cache the OAuth 2.0 token path. if empty, auto-generate")
	rootCmd.PersistentFlags().BoolVar(&options.CacheToken, "cacheToken", true, "cache the OAuth 2.0 token")
	rootCmd.PersistentFlags().BoolVar(&options.Strict, "strict", false, "if true, control go963 created event only")
	rootCmd.PersistentFlags().BoolVar(&options.Debug, "debug", false, "show HTTP traffic")

	rootCmd.PersistentFlags().MarkHidden("debug")

	// eventJson = flag.String("eventJson", "", "event format json file")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go963" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go963")
	}

	viper.SetEnvPrefix("GO963")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	v := reflect.Indirect(reflect.ValueOf(options))
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		viper.BindEnv(f.Name)
	}

	if err := viper.Unmarshal(&options); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	checkOptions()
	// fmt.Println(options.CalendarId)
	// fmt.Println(options.SecretFile)
}
