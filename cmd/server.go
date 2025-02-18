package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yoonhyunwoo/ibm-vpc-cni-k8s/pkg/cni"

	"github.com/spf13/cobra"
)

func ServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start the IBM VPC CNI server",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Starting IBM VPC CNI server...")
			http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "OK")
			})
			cni.StartServer() // CNI 플러그인 초기화
			log.Fatal(http.ListenAndServe(":8080", nil))
		},
	}
}
