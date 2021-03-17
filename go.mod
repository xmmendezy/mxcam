module github.com/xmmendezy/mxcam

go 1.16

require (
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	internal/controller v1.0.0
	internal/model v1.0.0
)

replace internal/controller => ./internal/controller

replace internal/model => ./internal/model

replace internal/view => ./internal/view
