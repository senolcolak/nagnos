#!/usr/bin/python3
# srassistant is a server for providing tmux screens to system engineers
#
# Copyright 2020 Nucleuss GmbH
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from flask import Flask, request
from flask_restful import Resource, Api


app = Flask(__name__)
api = Api(app)


commandz = []

##-- 1st screen, the list of commands for explanation
@app.route('/srcommands', methods=['GET'])
def get_commands(commands):
    return "ls"
##first route will get the available commands list
## always check if the commands list is updated or not



@app.route('/srcommands/<string:command>', methods=['GET'])
def getcommand(command):
    pass
##second route will give the details how tu use the command





##-- 2nd screen, analyse screen what can be done which problems exits
@app.route('/sranalyse')
def sranalyse():
    return "ubuntu-centos-debian"
#analyse the operating system, version, etc..
#run check_mk.sh to get an overview and check the SRANALYSE DB, which routines could be queried.

@app.route('/sranalyse/<string:package>', methods=['GET'])
def analyse_pkg(package):
    pass

@app.route('/sranalyse/<string:service>', methods=['GET'])
def analyse_srvs(service):
    pass



##3rd screen, disabled for automated management.
@app.route('/srnsrvs', methods=['GET'])
def srnsrvs():
    return "ls"



if __name__ == '__main__':
    app.run(port=5000, debug=True)
