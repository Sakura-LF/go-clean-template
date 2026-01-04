package postgres

import (
	log2 "log"
	"os"
	"time"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/pkg/logger"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

type Client struct {
	DB *gorm.DB
}

func NewDB(c config.DBConfig, log *logger.Logger) (*Client, error) {

	// 终端打印输入 sql 执行记录
	newLogger := gorm_logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags), // io writer
		gorm_logger.Config{
			SlowThreshold:             300 * time.Millisecond, // 慢查询 SQL 阈值
			Colorful:                  true,                   // 是否启动彩色打印
			IgnoreRecordNotFoundError: false,
			LogLevel:                  gorm_logger.Error, // Log lever
		},
	)
	db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal().Msg("failed to connect database")
	} else {
		log.Info().Msg("mysql connect success")
	}
	return db, nil
}
