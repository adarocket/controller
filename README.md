# Black Rocket, monitoring solution for blockchain nodes

Black Rocket Controller is a central point of Black Rocket solution, it gathers the information form end-point clients (informers) about node state, stores it and provides to other solution components (alerter and monitoring app).

## How to install Controller

### Installing Go
Controller requires Go 1.16 to compile, please refer to the [official documentation](https://go.dev/doc/install) for how to install Go in your system.

### Installing Controller:
```
go get github.com/adarocket/controller 
```

### Сonfiguration 
The next step is to create a configuration file. It is necessary to list in the "nodes" list all informers that need to be connected to the controller.
* ticker - name of the informer.
* uuid - unique id of the informer.

```
cat << EOF > ~/etc/ada-rocket/controller.conf
{
    "server_port":"5300",
    "nodes":[
        {
            "ticker":"informer 1 name",
            "uuid":"informer 1 uuid"
        },
          ...
        {
            "ticker":"informer N name",
            "uuid":"informer N uuid"
        }
    ]
}
EOF
```

### Setup ufw
If you have a firewall installed, you need to open port 5300. To do this, you need to run the following command:

```
sudo ufw allow 5300
```

## How to support

If you want to support this project, you can do it directly by sending coins to us, our wallets:

#### Ethereum or Binance Smart chain: 0x1af0637A6131f29389c2e68517D61bF5e2655a57
#### Cardano: addr1qyqa4x78s0l9vusy3kw4772najwzer2s0pk9l8t4hfrushsm50kqnkgssjre0nysnwz9uc20gsanmqsnwxdnxj4w7zfswl9fse
