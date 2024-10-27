# Build stage
FROM golang:1.23 as build

WORKDIR /app

# 必要なパッケージをインストール
RUN apt-get update -y && \
    apt-get install -y ffmpeg

# Goモジュールの依存関係を解決
COPY go.mod go.sum ./
RUN go mod download

# プロジェクトのすべてのファイルをコピーしてビルド
COPY . .
RUN go build -o main ./main.go

# Runtime stage
FROM golang:1.23

WORKDIR /app

# ffmpegのインストール
RUN apt-get update -y && \
    apt-get install -y ffmpeg

COPY --from=build /app/main .

# アプリケーションを実行
CMD ["go", "run", "main.go"]
