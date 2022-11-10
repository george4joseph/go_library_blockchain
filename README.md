# Library Blockchain 

A simple library blockchain  in Golang



## APIs

### GET http://localhost:3000
```
[
 {
  "Pos": 1,
  "Data": {
   "book_id": "",
   "user": "",
   "checkout_date": "",
   "is_genesis": true
  },
  "Timestamp": "2022-11-10 11:45:11.396579301 +0530 IST m=+0.000332246",
  "Hash": "bdc787bb3fc6ff744df1d9252f46d927fff04a4a532585490d9cf5c1d4121347",
  "PreHash": ""
 }
]
```

### POST http://localhost:3000/book
```
body:
{
    "title": "Go Up and Down",
    "author" : "Fernando Leo",
    "isbn": "324235235",
    "publish_date" : "2022-10-11"
}

Response:

{
 "id": "1551c8a3d2eebe5d626221cd7d60a28d",
 "title": "Go Up and Down",
 "author": "Fernando Leo",
 "isbn": "324235235",
 "publish_date": "2022-10-11"
}
```

### POST http://localhost:3000/issue
```
body:
{
    "book_id": "1551c8a3d2eebe5d626221cd7d60a28d",
    "user": "john",
    "checkout_date" : "2022-11-8"
}

Response:
{
 "book_id": "1551c8a3d2eebe5d626221cd7d60a28d",
 "user": "john",
 "checkout_date": "2022-11-8",
 "is_genesis": false
}
```
