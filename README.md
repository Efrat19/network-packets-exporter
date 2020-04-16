# network-packets-exporter 

![Publish Packets Exporter](https://github.com/Efrat19/network-packets-exporter/workflows/Publish%20Packets%20Exporter/badge.svg)

## Usage
The Network Packets Exporter uses google [gopacket library](https://github.com/google/gopacket) to capture all network packets and stream them out as prometheus-readable HTTP metrics, with useful labels like source/destination IP, port, and network protocol.
Great for debugging and collecting network statistics.

You are supposed to attach this image to a pod as a sidecar, and collect the exported metrics at the `METRICS_PORT`, which defaults to `9717`:

```yaml
containers:
- <your main container>

- name: network-metrics-sidecar
  image: efrat19/packets-exporter:stable-v1.0.2
```

Adding a Service:
```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    name: network-monitoring
  name: network-monitoring
  namespace: <your_namespace>
spec:
  ports:
  - name: http-metrics
    port: 9717
    protocol: TCP
    targetPort: 9717
  selector:
    <your_app_labels>
  type: ClusterIP
```
Adding the ServiceMonitor (for Prometheus-operator controlled systems):

```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: network-monitoring
  namespace: monitoring
  labels:
    <your_service_monitor_selector_labels>
spec:
  endpoints:
  - path: /metrics
    port: http-metrics
    interval: 10s
  namespaceSelector:
    matchNames:
      - <your_namespace>
  selector:
    matchLabels:
      name: network-monitoring
```

## Grafana Dashboard

Still Under development, Hopefully not for long :sleeping: 



