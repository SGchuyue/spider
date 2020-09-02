package model

/*
type Model struct {
	ID uint32 `gorm:
}
*/

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	Client      *mongo.Client
	Ctx         context.Context
	Option      *option
	Error       error `json:"error:default nil"`
	ctxTime     uint
	serviceAddr string
}

/*
func MongoInit() (e error) {
    // context是go的标准库包，不是很清楚这个包的作用，文档上面也没有写很清楚，知道的朋友希望指点
    //一下
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    // 连接uri
    uri := "mongodb://dcjt:dcjt123dalgurak@127.0.0.1:27017/admin"
    // 构建mongo连接可选属性配置
    opt := new(options.ClientOptions)
	// 设置最大连接的数量
    opt = opt.SetMaxPoolSize(uint16(10))
    // 设置连接超时时间 5000 毫秒
    du, _ := time.ParseDuration("5000")
    opt = opt.SetConnectTimeout(du)
    // 设置连接的空闲时间 毫秒
    mt, _ := time.ParseDuration("5000")
    opt = opt.SetMaxConnIdleTime(mt)
    // 开启驱动
    MongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri), opt)
    if err != nil {
        e = err
        return
    }
    // 注意，在这一步才开始正式连接mongo
    e = MongoClient.Ping(ctx, readpref.Primary())
    return

}
*/

// 连接数据库和测试
func (mongoDB *MongoDB) NewDBClient(opts ...optFunc) (*MongoDB, error) {
	mongoDB.PrepareDB(opts...)
	mongoDB.Ctx, _ = context.WithTimeout(context.TODO(), time.Duration(mongoDB.ctxTime)*time.Second)
	mongoDB.Client, mongoDB.Error = mongo.Connect(mongoDB.Ctx, options.Client().ApplyURI(mongoDB.serviceAddr))
	if mongoDB.Error != nil {
		return nil, err
	}
	return mongoDB, nil
}
