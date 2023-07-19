# vdbench2influxdb

Yet another simple program to upload vdbench results into InfluxDB 1.x for Grafana visualisation.

Inspired by https://github.com/xocoru/vdbench2influxdb

The original code is written in PHP;

Translated to Golang by ChatGPT;

Revised by Liouxiao.

## Build

```
GOOS=linux go build .
```

## Usage

1. Install InfluxDB 1.x and create database for metrics upload:
```
# influx 
Visit https://enterprise.influxdata.com to register for updates, InfluxDB server management, and monitoring.
Or run InfluxDB using docker:
```
docker pull influxb:1.8

mkdir data

docker run -d \
      -e INFLUXDB_REPORTING_DISABLED=true \
      -p 8086:8086 \
      -v $PWD/data:/var/lib/influxdb \
      influxdb:1.8 
```
Connected to http://localhost:8086 version 1.8
InfluxDB shell version: 1.8.10
> create database G200;
``` 
2. Install Grafana and define influxdb datasource. 
3. Run vdbench with any options that you like. 
4. Pick ```flatfile.html``` from vdbench output directory.
5. Run ```vdbench2influxdb flatfile.html sometag influxdb_url```  Use tag if you like to group data on Grafana graphs. For example, ```vdbench2influxdb flatfile.html "n/a" "http://127.0.0.1:8086/write?db=G200&precision=ms"```
6. Create or import Grafana dashboard. (For example ```grafana.json```, don't forgot change datasource)
7. Enjoy

![Grafana Demo Screen](https://raw.githubusercontent.com/xocoru/vdbench2influxdb/master/grafana-demo-graph.png)

PS: Sorry i am not a programmer at all ... so this PHP script could be mind blow. but it works. 
