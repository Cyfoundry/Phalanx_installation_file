#!/bin/bash

# 定義服務的順序
services_to_stop=(
    "nats-server.service"
    "endpoint.service"
    "jobmanager.service"
    "crawler.service"
    "msf.service"
    "nmap_subber@1.service"
    "nmap_subber@2.service"
    "nuclei_subber@1.service"
    "nuclei_subber@2.service"
    "sqlmap_server.service"
    "sqlmap_api.service"
    "sqlmap.service"
    "sqlidetector.service"
    "tunnel.service"
)

# 關閉服務
echo "正在關閉服務..."
for service in "${services_to_stop[@]}"; do
    systemctl stop "$service"
    if [ $? -eq 0 ]; then
        echo "已停止服務: $service"
    else
        echo "無法停止服務: $service"
    fi
    sleep 1  # 等待1秒
done

# 等待更長的時間（10秒）
echo "等待10秒..."
sleep 10

# 定義開啟服務的順序（與關閉相同）
services_to_start=(
    "nats-server.service"
    "endpoint.service"
    "jobmanager.service"
    "crawler.service"
    "msf.service"
    "nmap_subber@1.service"
    "nmap_subber@2.service"
    "nuclei_subber@1.service"
    "nuclei_subber@2.service"
    "sqlmap_server.service"
    "sqlmap_api.service"
    "sqlmap.service"
    "sqlidetector.service"
    "tunnel.service"
)

# 開啟服務
echo "正在開啟服務..."
for service in "${services_to_start[@]}"; do
    systemctl start "$service"
    if [ $? -eq 0 ]; then
        echo "已開啟服務: $service"
    else
        echo "無法開啟服務: $service"
    fi
    sleep 1  # 等待1秒
done

echo "完成"
