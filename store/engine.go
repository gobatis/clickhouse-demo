package store

import (
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/gobatis/gobatis"
	"log"
)

var Engine *gobatis.Engine

func InitEngine() {
	
	db := gobatis.NewDB("clickhouse", "tcp://192.168.1.161:9000?debug=false")
	
	Engine = gobatis.NewEngine(db)
	err := Engine.Init(gobatis.NewBundle("./store/sql"))
	if err != nil {
		log.Panicln("Init error:", err)
	}
	
	if err = Engine.Master().Ping(); err != nil {
		log.Panicln("Ping error:", err)
	}
	
	Engine.UseJsonTag()
	err = BindMapper(Engine)
	if err != nil {
		log.Panicln("Bind mapper error:", err)
	}
	
	err = Engine.Master().Migrate(Migration)
	if err != nil {
		log.Panicln("Migrate error:", err)
	}
}
