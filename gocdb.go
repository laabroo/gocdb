package gocdb

import (
	"strconv"

	"github.com/laabroo/gocdb/helpers"
)

//declaration couchdb
var lbrCouchDB CouchDB = CouchDB{}

func Init(cdb *CouchDB) {
	lbrCouchDB = *cdb
}

func FindById(docId string) (interface{}, error) {
	cdbURL := lbrCouchDB.BaseUrl() + "/" + docId
	data, err := helpers.HttpGet(cdbURL, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func FindByView(docName string, viewName string, limit int, skip int) (interface{}, error) {
	data, err := FindByViewFullParams(docName, viewName, limit, skip, nil, true, true)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func FindByViewWithKey(docName string, viewName string, limit int, skip int, key interface{}) (interface{}, error) {
	data, err := FindByViewFullParams(docName, viewName, limit, skip, key, true, true)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func FindByViewFullParams(docName string, viewName string, limit int, skip int, key interface{}, descending bool, includeDocs bool) (interface{}, error) {
	cdbURL := lbrCouchDB.GenerateViewUrl(docName, viewName)
	params := map[string]string{
		"limit":        strconv.Itoa(limit),
		"skip":         strconv.Itoa(skip),
		"descending":   strconv.FormatBool(descending),
		"include_docs": strconv.FormatBool(includeDocs),
	}

	if key != nil {
		params["key"] = "\"" + key.(string) + "\""
	}

	data, err := helpers.HttpGet(cdbURL, params)
	if err != nil {
		return nil, err
	}

	return data, nil
}
