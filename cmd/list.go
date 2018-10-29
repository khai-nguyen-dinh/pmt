package cmd

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/khainguyen95/pmt/cmd/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func ListInfo(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List info of application.",
		Long:  `List all info of application.`,
		Run: func(cmd *cobra.Command, args []string) {
			runListInfo(db)
		},
	}
	return cmd
}
func runListInfo(db *gorm.DB) {
	user := u.GetUserActive(db)
	data := u.ListInfo(db, user.ID)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "APPLICATION", "CREATE TIME"})
	table.SetFooter([]string{"", "", ""})
	if len(data) == 0 {
		fmt.Print("No application.Use '$pmt create' to create new application.")
	} else {
		tmp := make([][]string, len(data))
		for i := 0; i < len(data); i += 1 {
			a := []string{
				strconv.Itoa(i),
				data[i].Application,
				data[i].CreatedAt.String(),
			}
			tmp[i] = a
		}
		for _, v := range tmp {
			table.Append(v)
		}

		table.Render()
	}

}
