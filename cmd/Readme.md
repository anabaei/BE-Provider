# API

#### Post & PUT

```javascript
 Book='{"title": "this is", "published": 2001, "pages": 321, "genres": ["fiction", "mystery"], "rating": 1.4}'

curl -X POST -d "$Book" http://localhost:4001/v1/book

// PUT
curl -X PUT -d "$Book" http://localhost:4001/v1/books/78
```



 ###