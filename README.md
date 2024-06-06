# 一、参考文档
1. [bifromq](https://github.com/bifromq/bifromq)
2. [mqtt](https://www.emqx.com/zh/mqtt-guide)
3. [mqttx mqtt-client](https://www.emqx.com/zh/products/mqttx)
4. [eKuiper](https://ekuiper.org/docs/en/latest/getting_started/getting_started.html)
6. [dkron](https://dkron.io/)
# 二、环境准备
1. 安装
 
```shell
docker run -d --name bifromq -p 1883:1883 bifromq/bifromq:latest
```
```shell
docker run -p 9081:9081 -d --name kuiper -e MQTT_SOURCE__DEFAULT__SERVER="tcp://broker.emqx.io:1883" lfedge/ekuiper:latest
```
```shell
docker run --name kuiperManager -d -p 9082:9082 -e DEFAULT_EKUIPER_ENDPOINT="http://192.168.31.72:9081/" emqx/ekuiper-manager:1.8
```
