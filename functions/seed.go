package functions

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"sync/atomic"
)

func SeedData(
	database *mongo.Database,
	collection *mongo.Collection,
	count int,
	generate func() Entity) error {

	if err := collection.Drop(context.Background()); err != nil {
		return fmt.Errorf("failed to drop collection: %v", err)
	}

	var wg sync.WaitGroup
	data := make([]interface{}, count)
	returnsData := make([]interface{}, 0)

	bufferSize := 50
	numWorkers := 10
	taskCh := make(chan int, bufferSize)

	returnsDataCh := make(chan interface{}, bufferSize)
	if collection.Name() == "companies" {
		go func() {
			for rData := range returnsDataCh {
				returnsData = append(returnsData, rData)
			}
		}()
	}

	for i := 0; i < numWorkers; i++ {
		go func() {
			for idx := range taskCh {
				entity := generate()
				data[idx] = entity

				if collection.Name() == "companies" {
					generateReturn := GenerateReturn(entity)
					for j := 0; j < ReturnCount; j++ {
						returnsDataCh <- generateReturn()
					}
				}
				wg.Done()
			}
		}()
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		taskCh <- i
	}

	wg.Wait()
	close(taskCh)
	close(returnsDataCh)

	if _, err := collection.InsertMany(context.Background(), data); err != nil {
		return err
	}

	if collection.Name() == "companies" && len(returnsData) > 0 {
		returnsCollection := database.Collection("returns")
		if err := returnsCollection.Drop(context.Background()); err != nil {
			return fmt.Errorf("failed to drop returns collection: %v", err)
		}
		insertDataInChunks(database, "returns", returnsData, numWorkers)
	}

	return nil
}

func insertDataInChunks(database *mongo.Database, collectionName string, data []interface{}, numWorkers int) {
	chunkSize := 10000
	dataSize := len(data)
	numChunks := (dataSize + chunkSize - 1) / chunkSize

	var chunkWg sync.WaitGroup
	chunkWg.Add(numChunks)

	var chunkIndex int32

	for i := 0; i < numWorkers; i++ {
		go func() {
			for {
				chunkIdx := atomic.AddInt32(&chunkIndex, 1) - 1
				if int(chunkIdx) >= numChunks {
					break
				}
				start := int(chunkIdx) * chunkSize
				end := start + chunkSize
				if end > dataSize {
					end = dataSize
				}
				if _, err := database.Collection(collectionName).InsertMany(context.Background(), data[start:end]); err != nil {
					panic(err)
				}
				chunkWg.Done()
			}
		}()
	}

	chunkWg.Wait()
}
