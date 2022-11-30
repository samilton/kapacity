package commands

import (
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	conf "github.com/samilton/kapacity/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbose bool
	cfgFile string
	config  conf.TopicConfig

	rootCmd = &cobra.Command{
		Use:   "kapacity",
		Short: "Manage Kafka clusters",
		Long:  `Kapacity is a tool to manage Kafka clusters including the management of topics, the viewing consumer group status, and the management of offsets.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enabled more detailed logging")
}

func initConfig() {
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Error().Err(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".kapacity")
		viper.SetConfigType("yaml")
		viper.SetEnvPrefix("KPC")
		viper.BindEnv("ENV")
	}
	viper.AutomaticEnv()

	if verbose == false {
		verbose = viper.GetBool("verbose")
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

}
