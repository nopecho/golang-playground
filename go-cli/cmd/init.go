package cmd

import (
	"github/nopecho/golang-playground/go-cli/producer"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "$HOME 디렉토리에 환경변수 파일을 생성합니다.",
	Long:  `$HOME 디렉토리에 환경변수 파일을 생성합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		producer.CreateEnv()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
