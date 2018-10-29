package cmd

import (
	"github.com/jinzhu/gorm"
	"github.com/khainguyen95/pmt/cmd/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func ListUser(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show all user.",
		Long:  `Show all user.`,
		Run: func(cmd *cobra.Command, args []string) {
			runList(db)
		},
	}
	return cmd
}

func runList(db *gorm.DB) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "USERNAME", "CREATE TIME", "ACTIVE"})
	table.SetFooter([]string{"", "", "", ""})
	table.SetRowLine(true)
	users := utils.ListAllUser(db)
	tmp := make([][]string, len(users))
	for i := 0; i < len(users); i += 1 {
		status := ""
		if users[i].Active == 1 {
			status = "*"
		}
		a := []string{
			strconv.Itoa(int(users[i].ID)),
			users[i].Username,
			users[i].CreatedAt.String(),
			status,
		}
		tmp[i] = a
	}
	for _, v := range tmp {
		table.Append(v)
	}

	table.Render()
}
