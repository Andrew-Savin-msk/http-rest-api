FROM ubuntu:20.04

# Instaling golang on Ubuntu instead of go image on debian
RUN apt-get update && apt-get install -y wget tar

ENV GOLANG_VERSION 1.22.2
RUN wget https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz -O go.tgz && \
    tar -C /usr/local -xzf go.tgz && \
    rm go.tgz

ENV PATH="/usr/local/go/bin:$PATH"


# Установка go-migrate в Ubuntu
RUN apt-get update && apt-get install -y curl gnupg lsb-release

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - \
  && echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list \
  && apt-get update \
  && apt-get install -y migrate

RUN migrate --version

WORKDIR /go/httprestapi

ENV CONFIG_PATH=configs/apiserver.toml \
  CONFIG_PATH_DOCKER=configs/apiserverDocker.toml \
  PUBLIC_JWT_KEY_PATH=public_key.pem

COPY . .

RUN go mod tidy

# CMD while ! curl -sSf http://db:5432 > /dev/null; do \
#         sleep 1; \
#     done \
#     && echo "Контейнер с тегом 'db' готов к работе. Продолжаем выполнение команды CMD."

# CMD while true; do \
#         if curl -sSf http://db:5432 > /dev/null; then \
#             echo "Контейнер с тегом 'db' готов к работе. Продолжаем выполнение команды CMD." \
#             break; \
#         fi; \
#         sleep 1; \
#     done

# CMD while true; do \
#         if curl -sSf -w "%{http_code}" -o /dev/null http://example.com | grep -q "200"; then \
#             echo "Сервер доступен"; \
#             break; \
#         else \
#             echo "Ошибка: сервер недоступен"; \
#         fi; \
#         sleep 1; \
#     done \
#     && migrate -path migrations -database "postgres://postgres:Sassassa12@db:5432/restapi_dev?sslmode=disable" up \
#     && go run cmd/apiserver/main.go

CMD migrate -path migrations -database "postgres://postgres:Sassassa12@db:5432/restapi_dev?sslmode=disable" up \
  && go run cmd/apiserver/main.go