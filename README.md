# network-packets-exporter

## Usage
The Network Packets Exporter captures all network packets and stream them out as HTTP metrics, with useful labels like source/destination IP, port, and network protocol.
Great for debugging and collecting network statistics.

You are supposed to attach this image to a pod as a sidecar, and collect the exported metrics at the `METRICS_PORT`, which defaults to `9717` 

## Grafana Dashboard

Still Under development, Hopefully not for long :sleeping: 

## Helm Chart

Coming soon...
