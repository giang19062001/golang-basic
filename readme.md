
# Khởi tạo module Go
go mod init study

# chạy
go run main.go

# build -> tạo ra main.exe có thể chạy ở máy ko có môi trường go vì đây là file đã được go dịch thành ngôn ngữ máy
go build main.go

# cài thư viện
go get <tên thư viện git>
VD:  go get github.com/wagslane/go-tinytime      