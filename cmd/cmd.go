package cmd

import "github.com/khainguyen95/pmt/cmd/utils"

func init() {
	db := utils.InitDB()
	rootCmd.AddCommand(
		VersionPmt(),
		ListInfo(db),
		DetailInfo(db),
		CreateInfo(db),
		DeleteInfo(db),
		ConfigApp(db))
}
