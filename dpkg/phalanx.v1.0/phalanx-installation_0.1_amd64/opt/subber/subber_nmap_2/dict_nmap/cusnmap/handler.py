import os
import shlex
import subprocess
import sys
import re
import simplejson as json
import argparse
import asyncio
from xml.dom.minidom import parseString
from xml.etree import ElementTree as ET
from xml.etree.ElementTree import ParseError




from .parser import NmapCommandParser
from .utils import get_nmap_path, user_is_root # def function
from .exceptions import NmapNotInstalledError, NmapXMLParserError, NmapExecutionError

import xml


#OS platform
OS_TYPE = sys.platform

class CNmap(object):
    """
    This nmap class allows us to use the nmap port scanner tool from within python
    by calling nmap3.Nmap()
    """
    # @user_is_root
    def __init__(self, path=None):#代表宣告時會自動執行的函式
        """
        Module initialization

        :param path: Path where nmap is installed on a user system. On linux system it's typically on /usr/bin/nmap.
        """
        self.nmaptool = path if path else get_nmap_path()#找nmap的路徑
        self.default_args = "{nmap}  {outarg}  -  "
        self.maxport = 65535
        self.target = ""
        self.top_ports = dict()
        self.parser = NmapCommandParser(None)
        self.raw_ouput = None

    def speed(self, speed_str):#transfer input argument  "speed"
        if speed_str=="high":
            return "5"

        if speed_str=="middle":
            return "4"

        if speed_str=="low":
            return "3"
    
    def depth(self, depth_str):#transfer input argument "depth" 
        if depth_str=="full":
            message = "-p 1-65535 --script vulners,ssl-ccs-injection,ssl-dh-params,ssl-poodle,smb-vuln-ms17-010"
            return message # --script vuln

        if depth_str=="standard":
            return  "--top-ports 100 --script vulners,ssl-ccs-injection,ssl-dh-params,ssl-poodle,smb-vuln-ms17-010"
            

        if depth_str=="less":
            message = "--top-ports 10 --script vulners,ssl-ccs-injection,ssl-dh-params,ssl-poodle,smb-vuln-ms17-010"
            return message# --script vuln
        
    def check_target_valid(self,target_list):
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
          
        return " ".join(target_final)


    @user_is_root #force use root priviledge
    def default_command(self): #default nmap command
        """
        Returns the default nmap command
        eg nmap -sS -sV -O -oX - 
        """
        if OS_TYPE == 'win32':
            print("We don't support windows platform")
            exit(0)
        #format : used to substitute values into a string
        return self.default_args.format(nmap=self.nmaptool, outarg="-sS -sV -v -O -Pn --min-hostgroup 256  -oX") # --script-timeout 5s --scan-delay 1s


    # Unique method for repetitive tasks - Use of 'target' variable instead of 'host' or 'subnet' - no need to make difference between 2 strings that are used for the same purpose
    def scan_command(self, target, arg, args=None, timeout=None):
        self.target == target

        command_args = "{target}  {default}".format(target=target, default=arg)
        scancommand = self.default_command() + command_args 
        if (args):
            scancommand += " {0}".format(args)

        scan_shlex = shlex.split(scancommand) # ex:scan_shlex= nmap -oX - IP --top-ports 10
        output = self.run_command(scan_shlex, timeout=timeout)# do nmap command
        
        xml_root = self.get_xml_et(content)  #XMLoutput

        return xml_root

    @user_is_root #force use root priviledge
    def scan_top_ports(self , target , speed=None , depth=None , port=None , script=None , timeout=None , masterid=None):
        """
        Perform nmap's top ports scan

        :param: target can be IP or domain
        :param: default is the default top port

        """
        target =  self.check_target_valid(target)
        scan_command = self.default_command() + target
        
        #['192.168.163.109', '192.168.163.104', '192.168.163.229', '192.168.163.101', '192.168.163.231', '192.168.163.119']
        validspeed = ["high","middle","low"]
        validdepth =["full","standard","less"]

        if (speed not in validspeed):
            print("You need input high,middle,low")
            exit(0)
        else:
            speed=self.speed(speed)
            scan_command += " -T{0}".format(speed)

        
        if (depth not in validdepth):
            print("You need input full,standard, less")
            exit(0)
        else:
            depth=self.depth(depth)
            scan_command += " {0}".format(depth)
        
        if(port):
            try:
                port=int(port) #prevent user input is not a numbers
                if ( port > self.maxport):
                    print("Port can not be greater than default 65535")
                    exit(0)
                else:
                    scan_command += " -p {0}".format(port)
            except ValueError:
                print("Invalid input. Please enter a valid number.")
                exit()
        
        script_dir = "/usr/share/nmap/scripts" #check script in nmap script_DB or not

        if os.path.exists(os.path.join(script_dir, script)):
            scan_command += " --script {0}".format(script)

        if os.path.exists(os.path.join(script_dir, (script+".nse"))):
            scan_command += " --script {0}".format(script)
        
        scan_shlex = shlex.split(scan_command) 
        #print(scan_shlex)
        # Run the command and get the output 
        output = self.run_command(scan_shlex, timeout=timeout)
        output=parseString(output)
        output = output.toprettyxml()
        
        # with open( "1", "r") as file:
        #     output=file.read()
            # output=xml.dom.minidom.parseString(output)
            # output = output.toprettyxml()
        # exit()
        
        if not output:
            # Probaby and error was raise
            raise ValueError("Unable to perform requested command")
        
        # Begin parsing the xml response
        
        xml_root = self.get_xml_et(output) #output XML
        self.top_ports = self.parser.filter_top_ports(xml_root) #catch dat from XML file

        master_dict = {}
        if masterid:
            master_dict["masterid"] = masterid
            self.top_ports.append(master_dict)
        return self.top_ports  

    def run_command(self, cmd, timeout=None):
        """
        Runs the nmap command using popen

        @param: cmd--> the command we want run eg /usr/bin/nmap -oX -  nmmapper.com --top-ports 10
        @param: timeout--> command subprocess timeout in seconds.
        """
        if (os.path.exists(self.nmaptool)):
            sub_proc = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            try:
                output, errs = sub_proc.communicate(timeout=timeout)#output:stdout , errs: stderr
            except Exception as e:
                sub_proc.kill()
                raise (e)
            else:
                if 0 != sub_proc.returncode:
                    raise NmapExecutionError('Error during command: "' + ' '.join(cmd) + '"\n\n' + errs.decode('utf8'))

                # Response is bytes so decode the output and return
                return output.decode('utf8').strip()
        else:
            raise NmapNotInstalledError()
            
    def get_xml_et(self, command_output):
        """
        @ return xml ET
        """
        try:
            self.raw_ouput = command_output
            #parses a string  returns the root element object of the XML tree
            return ET.fromstring(command_output)
        except ParseError:
            raise NmapXMLParserError() #"Unable to parse xml output"



class CNmapAsync(CNmap):
    def __init__(self, path=None):
        super(NmapAsync, self).__init__(path=path)
        self.stdout = asyncio.subprocess.PIPE
        self.stderr = asyncio.subprocess.PIPE

    async def run_command(self, cmd, timeout=None):
        if (os.path.exists(self.nmaptool)):            
            process = await asyncio.create_subprocess_shell(cmd,stdout=self.stdout,stderr=self.stderr)
            
            try:
                data, stderr = await process.communicate()
            except Exception as e:
                raise (e)
            else:
                if 0 != process.returncode:
                    raise NmapExecutionError('Error during command: "' + ' '.join(cmd) + '"\n\n' + stderr.decode('utf8'))

                # Response is bytes so decode the output and return
                return data.decode('utf8').strip()
        else:
            raise NmapNotInstalledError()
    
    async def scan_command(self, target, arg, args=None, timeout=None):
        self.target == target

        command_args = "{target}  {default}".format(target=target, default=arg)
        scancommand = self.default_command() + command_args
        if (args):
            scancommand += " {0}".format(args)

        output = await self.run_command(scancommand, timeout=timeout)
        xml_root = self.get_xml_et(output)

        return xml_root
    
    @user_is_root    
    async def scan_top_ports(self, target ,speed=None,depth=None , port=None,script=None,timeout=None , masterid=None):
        """
        Perform nmap's top ports scan

        :param: target can be IP or domain
        :param: default is the default top port

        """
        target =  self.check_target_valid(target)
        scan_command = self.default_command() + target
        validspeed = ["high","middle","low"]
        validdepth =["full","standard","less"]

        if (speed not in validspeed):
            print("You need input high,middle,low")
            exit(0)
        else:
            speed=self.speed(speed)
            scan_command += " -T{0}".format(speed)


        if (depth not in validdepth):
            print("You need input full,half, less")
            exit(0)
        else:
            depth=self.depth(depth)
            scan_command += " {0}".format(depth)

        if(port):
            try:
                port=int(port) #prevent user input is not a numbers
                if ( port > self.maxport):
                    print("Port can not be greater than default 65535")
                    exit(0)
                else:
                    scan_command += " -p {0}".format(port)
            except ValueError:
                print("Invalid input. Please enter a valid number.")
                exit()

        script_dir = "/usr/share/nmap/scripts" #check script in nmap script_DB or not
        if os.path.exists(os.path.join(script_dir, script)):
            scan_command += " --script {0}".format(script)

        if os.path.exists(os.path.join(script_dir, (script+".nse"))):
            scan_command += " --script {0}".format(script)
        
        scan_shlex = shlex.split(scan_command) #list
        print(scan_shlex)
        # Run the command and get the output 
        output = await self.run_command(scan_command, timeout=timeout)
        if not output:
            # Probaby and error was raise
            raise ValueError("Unable to perform requested command")

        # Begin parsing the xml response
        xml_root = self.get_xml_et(output) #output XML
        self.top_ports = self.parser.filter_top_ports(xml_root) #catch dat from XML file
        master_dict = {}
        if masterid:
            master_dict["masterid"] = masterid
            self.top_ports.append(master_dict)
        return self.top_ports  
    


    
