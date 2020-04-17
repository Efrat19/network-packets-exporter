package src
import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var PacketLabels = []string{
	"srcIP",
	"dstIP",

	"srcPort",
	"dstPort",
	"protocol"}


var (
	counterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "network",
			Name:      "packets_total",
			Help:      "help",
		},
		PacketLabels,
		)

)



func countPacket(packetLabels []string){
	counterVec.WithLabelValues(packetLabels...).Add(1)
}

func init(){
	prometheus.MustRegister(counterVec)
	fmt.Fprintln(os.Stderr, "Registering counter vector")
	startHTTPserver()
}

func startHTTPserver() {
	port := os.Getenv("METRICS_PORT")
	if port == "" {
		fmt.Fprintln(os.Stderr, "no METRICS_PORT env specified, defaulting to 9717")
		port = "9717"
	}
	http.Handle("/", Adapt(http.FileServer(http.Dir("./static")), logAccess()))
	http.Handle("/metrics", Adapt(promhttp.Handler(), logAccess()))
	fmt.Fprintln(os.Stderr, "Starting server at port " + port)

	go func() {
		log.Fatal(http.ListenAndServe(":" + port, nil))
	}()
}