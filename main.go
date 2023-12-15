package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
	"os"
)

func main() {
	var cmd = &cobra.Command{}
	var dbId string
	var table string
	var help bool
	var configFile string

	cmd.PersistentFlags().StringVarP(&dbId, "database", "d", "", "Your database id in config.yaml")
	cmd.PersistentFlags().StringVarP(&table, "table", "t", "", "Your table name")
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Your config file path, default ./")
	cmd.PersistentFlags().BoolVarP(&help, "help", "h", false, "Show help message")

	_ = cmd.MarkPersistentFlagRequired("database") //不好用
	_ = cmd.MarkPersistentFlagRequired("table")    //不生效

	if err := cmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
	if help || (dbId == "" && table == "") {
		showHelp(cmd)
	}
	if dbId == "" {
		log.Fatal("Please use database or -d set database id")
	}
	if table == "" {
		log.Fatal("Please use table or -t set table")
	}
	c, err := loadConfig(configFile)
	if err != nil {
		log.Fatal("Load config file failed, " + err.Error())
	}
	if c == nil || len(c) == 0 {
		log.Fatal("Config is empty")
	}
	if dbId == "" {
		log.Fatal("Database id is empty")
	}
	if _, ok := c[dbId]; !ok {
		log.Fatal(fmt.Sprintf("Database id %s not exist in config file", dbId))
	}
	dbc := c[dbId]
	if dbc.User == "" || dbc.Database == "" || dbc.Password == "" || dbc.Host == "" {
		fmt.Println("数据库配置信息不完整，请检查host/user/password/database")
		os.Exit(0)
	}
	dbc.Table = table
	db, err := dbCon(dbc)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	structStr, err := genFromTable(dbc, db)
	println("\n" + structStr)
}

func showHelp(cmd *cobra.Command) {
	fmt.Println("Usage: sql2struct [options]")
	fmt.Println("Options:")
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("  --%s, -%s %s (default %v)\n", flag.Name, flag.Shorthand, flag.Usage, flag.DefValue)
	})
	os.Exit(0)
}
