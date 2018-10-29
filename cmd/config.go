package cmd

import (
	"github.com/jinzhu/gorm"
	config "github.com/khainguyen95/pmt/cmd/config"
	"github.com/spf13/cobra"
)

func ConfigApp(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Config User for application.",
		Long:  `Config User for application.`,
	}
	cmd.AddCommand(config.CreateUser(db), config.ChangeUser(db), config.ListUser(db))
	return cmd
}
