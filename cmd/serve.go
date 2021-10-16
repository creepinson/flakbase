package cmd

import (
	"github.com/spf13/cobra"

	"github.com/creepinson/flakbase/pkg/net"
)

var (
	flagHost  string
	flagMongo string
	flagRule  string
)

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start Flakbase server",
	Args:  cobra.NoArgs,
	Run:   serve,
}

func init() {
	cmdServe.Flags().StringVarP(&flagHost, "host", "", "0.0.0.0:9527", "host name to serve")
	cmdServe.Flags().StringVarP(&flagMongo, "mongo", "m", "", "mongodb config file")
	cmdServe.Flags().StringVarP(&flagRule, "rule", "", "", "security rule json file")
}

func serve(cmd *cobra.Command, args []string) {
	net.Run(&net.Config{
		Host:  flagHost,
		Rule:  flagRule,
		Mongo: flagMongo,
	})
}
