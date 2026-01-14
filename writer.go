package mongox

import (
	"context"
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/mongo"
)

// NewWriter 创建MongoDB日志写入器
func NewWriter(source, collection string) io.Writer {
	if Initialized() {
		client := GetClient(source)
		database := client.GetConfig().Database
		writer := Writer{}
		writer.collection = client.GetClient().Database(database).Collection(collection)
		return &writer
	}
	return nil
}

// Writer 日志写入
type Writer struct {
	collection *mongo.Collection
}

func (w *Writer) Write(bytes []byte) (int, error) {
	go func() {
		// 异步写入
		var doc interface{}
		if err := json.Unmarshal(bytes, &doc); err == nil {
			_, _ = w.collection.InsertOne(context.Background(), doc)
		}
	}()
	return 0, nil
}
