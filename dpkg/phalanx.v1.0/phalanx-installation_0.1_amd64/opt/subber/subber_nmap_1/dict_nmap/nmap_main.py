# import cusnmap
from cusnmap import CNmap as cnp
from cusnmap import CNmapAsync as cnpasyn
import sys
import json
import argparse
import asyncio
import time

parser = argparse.ArgumentParser()
parser.add_argument("-t","--target ", help="input the target",dest="target", default="127.0.0.1",type=str)
parser.add_argument("-s","--speed ", help="input the scan speed(high,middle,low)",dest="speed", default="high",type=str) #high(T5),middle(T3),low(T1)
parser.add_argument("-d","--depth ", help="input the scan depth(full,standard,less)",dest="depth", default="standard",type=str)#full(65535),standard(top 100),less(top 10)
parser.add_argument("-p","--port ", help="input the port",dest="port", default="",type=str)
parser.add_argument("-m","--masterid ", help="input the masterid",dest="masterid", default="",type=str)
parser.add_argument("--script ", help="input the script",dest="script", default="default",type=str)
parser.add_argument("--async", help="input the script",dest="script", default="default",type=str)
parser.add_argument("--asyncrun", help="Use async run", dest='asyncrun',default=False,action='store_true')
#parser.set_defaults(asyncrun=False)
#parser.add_argument('input_file', type=str, help='input file name')
args = parser.parse_args()

target=args.target 
speed=args.speed
depth=args.depth
port=args.port
script=args.script
asyncrun=args.asyncrun
masterid=args.masterid
# with open(args.input_file, 'r') as f:
#     ips = f.read().split(',')
# for ip in ips:
#     print(ip.strip())
#python script.py ips.txt
if asyncrun:
    nmap = cnpasyn()
else:
    nmap = cnp()

#Default command: nmap {IP/CIDR/domains} -sS -sV -O -oX

if asyncrun:
    results = asyncio.run(nmap.scan_top_ports(target,speed,depth,port,script,5,masterid))#use asyncio function
else:
    results = nmap.scan_top_ports(target,speed,depth, port,script,None,masterid)
end = time.time()
json_formatted_str = json.dumps(results, indent=2)
# json_formatted_str = json.dumps(results, indent=2)
print(json_formatted_str)
# print("執行時間：%f 秒" % (end - start))
# json_str = json.dumps(results, indent=2)
# with open("2.json","w") as f:
#     f.write(json_str)