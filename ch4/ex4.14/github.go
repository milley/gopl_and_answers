// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package main

import (
	"fmt"
	"time"
)

// IssuesURL issue链接
const IssuesURL = "https://api.github.com/search/issues"

// APIURL api url
const APIURL = "https://api.github.com"

// IssuesSearchResult issue查询结果
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// CacheURL 缓存url
func (i Issue) CacheURL() string {
	return fmt.Sprintf("/issues/%d", i.Number)
}

// User 用户
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//!-
