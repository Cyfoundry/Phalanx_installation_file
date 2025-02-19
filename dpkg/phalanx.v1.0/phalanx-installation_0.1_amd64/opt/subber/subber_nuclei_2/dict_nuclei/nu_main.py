import subprocess
import json
import argparse
import re
import os


def check_target_valid(target_list):
    target_list=target_list.split(",")
    cidr_regex = r'^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$'
    ip_regex = r'^(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$'
    #iprange_regex=r'^(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)[-](?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$'
    domain_regex = re.compile(r'^([A-Za-z0-9]\.|[A-Za-z0-9][A-Za-z0-9-]{0,61}[A-Za-z0-9]\.){1,3}[A-Za-z]{2,6}$')

    target_final=[]
    for target in target_list:
        target=target.strip()
        if re.match(cidr_regex, target):
            target_final.append(target)
            continue

        if re.match(ip_regex, target):
            target_final.append(target)
            continue

        # if re.match(iprange_regex, target):
        #     target_final.append(target)
        #     continue

        if re.match(domain_regex, target):
            target_final.append(target)
            continue

        print(f"{target} is not a valid IP/CIDR/domain,will be neglected ,please check again.")
          
    return target_final
    
def run_command(target_final):
    workdir = os.path.dirname(__file__)
    if os.path.exists(f'{workdir}/nuclei'):#/home/kali/module
    # if os.path.exists('nuclei'):   
        if len(target_final) == 0:
            return 0
        else:    
            target_final=target_final[0]
            #cmd = ['./module/nuclei','-u', target_final, '-es',' info,low','-j']
            #cmd = [f'{workdir}/nuclei','-u', target_final,'-retries','10','-bs','20','-c', '2','-rl','50','-timeout','15','-nmhe','-j']
            cmd = [f'{workdir}/nuclei','-u', target_final,'-j','-timeout','5','-rl','500','-c','200']
            #cmd = ['./module/nuclei.exe','-u', target_final, '-es',' info,low','-j']
            sub_proc = subprocess.Popen(cmd, shell=False,stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            try:
                output, errs = sub_proc.communicate()#output:stdout , errs: stderr   
            except Exception as e:
                sub_proc.kill()
                raise (e)
            else:
                if 0 != sub_proc.returncode:
                    print(errs)
                    print ('Your command is error')
                    exit()
                # print(output.decode('utf8').strip())
                # Response is bytes so decode the output and return
                
                return output.decode('utf8').strip()
        
    else:
        print("There is no nuclei elf , please install")
        exit()

def parse_result(result,target):
    result = result.split("\n")
    # if result[0]=="":
    #     data1=None
    #     return 
    # else:
    data1 = json.loads(result[0])
    
    vulnerability_dict = {"ip": data1.get("ip", target[0]), "ports": []}
    for line in result:
        if line != None and line != "":
            data = json.loads(line)
            if data != None:
                if "ip" not in data:
                    continue
                port_type = data.get("type", "")
                if ( "matched-at" in data):
                    if(data["matched-at"].split(":")[1].isnumeric()):
                        port = data.get("matched-at", "").split(":")[1]
                    if(not data["matched-at"].split(":")[1].isnumeric()):
                        if port_type == "http":
                            port = "80"
                        if port_type == "https":
                            port = "443"
                if port is None:
                    continue
                
                vulnerability = data.get("template-id", "")
                cvss_score = data.get("info", {}).get("classification", {}).get("cvss-score")
                severity = data.get("info", {}).get("severity", "")
                description = data.get("info", {}).get("description", "")
                
                for port_dict in vulnerability_dict["ports"]:
                    if port_dict["port"] == port:
                        new_vuln = {"vulnerability": vulnerability, "cvss_score": cvss_score, "severity": severity, "description":description}
                        if new_vuln not in port_dict["vulnerabilities"]:
                            port_dict["vulnerabilities"].append(new_vuln)
                        #port_dict["vulnerabilities"].append({"vulnerability": vulnerability, "cvss_score": cvss_score, "severity": severity})
                        break
                else:
                    vulnerability_dict["ports"].append({"port": port, "port_type": port_type, "vulnerabilities": [{"vulnerability": vulnerability, "cvss_score": cvss_score, "severity": severity, "description":description}]})

    severity_scores = {'critical': 10, 'high': 7, 'medium': 3, 'low': 1, 'info': 0}
    score = 0
    # 假設結果是按照上面給出的格式解析的
    for line in result:
        if line:
            data = json.loads(line)
            severity = data.get("info", {}).get("severity", "info").lower()  # 預設為info
            score = max(score, severity_scores.get(severity, 0))

    vulnerability_dict["score"] = score  # 添加每個IP的最高評分
    return vulnerability_dict

def main(target, masterid=None):
    target = check_target_valid(target)
    results = ""
    results = run_command(target)
    if results == 0:  # target is not valid
        return None
    elif results == {} or results == "":  # target is valid,but results is not Null
        results = {"ip": target[0], "ports": [], "score": 0}
    else:
        results = parse_result(results, target)

    if masterid:
        results["masterid"] = masterid
    
    return results
