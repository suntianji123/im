package data

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/im/common/conf"
	"github.com/im/common/data/ent"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var Data = &data{}

type data struct {
	Db  *ent.Client
	Rdb *redis.Client
}

func (d *data) Init() error {

	dataConfig := &conf.DataConfig{
		DbConfig:    &conf.DatabaseConfig{},
		RedisConfig: &conf.RedisConfig{},
	}
	err := mapstructure.Decode(config.GetRootConfig().Custom.GetDefineValue("data", map[string]interface{}{}), &dataConfig)
	if err != nil {
		logger.Errorf("反序列化配置文件失败:%v", err)
	}
	drv, err := sql.Open(
		dataConfig.DbConfig.Driver,
		dataConfig.DbConfig.Source,
	)
	if err != nil {
		logger.Errorf("打开数据库失败:%v", err)
		return err
	}

	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		logger.Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		logger.Errorf("failed opening connection to sqlite: %v", err)
		return err
	}
	//Run the auto migration tool.
	//if err = client.Schema.Create(context.Background()); err != nil {
	//	logger.Errorf("failed creating schema resources: %v", err)
	//	return err
	//}

	rdb := redis.NewClient(&redis.Options{
		Addr:         dataConfig.RedisConfig.Addr,
		DB:           dataConfig.RedisConfig.DB,
		DialTimeout:  time.Duration(dataConfig.RedisConfig.DialTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(dataConfig.RedisConfig.WriteTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(dataConfig.RedisConfig.ReadTimeout) * time.Millisecond,
	})
	rdb.AddHook(redisotel.TracingHook{})
	d.Db = client
	d.Rdb = rdb
	return nil
}
