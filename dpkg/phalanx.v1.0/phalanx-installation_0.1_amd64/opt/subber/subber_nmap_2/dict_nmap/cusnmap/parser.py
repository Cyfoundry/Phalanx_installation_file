import csv
import io
import os
import re
import shlex
import subprocess
import sys
from xml.etree import ElementTree as ET

class NmapCommandParser(object):
    """
    Object for parsing the xml results
    
    Each function below will correspond to the parse
    for each nmap command or option.
    """
    def __init__(self, xml_et):
        self.xml_et = xml_et#element
        self.xml_root = None
        self.port_result_array=[]
            
    def filter_top_ports(self, xmlroot):
        """
        Given the xmlroot return the all the ports that are open from
        that tree
        """
        try:
            if xmlroot.findall("host"):
                scanned_host = xmlroot.findall("host")
            elif xmlroot.findall("hosthint"):
                scanned_host = xmlroot.findall("hosthint")
            else:
                raise ValueError("Unable to find host or hosthint in XML.")

            for hosts in scanned_host:
                port_result_dict = {}
                for addess in hosts.findall("address"): 
                    if ((addess.attrib['addrtype']) == "ipv4"):
                        port_result_dict["ip"]=addess.attrib["addr"]
                    if ((addess.attrib['addrtype']) == "mac"):
                        port_result_dict["MACaddress"]=addess.attrib["addr"]
                    else:
                        port_result_dict["MACaddress"]= "unknown"
                      
                
                if (hosts.find("os/osmatch")): #delete the None
                    port_result_dict["os"]=self.parse_os(hosts)
                else:
                    port_result_dict["os"]=[]
                
                if (hosts.find("ports/port")): #delete the None
                    port_result_dict["hostname"]=self.parse_hostname(hosts)
                else:
                    port_result_dict["hostname"]="unknown"
                
                 
                if (hosts.find("ports/port")):
                   port_result_dict["ports"]= self.parse_ports(hosts)

                self.port_result_array.append(port_result_dict)
                
        except Exception as e:
            raise(e)
        finally:
            return self.port_result_array
   
    
    def parse_os(self, os_results):
        """
        parses os results
        """
        
        #os = os_results.find("")
        os_list = []
        if os is not None:
            if os_results.find("os/osmatch"): 
           # for match in os_results.findall("os/osmatch"):
                try:
                    match=os_results.find("os/osmatch")
                    os_list.append(match.find("osclass").find("cpe").text)
                except:
                    os_list.append("unknow")
            return os_list
        else:
            return {}

    def parse_hostname(self, hostname_results):
        hostname_list = []
        for port in hostname_results.findall("ports/port"):
            if (port.find('service') is not None):
                if( "hostname" in port.find("service").attrib ):
                    hostname=port.find("service").attrib["hostname"]
                    hostname_list.append(hostname.strip())
                
        if len(hostname_list) >0:
            return hostname_list[0]  
        else:
            return "unknown"
            
    def parse_ports(self, xml_hosts):
        
        """
        Parse parts from xml
        """
        open_ports_list = []
        for port in xml_hosts.findall("ports/port"): 
            open_ports = {}
            if(port.find('state') is not None):
                if ((port.find("state").attrib["state"]) == "open"):
                    if ((port.attrib['protocol']) == "tcp"):
                        open_ports["port"]=port.attrib['portid']+"/"+port.attrib['protocol']
                        open_ports["state"]=port.find("state").attrib["state"]
        
            if(port.find('service') is not None ) and ("port" in open_ports):
                if ((port.find("state").attrib["state"]) == "open"):
                    if port.find("service").attrib["name"]:
                        open_ports["service"]=port.find("service").attrib["name"]
                    else:
                        open_ports["service"]="unknown"
            

                for cp in port.find("service").findall("cpe"):
                    open_ports["cpe"] =cp.text
                    if not cp.text:
                        open_ports["cpe"]="unknown"
            
                if( "product" in port.find("service").attrib ):
                    open_ports["product"]=port.find("service").attrib["product"]
                else:
                    open_ports["product"]="unknown"

                if("version" in port.find("service").attrib ):
                    open_ports["version"]=port.find("service").attrib["version"]
                else:
                    open_ports["version"]="unknown"    

            if (port.find('service') is None) and ("port" in open_ports):
                open_ports["service"]="unknown"
                open_ports["cpe"]="unknown"
                open_ports["product"]="unknown"
                open_ports["version"]="unknown"  

            vulnlist=[]
            if(port.findall("script") ): 
                for script in port.findall('script'):
                    if (script.text is not None and script.text != "false"):
                        if ((script.attrib["id"]) == "vulners"):       
                            for table in script.findall('table/table'):
                                vulndict={}
                                pattern=r"(?<![\S])CVE-(\d{4})-\d+"
                                if re.search(pattern, table.find('elem[@key="id"]').text):
                                    if int(re.match(pattern, table.find('elem[@key="id"]').text).group(1)) > 2012:
                                        vulndict["vulnerability"]=table.find('elem[@key="id"]').text
                                        cvss_score =table.find('elem[@key="cvss"]').text
                                        vulndict["cvss_score"]=cvss_score
                                        if float(cvss_score) < 4:
                                            vulndict["severity"] = "low"
                                        elif float(cvss_score) >= 4 and float(cvss_score) < 7:
                                            vulndict["severity"] = "medium"
                                        elif float(cvss_score) >= 7 and float(cvss_score) < 9:
                                            vulndict["severity"] = "high"
                                        elif float(cvss_score) >= 9 and float(cvss_score) <= 10:
                                            vulndict["severity"] = "critical"
                                        else:
                                            vulndict["severity"] = "unkown"

                                        vulnlist.append(vulndict)

                        if ((script.attrib["id"]) == "ssl-ccs-injection"):
                                vulndict={}
                                vulndict["vulnerability"] = script.find("table/elem[@key='title']").text
                                vulndict["severity"] =re.search(r'Risk factor: (\w+)', script.attrib["output"]).group(1).strip().lower()
                                vulndict["cvss_score"] ="9.2"
                                vulndict["description"] =script.find("table/table[@key='description']").find("elem").text.strip()
                                vulnlist.append(vulndict)

                        if ((script.attrib["id"]) == "ssl-dh-params") or ((script.attrib["id"]) == "ssl-poodle"):
                            if (script.text is not None and script.text != "false"):
                                vulndict={}
                                vulndict["vulnerability"] = script.find("table/elem[@key='title']").text
                                vulndict["severity"] = "low"
                                vulndict["cvss_score"] ="3.1"
                                vulndict["description"] =script.find("table/table[@key='description']").find("elem").text.strip()
                                vulnlist.append(vulndict)
            
            if(open_ports != {}) and ("port" in open_ports):
                if open_ports["port"] == "445/tcp":
                    for script in xml_hosts.findall("hostscript/script"):
                        v_dict={}
                        s = script.attrib["id"]
                        match = re.search(r"ms\d{2}-\d{3}", s)
                        if match is not None and script.text != "false":
                            v_dict["vulnerability"]=match.group()
                            severity=re.search(r'Risk factor: (\w+)', script.attrib["output"]).group(1).strip().lower()
                            v_dict["severity"]= severity
                            v_dict["cvss_score"] ="unkown"
                            v_dict["description"] = script.find('table/table[@key="description"]').find("elem").text.strip()
                            vulnlist.append(v_dict)
            if vulnlist:
                open_ports["vulnerabilities"] = vulnlist

            if open_ports :
                open_ports_list.append(open_ports)
    
         
        return open_ports_list
