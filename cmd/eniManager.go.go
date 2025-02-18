package cmd

import (
	"log"

	"github.com/yoonhyunwoo/ibm-vpc-cni-k8s/pkg/eni"

	"github.com/spf13/cobra"
)

func ENIManagerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "eni-manager",
		Short: "Manage Elastic Network Interfaces (ENI)",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Starting ENI manager...")
			eni.ManageENI()
		},
	}
}
