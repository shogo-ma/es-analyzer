package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	elastic "gopkg.in/olivere/elastic.v5"
)

func main() {
	var (
		host     string
		analyzer string
		query    string
	)

	// host flag
	flag.StringVar(&host, "host", "http://localhost:9200", "host url")
	flag.StringVar(&host, "h", "http://localhost:9200", "host url")

	// analyzer flag
	flag.StringVar(&analyzer, "analyzer", "standard", "analyzer")
	flag.StringVar(&analyzer, "a", "standard", "analyzer")

	// query flag
	flag.StringVar(&query, "query", "", "query")
	flag.StringVar(&query, "q", "", "query")
	flag.Parse()

	if host == "" {
		log.Println("Error: host not found")
		os.Exit(1)
	}

	client, err := elastic.NewClient(
		elastic.SetURL(host),
	)

	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	res, err := client.IndexAnalyze().
		Analyzer(analyzer).
		Text(query).
		Do(context.Background())

	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("token\tstart_offset\tend_offset\ttype\tposition")
	for _, t := range res.Tokens {
		fmt.Printf(
			"%s\t%d\t%d\t%s\t%d\n",
			t.Token,
			t.StartOffset,
			t.EndOffset,
			t.Type,
			t.Position,
		)
	}
}
