# cloudeos-sidecar

Sidecar for cloudeos pod to provide addtional metadata

## Current functionality

* Applies annotation to the cloudeos pod with the URL of cloudvision `arista/cloudvisionurl`.
Example output:

```bash
Name:         cloudeos-cj7cv
Namespace:    kube-system
Priority:     0
Node:         kube-67/10.90.224.67
Start Time:   Fri, 22 May 2020 10:43:30 -0700
Labels:       app=cloudeos
             controller-revision-hash=7765764df8
             pod-template-generation=7
Annotations:  arista/cloudvisionurl: https://10.90.224.175
```
