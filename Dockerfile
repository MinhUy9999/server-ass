# Sử dụng hình ảnh Go chính thức
FROM golang:1.22 AS builder

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy các tệp cần thiết vào container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Biên dịch ứng dụng thành tệp nhị phân
RUN go build -o master .

# Tạo một lightweight container để chạy ứng dụng
FROM debian:bullseye-slim


# Thư mục làm việc
WORKDIR /root/

# Copy tệp nhị phân từ builder sang container này
COPY --from=builder /app/master .

# Expose cổng server
EXPOSE 3000

# Lệnh khởi chạy ứng dụng
CMD ["./main"]
