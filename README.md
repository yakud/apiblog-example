# Api blog example

## How to start:

#### Init system:
```sh
export GOPROJECT=$(GOPATH)/src/github.com/yakud/apiblog-example
mkdir -p $(GOPROJECT)
git clone https://github.com/yakud/apiblog-example.git $(GOPROJECT)
cd $(GOPROJECT)
make install
make up
make run
```