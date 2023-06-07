package aoc

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
)

func cacheFilename(year, day int) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.cache/aoc/%v/%v", homeDir, year, day), nil
}

func getFromCache(year, day int) (string, error) {
	filename, err := cacheFilename(year, day)
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var ErrCantMakeCacheDir = errors.New("can't make cache directory")

func saveToCache(year, day int, value string) error {
	filename, err := cacheFilename(year, day)
	if err != nil {
		return err
	}

	dir := filepath.Dir(filename)

	err = os.MkdirAll(dir, 0o755)
	if err != nil {
		return errors.Join(err, ErrCantMakeCacheDir)
	}

	return os.WriteFile(filename, []byte(value), 0o644)
}

var ErrNoSessionCookie = errors.New("no session cookie found")

func getFromWebsite(year, day int) (string, error) {
	session_cookie := os.Getenv("AOC_SESSION")

	if session_cookie == "" {
		fmt.Println("Error: no session cookie found")
		return "", ErrNoSessionCookie
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	client := &http.Client{Jar: jar}
	cookie := http.Cookie{Name: "session", Value: session_cookie}

	urlString := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	url, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	client.Jar.SetCookies(url, []*http.Cookie{&cookie})

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "github.com/tyler569/aoc-go by me@tyler569.com")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Input(year, day int) string {
	v, err := getFromCache(year, day)
	if err != nil && os.IsNotExist(err) {
		v, err = getFromWebsite(year, day)
		if err != nil {
			panic(err)
		}
		err := saveToCache(year, day, v)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to save to cache:", err)
		}
	} else if err != nil {
		panic(err)
	}
	return v
}
