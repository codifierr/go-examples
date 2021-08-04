package main

import (
	"context"
	"fmt"
	"os"

	"log"

	"gopkg.in/olivere/elastic.v5"
)

var esClient *elastic.Client

func main() {
	fmt.Println("Hello elastic search")
	StartESClient()

	result1, _ := searchIndices("stats-alias")
	if len(result1) > 0 {
		fmt.Println(result1)
	} else {
		fmt.Println("indice not found")
	}

}

func StartESClient() error {
	answer := make([]string, 1)
	answer[0] = fmt.Sprintf("http://%s:%d", "localhost", 9200)
	var err error
	esClient, _ = elastic.NewClient(elastic.SetURL(answer...), elastic.SetSniff(false), elastic.SetTraceLog(log.New(os.Stdout, "ES", log.LstdFlags)))
	if elastic.IsConnErr(err) {
		return err
	}
	return nil
}

func searchIndices(alias string) ([]string, error) {
	res, err := esClient.Aliases().Index(alias).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res.IndicesByAlias(alias), nil
}
