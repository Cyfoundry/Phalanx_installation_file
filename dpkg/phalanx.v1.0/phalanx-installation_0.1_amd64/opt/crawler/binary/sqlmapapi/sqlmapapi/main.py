import argparse
from sqlmapapi import SQLMapAPI
import os
import json

from socket_pro.client import Client
import toml

path = os.path.abspath(__file__)
paths = path.split("/")
path = "/".join(paths[0:len(paths)-1])
default_config = f"{path}/config/config.toml"

parser = argparse.ArgumentParser(description='Scan a target URL using SQLMapAPI')
parser.add_argument('-c', '--config', dest='config', required=False, help='The path of config', default=default_config)
parser.add_argument('-f', '--file', dest='file', required=False, help='The json file contains GET, POST methods.')
parser.add_argument('-d', '--data', dest='data', required=False, help='use json data')
parser.add_argument('-O', '--output', dest='output', required=False, help='OUTPUT [Output file]')

args = parser.parse_args()
config = args.config
json_file = args.file
string = args.data
toml_dict = toml.load(config)
server_ip = toml_dict['host']
server_port = toml_dict['port']
output = args.output

if __name__ == "__main__":
    ## GET method
    sqlmap = SQLMapAPI(server_ip, server_port)
    if json_file:
        if not os.path.isfile(json_file):
            print("No data from SQLidetector.")
        valid_get = sqlmap.detector(json_file, 'get')
        payload_post, valid_post = sqlmap.detector(json_file, 'post')
        scan_result = sqlmap.scan_all(valid_get, valid_post, payload_post, server_ip)
        if output:
            sqlmap.output_file(scan_result, output)
        else:
            print(json.dumps(scan_result))

    elif string:
        if not string:
            print("No data from SQLidetector.")
        valid_get = sqlmap.detector_str(string, 'get')
        payload_post, valid_post = sqlmap.detector_str(string, 'post')
        scan_result = sqlmap.scan_all(valid_get, valid_post, payload_post, server_ip)
        if output:
            sqlmap.output_file(scan_result, output)
        else: 
            print(json.dumps(scan_result))


