/*
Copyright © 2022 NAME HERE <a2110560@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"project/router"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.GET("", func(context *gin.Context) {
			context.String(http.StatusOK, "pong")
		})
		r.GET("/ping", func(context *gin.Context) {
			context.String(http.StatusOK, "pong")
		})
		router.NewLineRoute(r.Group("/line"))
		router.NewUserRoute(r.Group("/users"))
		err := r.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")

}
