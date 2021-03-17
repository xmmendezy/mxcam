module controller

go 1.16

require (
	github.com/olekukonko/tablewriter v0.0.5
	gorm.io/gorm v1.21.3
	internal/model v1.0.0
	internal/view v1.0.0
)

replace internal/model => ../model

replace internal/view => ../view
