# Implmentation notes

`[aaronb@accra ~]$ curl localhost:6060/rest/Eos/TerminAttr/connection`

```json
{
    "alpha-cvp1.sjc.aristanetworks.com:9910": {
        "_ptr": "/Eos/TerminAttr/connection/alpha-cvp1.sjc.aristanetworks.com:9910"
    },
    "alpha-cvp2.sjc.aristanetworks.com:9910": {
        "_ptr": "/Eos/TerminAttr/connection/alpha-cvp2.sjc.aristanetworks.com:9910"
    },
    "alpha-cvp3.sjc.aristanetworks.com:9910": {
        "_ptr": "/Eos/TerminAttr/connection/alpha-cvp3.sjc.aristanetworks.com:9910"
    },
    "apiserver.cv-dev.corp.arista.io:443": {
        "_ptr": "/Eos/TerminAttr/connection/apiserver.cv-dev.corp.arista.io:443"
    }
}
```

You can learn more about each connection if you append the URL/IP to the path:
[aaronb@accra ~]$ curl localhost:6060/rest/Eos/TerminAttr/connection/alpha-cvp1.sjc.aristanetworks.com:9910

```json
{
    "Cluster": "alpha",
    "SourceAddress": "172.20.253.81:59092",
    "SourceVrf": "default",
    "State": "connected",
    "name": "alpha-cvp1.sjc.aristanetworks.com:9910"
}
```