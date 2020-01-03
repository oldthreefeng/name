/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Buildstamp = ""
	Githash    = ""
	Goversion  = ""
	Version    = "1.0.0"
	Author     string
	LastName   string
	Gender     string
	Time       string
	Wuxing     string
	rootCmd    = &cobra.Command{
		Use:   "name ",
		Short: "name 是一个用于孩子起名的工具",
		Long: `基于golang开发的一个用于通过名字的组合，然后去网站上打分，解析HTML网页，获取这个名字的分数，
然后按分数从高到底保存下来，以供选择, 文档查看: https://github.com/oldthreefeng/name`,
		Run: func(cmd *cobra.Command, args []string) {

			Name()
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show versions",
		Run: func(cmd *cobra.Command, args []string) {
			version()
		},
	}
)

func version() {
	fmt.Printf(`name is cli tools to get your baby names from Golang.
run "name -h" get more help, more see https://github.com/oldthreefeng/name
`)
	fmt.Printf("name version :       %s\n", Version)
	fmt.Printf("Git Commit Hash:     %s\n", Githash)
	fmt.Printf("UTC Build Time :     %s\n", Buildstamp)
	fmt.Printf("Go Version:          %s\n", Goversion)
	fmt.Printf("Author :             %s\n", Author)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Author, "author", "", "louis.hong", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&LastName, "lastName", "l", "", "your baby LastName")
	rootCmd.PersistentFlags().StringVarP(&Time, "birth-time", "t", "2020-01-01-12-00", "your baby birth-time year-month-day-hour-minutes")
	rootCmd.PersistentFlags().StringVarP(&Wuxing, "wuxing", "w", "J", "your baby needs wuxing , JMSHT meaning jin(金) mu(木) shui(水) huo(火) tu(土) ")
	rootCmd.PersistentFlags().StringVarP(&Gender, "gender", "g", "1", "your baby gender: 1/2 boy/girl ")
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
