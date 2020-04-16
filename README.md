# network-packets-exporter CURRENTLY UNDER HEAVY DEVELOPMENT

## Usage
The Network Packets Exporter captures all network packets and stream them out as HTTP metrics, with useful labels like source/destination IP, port, and network protocol.
Great for debugging and collecting network statistics.

You are supposed to attach this image to a pod as a sidecar, and collect the exported metrics at the `METRICS_PORT`, which defaults to `9717`:

```yaml
containers:
- <your main container>

- name: network-metrics-sidecar
  image: efrat19/packets-exporter:stable
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
Adding the ServiceMonitor:

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

## Exported Metrics



## Grafana Dashboard

Still Under development, Hopefully not for long :sleeping: 

## Helm Chart

sidecar
service
servicemonitor


Coming soon...

## RoadMap

- [ ] Tests


