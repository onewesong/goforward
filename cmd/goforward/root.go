/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/onewesong/goforward/internal/config"
	"github.com/onewesong/goforward/internal/constants"
	"github.com/onewesong/goforward/internal/manager"
	"github.com/onewesong/goforward/internal/models"
	"github.com/onewesong/goforward/internal/pkg/log"
	"github.com/onewesong/goforward/internal/router"
	"github.com/spf13/cobra"
)

var (
	cfgFile         string
	showVersion     bool
	logWay          string
	logFile         string
	logLevel        string
	logMaxDays      int64
	disableLogColor bool

	listenAddr  string
	forwardLink string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file of frps")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "version of frps")
	rootCmd.PersistentFlags().StringVarP(&logWay, "log_way", "", "console", "log file")
	rootCmd.PersistentFlags().StringVarP(&logFile, "log_file", "", "console", "log file")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log_level", "", "info", "log level")
	rootCmd.PersistentFlags().Int64VarP(&logMaxDays, "log_max_days", "", 3, "log max days")
	rootCmd.PersistentFlags().BoolVarP(&disableLogColor, "disable_log_color", "", false, "disable log color in console")
	rootCmd.PersistentFlags().StringVarP(&listenAddr, "listen_addr", "l", "0.0.0.0:5668", "forward api listen addr")
	rootCmd.PersistentFlags().StringVarP(&forwardLink, "forward_link", "f", "", "forward link, e.g. -f=127.0.0.1:12345->[2400:3200::1]:443,127.0.0.1:12346->[2400:3200:baba::1]:443")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goforward",
	Short: "forward ipv4/6:port to ipv4/6:port",

	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Println(constants.VERSION)
			return nil
		}
		var cfg *config.Config
		var err error
		if cfgFile != "" {
			cfg, err = config.LoadCfg(cfgFile)
		} else {
			cfg, err = parseCfgFromCmd()
		}
		if err != nil {
			return err
		}
		err = runServer(*cfg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runServer(cfg config.Config) (err error) {
	log.InitLog(cfg.LogWay, cfg.LogFile, cfg.LogLevel, cfg.LogMaxDays, cfg.DisableLogColor)

	if cfgFile != "" {
		log.Info("goforward uses config file: %s", cfgFile)
	} else {
		log.Info("goforward uses command line arguments for config")
	}

	r := router.SetupRouter()
	go func() {
		err = r.Run(cfg.ListenAddr)
		if err != nil {
			log.Fatal("start api listen err: %s", err)
		}
	}()

	m := manager.GetInstance()
	m.AddForward(cfg.ForwardLinks, false)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("agent get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			log.Info("agent exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}

func parseCfgFromCmd() (*config.Config, error) {
	cfg := &config.Config{}
	forwardLinks, err := models.NewForwardLinks(forwardLink)
	if err != nil {
		return nil, err
	}
	cfg.ListenAddr = listenAddr
	cfg.ForwardLinks = forwardLinks
	cfg.LogWay = logWay
	cfg.LogFile = logFile
	cfg.LogLevel = logLevel
	cfg.LogMaxDays = logMaxDays
	cfg.DisableLogColor = disableLogColor
	return cfg, nil
}
