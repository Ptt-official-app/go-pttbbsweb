#!/bin/bash

if [ "$#" != "2" ]; then
    echo "usage: docker_initbbs.sh [BBSHOME] [GOBBS_IMAGE]"
    exit 255
fi

BBSHOME=$1
GOBBS_IMAGE=$2
echo "BBSHOME: ${BBSHOME} GOBBS_IMAGE: ${GOBBS_IMAGE}"

project=go-initbbs

docker run --rm --name ${project} -v ${BBSHOME}:/home/bbs ${GOBBS_IMAGE} /opt/bbs/bin/initbbs -DoIt

docker run --rm --name ${project} -v ${BBSHOME}:/home/bbs ${GOBBS_IMAGE} sh -c 'mkdir -p /home/bbs/bin && cp /opt/bbs/bin/* /home/bbs/bin'

docker run --rm --name ${project} -v ${BBSHOME}:/home/bbs ${GOBBS_IMAGE} sh -c 'mkdir -p /home/bbs/etc && cp /opt/bbs/etc/* /home/bbs/etc'

docker run --rm --name ${project} -v ${BBSHOME}:/home/bbs ${GOBBS_IMAGE} sh -c 'mkdir -p /home/bbs/wsproxy && cp -R /opt/bbs/wsproxy/* /home/bbs/wsproxy'
