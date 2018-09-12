# Api blog example

## How to start:

#### Init system:
```sh
export GOPROJECT=${GOPATH}/src/github.com/yakud/apiblog-example

mkdir -p ${GOPROJECT}
git clone https://github.com/yakud/apiblog-example.git ${GOPROJECT}
cd ${GOPROJECT}

# Обновляем go окружение
make install

# Поднимаем docker containers
make up

# Запукаем go-шное приложение
make run
```