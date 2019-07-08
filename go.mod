module gosshtool

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190621222207-cc06ce4a13d4
	golang.org/x/net => github.com/golang/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190626221950-04f50cda93cb
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190628021728-85b1a4bcd4e6
)

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/astaxie/beego v1.11.1
	github.com/fatih/color v1.7.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogather/com v0.0.0-20180806101416-e1e2f00d4337
	github.com/kr/fs v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/pkg/sftp v1.10.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
)
