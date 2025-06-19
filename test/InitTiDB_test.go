package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"seckillProject/db"
	"testing"
)

func TestInitTiDB(t *testing.T) {
	_, err := gorm.Open(mysql.Open("4FWYBnvsYcngMuY.root:qe8e8rlOtuNnMcDu@tcp(gateway01.ap-southeast-1.prod.aws.tidbcloud.com:4000)/test?tls=true"), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	fmt.Printf("success")
}

func TestConfigEnvTiDB(t *testing.T) {
	_, err := db.OpenDB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
}
