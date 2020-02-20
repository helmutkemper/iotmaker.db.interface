package iotmaker_db_interface

type DbFunctionsInterface interface {
	Connect(connection ...interface{}) error
	Disconnect() error
	Find(collection, query interface{}, pointerToResult *[]map[string]interface{}) error
	Count(collection, query interface{}) (error, int64)
	Insert(collection, data interface{}) error
}
