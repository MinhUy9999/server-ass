# Sử dụng hình ảnh Go chính thức
FROM golang:1.22 AS builder

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy các tệp cần thiết vào container
COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Biên dịch ứng dụng thành tệp nhị phân
RUN go build -o myapp .

# Expose cổng server
EXPOSE 3000

# Lệnh khởi chạy ứng dụng
CMD ["./myapp"]
