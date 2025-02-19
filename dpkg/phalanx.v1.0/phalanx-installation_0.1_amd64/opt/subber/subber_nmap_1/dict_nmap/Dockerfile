# 使用 Python 官方映像作為基礎映像
FROM python:3.11.2

# 將工作目錄設置為項目目錄
WORKDIR .

# 複製所有 Python 檔案到容器中的工作目錄
ADD . .

# 安裝依賴庫
RUN pip install -r requirements.txt
RUN apt-get update && apt-get install -y nmap


# 定義容器啟動時運行的命令
ENTRYPOINT ["python3", "tes1.py"]