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
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.0.1
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/astaxie/beego v1.11.1
	github.com/denisenkom/go-mssqldb v0.0.0-20190910235433-dd40567e5238 // indirect
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.0.1 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)
