package cron

import (
	"fmt"

	"github.com/argoproj/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/argoproj/argo/v2/cmd/argo/commands/client"
	cronworkflowpkg "github.com/argoproj/argo/v2/pkg/apiclient/cronworkflow"
)

// NewSuspendCommand returns a new instance of an `argo suspend` command
func NewSuspendCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "suspend CRON_WORKFLOW...",
		Short: "suspend zero or more cron workflows",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, apiClient := client.NewAPIClient()
			serviceClient := apiClient.NewCronWorkflowServiceClient()
			namespace := client.Namespace()
			for _, name := range args {
				cronWf, err := serviceClient.SuspendCronWorkflow(ctx, &cronworkflowpkg.CronWorkflowSuspendRequest{
					Name:      name,
					Namespace: namespace,
				})
				errors.CheckError(err)
				cronWf.Spec.Suspend = true
				fmt.Printf("CronWorkflow '%s' suspended\n", name)
			}
		},
	}
	return command
}