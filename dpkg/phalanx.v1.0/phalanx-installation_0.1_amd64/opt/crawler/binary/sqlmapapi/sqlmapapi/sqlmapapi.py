import requests
import json
import sys
import time
import os
from socket_pro.client import Client

class SQLMapAPI:
    def __init__(self, server_ip, server_port, cookies=None):
        self.server_ip = server_ip
        self.server_port = server_port
        self.cookies = {"PHPSESSID": cookies}
        self.task_id = None
    def detector(self, file, method=None):
        target = []
        if os.path.isfile(file):
            with open(file, 'r') as f:
                # Load the contents of the file as a JSON array
                data = json.load(f)
            if method == 'get':
                target = []
                for d in data['sqlidetector']:
                    if d['Result'] == True:
                        ## append the urls for SQLMAP
                        target.append(d['Target'])
                return target

            elif method == 'post':
                payload, url = [], []
                for d in data['crawler']:
                    if d['Method'] == 'post':
                        ## append the payload of post method
                        payload.append(d['Result'])
                        url.append(d['Uri'])
                return payload, url
    def detector_str(self, string, method):
        data = json.loads(string)
        if method == 'get':
                target = []
                for d in data['sqlidetector']:
                    if d['Result'] == True:
                        ## append the urls for SQLMAP
                        target.append(d['Target'])
                return target

        elif method == 'post':
            payload, url = [], []
            for d in data['crawler']:
                if d['Method'] == 'post':
                    ## append the payload of post method
                    payload.append(d['Result'])
                    url.append(d['Uri'])
            return payload, url
        #return data['crawler']
    def new_scan(self, link):
        self.target_url = link
        api_url = f"http://{self.server_ip}:{self.server_port}/task/new"
        try:
            response = requests.get(api_url)
            response.raise_for_status()
            self.task_id = response.json()["taskid"]
            #print(f"Scan started with task ID {self.task_id}")
        except requests.exceptions.RequestException as e:
            print("Error new scan:", e)
            exit(1)
        ## set the scan
        api_url = f"http://{self.server_ip}:{self.server_port}/option/{self.task_id}/set"
        params = {"url": self.target_url, "cookie": json.dumps(self.cookies),"threads": 10, "getDbs":True, "getTables": True, "getColumns": True}
        try:
            response = requests.post(api_url, json=params)
            response.raise_for_status()
        except requests.exceptions.RequestException as e:
            print("Error setting scan:", e)
            sys.exit(1)

    def start_scan(self, method, file = None):
        headers = {'content-type': 'application/json'}
        if method == 'post':
            api_url = f"http://{self.server_ip}:{self.server_port}/scan/{self.task_id}/start"
            params = {"cookies": json.dumps(self.cookies), "method": "POST","requestFile": file}
            try:
                response = requests.post(api_url, data=json.dumps(params), headers=headers)
                response.raise_for_status()
            except requests.exceptions.RequestException as e:
                print("Error starting scan:", e)
                exit(1)
        else:
            api_url = f"http://{self.server_ip}:{self.server_port}/scan/{self.task_id}/start"
            params = {"url": self.target_url, "cookies": json.dumps(self.cookies)}
            try:
                response = requests.post(api_url, data=json.dumps(params), headers=headers)
                response.raise_for_status()
            except requests.exceptions.RequestException as e:
                print("Error starting scan:", e)
                exit(1)

    def check_scan_status(self):
        api_url = f"http://{self.server_ip}:{self.server_port}/scan/{self.task_id}/status"
        status = "running"
        while status == "running":
            try:
                response = requests.get(api_url)
                response.raise_for_status()
                status = response.json()["status"]
                time.sleep(5) # wait for 5 seconds before checking again
            except requests.exceptions.RequestException as e:
                print("Error checking scan status:", e)
                exit(1)
        #print(f"Scan finished with status: {status}")

    def retrieve_scan_results(self):
        api_url = f"http://{self.server_ip}:{self.server_port}/scan/{self.task_id}/data?url={self.target_url}"
        try:
            response = requests.get(api_url)
            response.raise_for_status()
            results = json.loads(response.content)
            if results['data'] or results['error']:
                return results
            else:
                return ''
        except requests.exceptions.RequestException as e:
            print("Error retrieving scan results:", e)
            exit(1)
    def get_log(self):
        '''
        check the error for the task
        '''
        api_url = f"http://{self.server_ip}:{self.server_port}/option/{self.task_id}/list"
        #api_url = f"http://{self.server_ip}:{self.server_port}/scan/{self.task_id}/log"
        try:
            response = requests.get(api_url)
            response.raise_for_status()
            results = json.loads(response.content)
            if results:
                return json.dumps(results)
            else:
                print("No results found.")
        except requests.exceptions.RequestException as e:
            print("Error retrieving scan results:", e)
            exit(1)
    def process_result(self, result_json):
        '''
        modify the format of the result from sqlmapapi
        '''
        for element in result_json['data']:
            if isinstance(element['value'], list) & len(element['value']) == 1:
                element['value'] = element['value'][0]
            if 'data' in element['value']:
                if isinstance(element['value']['data'], dict):
                    element['value']['data'] = list(element['value']['data'].values())
                if element['value']['data'] != None:
                    title = []
                    for i in element['value']['data']:
                        title.append(i['title'])
                    element['value']['data'] = title
            if 'ptype' in element['value']:
                element['value'].pop('ptype')
            if 'conf' in element['value']:
                element['value'].pop('conf')
            if 'clause' in element['value']:
                element['value'].pop('clause')
            if 'notes' in element['value']:
                element['value'].pop('notes')
            if element['type'] == 0:
                result_json['target'] = element
            elif element['type'] == 1:
                result_json['payload'] = element
            elif element['type'] == 12:
                result_json['database'] = {"value":{"name": element['value']}}
            elif element['type'] == 13:
                result_json['tables'] = element
            elif element['type'] == 14:
                result_json['columns'] = element
            element.pop('type')
            element.pop('status')
        result_json.pop('data')
        return result_json
    def scan_all(self, valid_get, valid_post, payload_post, server_ip):
        scan_result = []
        file_name = []
        ## params: host ip, host port, buffer size
        ## NOTICE: the port for sending data is different from slqlmap
        client = Client(server_ip, 9898, 4096)
        for payload in payload_post:
            ## write the file to the sqlmapapi server
            name = client.transfer(payload)
            file_name.append(os.path.basename(name).decode("utf-8"))
        # scan with SQLMAP(GET)
        if valid_get:
            for i, link in enumerate(valid_get):
                self.new_scan(link)
                self.start_scan('get')
                self.check_scan_status()
                result = self.retrieve_scan_results()
                ## result[1] : data, result[2]:error
                if result != '':
                    result_new = self.process_result(result)
                    scan_result.append(result_new)
        # # scan with SQLMAP(POST)
        if valid_post and file_name:
            for link, file in zip(valid_post, file_name):
                self.new_scan(link)
                self.start_scan('post', file)
                self.check_scan_status()
                result = self.retrieve_scan_results()
                if result != '':
                    result_new = self.process_result(result)
                    scan_result.append(result_new)
                client.remove_file(file)
        return scan_result


    ##http://192.168.10.188:8092/index.php?blend=1
if __name__ == '__main__':
    server_ip = '127.0.0.1'
    server_port = '8787'
    sqlmap = SQLMapAPI(server_ip, server_port)
    json_file = input("input a json file: ")
    if json_file:
        if not os.path.isfile(json_file):
            print("No data from SQLidetector.")
        valid_get = sqlmap.detector(json_file, 'get')
        payload_post, valid_post = sqlmap.detector(json_file, 'post')
        scan_result = sqlmap.scan_all(valid_get, valid_post, payload_post, server_ip)
        print(json.dumps(scan_result))

