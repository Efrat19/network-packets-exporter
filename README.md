# network-packets-exporter

## Usage
The Network Packets Exporter captures all network packets and stream them out as HTTP metrics, with useful labels like source/destination IP, port, and network protocol.
Great for debugging and collecting network statistics.

You are supposed to attach this image to a pod as a sidecar, and collect the exported metrics at the `METRICS_PORT`, which defaults to `9717`:

```yaml
containers:
- <your main container>

- name: network-metrics-sidecar
  image: efrat19/packets-exporter:23b366679e824e1b2a794ea5b7efdb3778724b64
#  env:
#  - name: METRICS_PORT
#    value: "9717"
#  - name: IFACE
#    value: eth0
```



## Exported Metrics



## Grafana Dashboard

Still Under development, Hopefully not for long :hugging_face: 

## Helm Chart

sidecar
service
servicemonitor


Coming soon...

## RoadMap

- [ ] Tests


