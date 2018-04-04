package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/ssube/togo/client"
)

func init() {
	projectsCmd := &cobra.Command{
		Use:   "projects",
		Short: "list projects",
		Run: func(cmd *cobra.Command, args []string) {
			projects, err := rootClient.GetProjects()
			if err != nil {
				log.Printf("error adding task: %s", err.Error())
			}
			client.PrintProjects(os.Stdout, projects)
		},
	}

	rootCmd.AddCommand(projectsCmd)
}
