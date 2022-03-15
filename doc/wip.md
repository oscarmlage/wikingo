# Minimal personal wiki in Go

## Features

* Markdown -> HTML
* Basic HTML (with Water CSS, for example)
  * requires serving static content
* One level pages (e.g. http://host/PageName/)
  * view version (default latest)
    * old version -> restore
    * delete
  * edit
  * list history (changes -> view version)
* Normalize URLs
  * e.g. reidrect page/ -> page

### Stretch goals

1. List pages
  * optionally filter by tag
  * pagination
2. Search
  * query into title (e.g. PageName), tags
  * full text?
  * n results, pagination not needed
3. Auth for protected pages
  * edit, restore
4. Atom feed for updated pages (last n)
5. Detect the page has changed before post
  * show warning

## Stack

* Storage: sqlite with [GORM](https://gorm.io/)
  * easy queries
  * struct -> tables
  * auto-migrations
* HTTP server: [Echo](https://echo.labstack.com/)
  * easy to use
  * routes
  * static content
  * middleware (e.g. auth, remove trailing slash)
* Templates
  * needed? [html/template](https://pkg.go.dev/html/template)
  * not super-easy to use
* Markdown -> HTML: [Blackfriday](https://github.com/russross/blackfriday)

## Routes

| Method | URL                     | Notes                                     |
| ---    | ---                     | ---                                       |
| GET    | /                       | alias for WikiHome                        |
| GET    | /:page(/:version)?      | version is optional, defaults to "latest" |
| GET    | /:page/edit             | protected                                 |
| POST   | /:page/edit             | protected                                 |
| GET    | /:page/history          |                                           |
| POST   | /:page/restore/:version | protected                                 |
| DELETE | /:page/:version         | protected                                 |
| GET    | /list(/:tag)?           | tag is optional, stretch goal             |
| POST   | /search                 | stretch goal                              |
| GET    | /atom.xml               | stretch goal                              |

:page in format CamelCase (starts uppercase)  
:version is numeric  
:tag is string

## Model: Page

* Page name
* version
* change description
* created on
* updated on
* tags (e.g. tag1:tag2:tag3)
* body


(thanks [reidrac](http://github.com/reidrac))
