package wget

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"os"
	"time"
)

var (
	defaultOutput  = "index.html"
	defaultTimeout = 5 * time.Second
)

type Options struct {
	URL     string
	Output  string
	Timeout time.Duration
}

func Exec(args []string) error {
	options := Options{}

	flagSet := flag.NewFlagSet("", flag.ExitOnError)
	flagSet.StringVar(&options.URL, "u", "", "url")
	flagSet.StringVar(&options.Output, "o", defaultOutput, "output")
	flagSet.DurationVar(&options.Timeout, "t", defaultTimeout, "timeout")

	err := flagSet.Parse(args)
	if err != nil {
		return fmt.Errorf("failed to parse flags: %s", err)
	}

	if options.URL == "" {
		return fmt.Errorf("url is required")
	}

	url, err := urlpkg.Parse(options.URL)
	if err != nil {
		return err
	}

	if url.Scheme == "" {
		url.Scheme = "https"
	}

	ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
	defer cancel()

	data, err := fetch(ctx, url.String())
	if err != nil {
		return err
	}

	file, err := os.Create(options.Output)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(data))
	if err != nil {
		return err
	}

	return nil
}

func fetch(ctx context.Context, url string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
