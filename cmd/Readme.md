# API

#### Post & PUT

```javascript
 Book='{"title": "this is", "published": 2001, "pages": 321, "genres": ["fiction", "mystery"], "rating": 1.4}'

curl -X POST -d "$Book" http://localhost:4001/v1/book

// PUT
curl -X PUT -d "$Book" http://localhost:4001/v1/books/78
```


 ### Local DB 
 * Connect to postgresql, using a dependency
 ```
 go get github.com/lib/pq"
 // added to go.mod
 ```
 * Define variable config
 ```
 export READINGLIST_DB_DSN="postgres://readinglist:password@localhost/readinglist?sslmode=disable"
 ```
 