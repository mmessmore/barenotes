/*
Copyright © 2022 Mike Messmore <mike@messmore.org>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "messynotes",
	Short: "Wrapper CLI to use hugo as a personal note/todo system",
	Long: `The messynotes hugo theme is designed to be a minimalistic
system for maintaining personal notes and todo items.

This is a wrapper cli for providing simple access to maintain and use it
in this fashion.

This will try its best to choose the text editor of your choice.  The order
of precidence (first set wins): command line argument, config file,
$VISUAL, $EDITOR, editor command (for update-alternatives), and vi.  If
none are found, commands like 'new' will fail and you will need to
specify in the config or on the command line.

For the browser it will walk command line argument, config file, the 'open'
command, and then the 'xdg-open' command.  If none work, it will fail.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.messynotes.yaml)")
	rootCmd.PersistentFlags().StringP("root", "r", ".", "Root of hugo repository")
	rootCmd.PersistentFlags().StringP("editor", "e", "", "Text editor binary")
	rootCmd.PersistentFlags().StringP("browser", "b", "", "Web browser binary/launcher")
	rootCmd.PersistentFlags().StringP("hugo", "H", "hugo", "Hugo binary")
	rootCmd.PersistentFlags().StringP("git", "g", "git", "Git binary")
	viper.BindPFlags(rootCmd.PersistentFlags())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".messynotes")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
