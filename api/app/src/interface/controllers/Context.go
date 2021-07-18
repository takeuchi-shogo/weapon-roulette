package controllers

type Context interface {
	Query(key string) string
	JSON(code int, obj interface{})
}
