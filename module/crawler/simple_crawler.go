package crawler

import (
	"context"

	"github.com/gocolly/colly/v2"
)

type SimpleCrawlerConstructor struct {
	Url           string
	QuerySelector string
}

func (s SimpleCrawlerConstructor) New() Crawler {
	return &simpleCrawler{SimpleCrawlerConstructor: s, c: colly.NewCollector(colly.Async(false))}
}

type simpleCrawler struct {
	SimpleCrawlerConstructor
	c *colly.Collector
}

func (s *simpleCrawler) Crawl(ctx context.Context) (result <-chan []byte, errChannel <-chan error) {
	ch := make(chan []byte)
	result = ch
	errCh := make(chan error)
	errChannel = errCh

	s.c.OnScraped(func(r *colly.Response) {
		close(ch)
		close(errCh)
	})

	s.c.OnHTML(s.QuerySelector, func(e *colly.HTMLElement) {
		select {
		case <-ctx.Done():
			errCh <- ctx.Err()
			return
		case ch <- []byte(e.Text):
		}
	})

	go func() {
		err := s.c.Visit(s.Url)
		if err != nil {
			errCh <- err
		}
	}()
	return
}
