module beego_admin

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190618222545-ea8f1a30c443
	golang.org/x/net => github.com/golang/net v0.0.0-20190619014844-b5b0513f8c1b
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190620070143-6f217b454f45
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190620154339-431033348dd0
)

require (
	github.com/astaxie/beego v1.11.1
	github.com/go-sql-driver/mysql v1.4.1
)
