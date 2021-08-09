# Gobatis ClickHouse 读写 DEMO

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gobatis/clickhouse-demo/store"
	"time"
)

func main() {
	
	var err error
	
	// 初始化 Gobatis 引擎
	store.InitEngine()
	
	tx, err := store.Engine.Master().Begin()
	if err != nil {
		panic(err)
	}
	
	// 插入数据
	err = store.Mapper.InsertLog(tx, &store.Log{
		Id:        7,
		Url:       "/path/to/target",
		Input:     "{}",
		Output:    `{"code": 200}`,
		Error:     "",
		CreatedAt: time.Now(),
	})
	
	if err != nil {
		fmt.Println("数据写入失败", err)
		_ = tx.Rollback()
		return
	} else {
		err = tx.Commit()
		if err != nil {
			fmt.Println("数据提交失败", err)
			return
		}
	}
	
	fmt.Println("数据写入成功")
	
	// 查询数据
	items, err := store.Mapper.QueryLog()
	if err != nil {
		fmt.Println("数据查询失败", err)
		return
	}
	
	fmt.Printf("共有 %d 条数据：\n", len(items))
	for _, v := range items {
		d, _ := json.Marshal(v)
		fmt.Println(string(d))
	}
}
```