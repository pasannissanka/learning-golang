/*
Copyright © 2023 Pasan Nissanka pasannissanka@outlook.com
*/
package hnapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

type Post struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

func (p Post) String() string {
	t := color.New(color.FgCyan)
	a := color.New(color.FgGreen, color.Underline)
	b := color.New(color.FgYellow, color.Bold)

	return fmt.Sprintf("%s: %s %s: %s\n", b.Sprint("Title: "), t.Sprint(p.Title), b.Sprint("URL: "), a.Sprint(p.URL))
}

func GetItems() ([]int, error) {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var topItems []int
		body, err := io.ReadAll(resp.Body)
		json.Unmarshal(body, &topItems)

		if err != nil {
			return nil, err
		}

		return topItems, nil
	}

	return nil, errors.New("something went wrong")
}

func GetPost(id int) (*Post, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var post Post
		body, err := io.ReadAll(resp.Body)
		json.Unmarshal(body, &post)

		if err != nil {
			return nil, err
		}

		return &post, nil
	}

	return nil, errors.New("something went wrong")
}

func FetchStories(items []int) ([]Post, error) {
	var posts []Post

	for _, item := range items[:5] {
		post, err := GetPost(item)
		if err != nil {
			return nil, err
		}
		posts = append(posts, *post)
	}

	return posts, nil
}
