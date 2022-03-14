export DL=test/logs/dnsmasq.log
export DLE=test/logs/dnsmasq.leases
make build-dev; ./builds/dev iptables $DL $DLE
