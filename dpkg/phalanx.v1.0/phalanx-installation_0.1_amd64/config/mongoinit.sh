#!/bin/bash
mongo -u root --authenticationDatabase admin -p tcfc202303 <<EOF
rs.initiate();
var conf = rs.conf()
conf.members[0].host = "mongo1:27017"
rs.reconfig(conf, { force: true })
EOF