// Package rulefile
// TODO: doc
// Example:
// 8080 example.com:80
// 8443 example.com:443 direct dns=1.1.1.1,tfo
// 192.168.1.1:8443 example.com:443
// eth0:8443 example.com:443 direct interface=eth1
// tcp://eth0:8443 example.com:443,www.example.com:443 loadbalance policy=round-robin
package rulefile
