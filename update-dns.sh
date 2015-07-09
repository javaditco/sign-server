#!/bin/bash
# script to update your dnsmasq configuration, to make it easier to use the hostname instead ip
# important: make sure that your dnsmasq is listening in the docker0 interface and that its
# support the /etc/dnsmasq.d config directory. run it with sudo

if [ ! -d /etc/dnsmasq.d/ ]; then
	mkdir /etc/dnsmasq.d
fi

container='sign-server'
new_ip=$(docker inspect --format '{{ .NetworkSettings.IPAddress }}' sign-server)
echo "host-record=sign-server-vpereira.gigantic.io,$new_ip" > /etc/dnsmasq.d/0host_$container
systemctl restart dnsmasq
