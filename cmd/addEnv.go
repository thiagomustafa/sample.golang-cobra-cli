/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var env string
var requiredServersName = map[string]string{
	"BILLING":             "",
	"BILLING_INTEGRATION": "",
}

// addEnvCmd represents the addEnv command
var addEnvCmd = &cobra.Command{
	Use:   "addEnv",
	Short: "Set up environment variables",
	Long:  `Set up the required environment variables for the CLI to work properly`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addEnv called")
		fmt.Println("Environment name: ", env)
		askForServerNames()
		fmt.Println("Server names: ", requiredServersName)
		storeEnvs()
	},
}

func init() {
	rootCmd.AddCommand(addEnvCmd)

	// Here you will define your flags and configuration settings.

	addEnvCmd.Flags().StringVar(&env, "env", "", "Environment name")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addEnvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addEnvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func askForServerNames() {
	reader := bufio.NewReader(os.Stdin)
	for key := range requiredServersName {
		fmt.Printf("Enter the server name for %s: ", key)
		input, _ := reader.ReadString('\n')
		requiredServersName[key] = strings.TrimSpace(input)
	}
}

func storeEnvs() {
	// WIP: This is not working yet

	// viper.SetConfigName(".my-cli-env")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")

	// for key, value := range requiredServersName {
	// 	viper.Set(key, value)
	// }

	// err := viper.WriteConfigAs(fmt.Sprintf("%s.env", env))
	// if err != nil {
	// 	fmt.Printf("Error writing config file: %v\n", err)
	// }

	// viper.SetConfigName(".my-cli")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("$HOME")
	// viper.AddConfigPath(".")

	// if err := viper.ReadInConfig(); err != nil {
	//     if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	//         // Config file not found; ignore error if desired
	//     } else {
	//         // Config file was found but another error was produced
	//         fmt.Println("Error reading config file:", err)
	//         return
	//     }
	// }

	// viper.Set(fmt.Sprintf("environments.%s.server", env), server)

	// configFile := filepath.Join(os.Getenv("HOME"), ".my-cli.yaml")
	// if err := viper.WriteConfigAs(configFile); err != nil {
	//     fmt.Println("Error writing config file:", err)
	// }
}
