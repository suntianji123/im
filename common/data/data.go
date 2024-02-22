package data

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/im/common/conf"
	"github.com/im/common/data/ent"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var DataM *Data

type Data struct {
	db     *ent.Client
	rdb    *redis.Client
	config *conf.DataConfig
}

func NewData(config *conf.DataConfig) *Data {
	DataM = &Data{
		config: config,
	}
	return DataM
}

func (d *Data) Init() error {

	drv, err := sql.Open(
		d.config.DbConfig.Driver,
		d.config.DbConfig.Source,
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
		Addr:         d.config.RedisConfig.Addr,
		DB:           d.config.RedisConfig.DB,
		DialTimeout:  time.Duration(d.config.RedisConfig.DialTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(d.config.RedisConfig.WriteTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(d.config.RedisConfig.ReadTimeout) * time.Millisecond,
	})
	rdb.AddHook(redisotel.TracingHook{})
	d.db = client
	d.rdb = rdb
	return nil
}

func (p *Data) Close() error {
	err := p.db.Close()
	if err != nil {
		logger.Errorf("Data close DB failed:%v", err)
	}
	err = p.rdb.Close()
	if err != nil {
		logger.Errorf("Data close redis failed:%v", err)
	}
	return err
}

func (d *Data) GetDBClient() *ent.Client {
	return d.db
}

func (d *Data) GetRedisClient() *redis.Client {
	return d.rdb
}
