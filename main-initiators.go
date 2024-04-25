package main

import (
	"sync"

	postgres_main "github.com/pseudoelement/go-server/src/db/postgres"
)

func initAllTables(queries []postgres_main.TableCreator) error {
	var wg sync.WaitGroup
	errors_ch := make(chan error, len(queries))

	for _, q := range queries {
		wg.Add(1)
		go func(query postgres_main.TableCreator) {
			defer wg.Done()
			err := query.CreateTable()
			if err != nil {
				errors_ch <- err
			} else {
				errors_ch <- nil
			}
		}(q)
	}
	wg.Wait()
	close(errors_ch)

	for e := range errors_ch {
		if e != nil {
			return e
		}
	}

	return nil
}
