module web

go 1.12

replace common => ../common

require (
	common v0.0.0-00010101000000-000000000000
	github.com/foolin/gin-template v0.0.0-20190415034731-41efedfb393b
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/ugorji/go v1.1.4 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)
