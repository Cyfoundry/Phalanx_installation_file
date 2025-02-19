import pandas as pd
from pymongo import MongoClient
 
 
#print(data)
client = MongoClient("mongodb://root:tcfc202303@mongo1:27017/")
db = client["reconnaissance"]  # 替换为您的数据库名
collection = db["vulndata"]  # 替换为您的集合名
 
 
df = pd.read_excel("nasl_data.xlsx")
df = df.where(pd.notnull(df), "")
# 将 DataFrame 转换为字典列表
data = df.to_dict(orient='records')
collection.insert_many(data)
df = pd.read_excel("cve_allitem2.xlsx")
df = df.where(pd.notnull(df), "")
data = df.to_dict(orient='records')
collection.insert_many(data)
 
 
collection = db["scan_tool"]  # 替换为您的集合名
collection.insert_one({
  "step": 3,
  "tool_name": "nmap",
  "rename": "module1"
})
collection.insert_one({
  "step": 4,
  "tool_name": "nuclei",
  "rename": "module2"
})