package gocdb

import "strconv"

type CouchDB struct {
	Schema       string `default:"http"`
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

func (cdb *CouchDB) NewInstace(databaseName string) CouchDB {
	if cdb != nil && cdb.DatabaseName == databaseName {
		return *cdb
	}

	return CouchDB{
		Schema:       cdb.Schema,
		Host:         cdb.Host,
		Port:         cdb.Port,
		User:         cdb.User,
		Password:     cdb.Password,
		DatabaseName: databaseName,
	}
}

func (cdb *CouchDB) BaseUrl() string {
	baseUrl := cdb.Schema + "://" + cdb.User + ":" + cdb.Password + "@" + cdb.Host + ":" + strconv.Itoa(cdb.Port) + "/" + cdb.DatabaseName
	return baseUrl
}

func (cdb *CouchDB) GenerateViewUrl(designDocName string, viewName string) string {
	viewUrl := cdb.BaseUrl() + "/_design/" + designDocName + "/_view/" + viewName
	return viewUrl
}
