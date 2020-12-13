package backend

import "reflect"

func WithPrefix(route string) string {
	return HTTP_PREFIX + route
}

//Convert
//
//  convert from pttbbs to backend using reflect
//  dataPttbbs is by ptr. (*Type)
//  dataBackend is by ptr of ptr. (**Type)
//
//Ex:
//   dataPttbbs := &api.LoginResult{} //from go-pttbbs/api
//   dataBackend := &LoginResult{}    //from backend
//   Convert(dataPttbbs, &dataBackend)
//   reflect.DeepEqual((*LoginResult)(dataPttbbs), dataBackend)
func Convert(dataPttbbs interface{}, dataBackend interface{}) {
	valueBackend := reflect.ValueOf(dataBackend)
	typeBackend := valueBackend.Elem().Type()
	valuePttbbs := reflect.ValueOf(dataPttbbs)
	converted := valuePttbbs.Convert(typeBackend)
	valueBackend.Elem().Set(converted)
}
