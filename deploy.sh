#!/bin/bash
go version 2>/dev/null 
if [ $? -eq 0 ]
then
    echo "go present"
else
    echo "we need install go"

PACK=go1.22.3.linux-amd64.tar.gz
CWDR=$(pwd)

curl -O https://dl.google.com/go/$PACK
tar -xvf $PACK -C /opt
rm $PACK
mkdir -p /opt/go
chown -R root:root /opt/go
/opt/go/bin/go version


fi

echo "Start go build"
/opt/go/bin/go mod tidy 
/opt/go/bin/go build -o ./cicd_agent ./cmd/cicd_agent
echo "End go build"

ls -lah

mkdir -p /opt/cicd_agent
mv ./cicd_agent /opt/cicd_agent

cp ./cicd_agent.service /lib/systemd/system

ln -s /lib/systemd/system/cicd_agent.service /etc/systemd/system/multi-user.target.wants/cicd_agent.service

systemctl restart cicd_agent && systemctl enable cicd_agent
systemctl daemon-reload

systemctl status cicd_agent.service
