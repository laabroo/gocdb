# gocdb
Simple go module for couchdb.

Available function :
1. Find by doc id
2. Find by view
3. Find by key

How to use: 
```
//Don't forget to change this value by your own setup
const (
	CdbSchema         = "http"
	CdbHost           = "127.0.0.1"
	CdbPort           = 5984
	CdbUser           = "YOUR_USER"
	CdbPassword       = "YOUR_PASSWORD"
	CdbDocName        = "YOUR_DOC_NAME"
	CdbIndexViewName  = "YOUR_INDEX_VIEW_NAME"
	CdbSingleViewName = "YOUR_SINGLE_VIEW_NAME"
	CdbKey            = "YOUR_KEY"
	CdbDbNameArticle  = "YOUR_DATABASE_NAME"
	CdbDocId          = "YOUR_DOC_ID"
)

//Declaration couchdb
var couchDB = gocdb.CouchDB{
	Schema:       CdbSchema,
	Host:         CdbHost,
	Port:         CdbPort,
	User:         CdbUser,
	Password:     CdbPassword,
	DatabaseName: CdbDbNameArticle,
}

func main() {
	//initialize couchdb
	gocdb.Init(&couchDB)

	//find data by document id
	data, err := gocdb.FindById(CdbDocId)
	if err != nil {
		log.Panic(err)
	}
	log.Println(data)

	//find data by view name
	data2, err2 := gocdb.FindByView(CdbDocName, CdbIndexViewName, 2, 0)
	if err2 != nil {
		log.Panic(err2)
	}
	log.Println(data2)

	//find data by key
	data3, err3 := gocdb.FindByViewWithKey(CdbDocName, CdbSingleViewName, 1, 0, key)
	if err3 != nil {
		log.Panic(err3)
	}
	log.Println(data3)
}
```



