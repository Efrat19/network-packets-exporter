package main
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
	http.Handle("/", Adapt(http.FileServer(http.Dir("./static")), logAccess()))
	http.Handle("/metrics", Adapt(promhttp.Handler(), logAccess()))
	fmt.Fprintln(os.Stderr, "Starting server at port 9717")

	go func() {
		log.Fatal(http.ListenAndServe(":9717", nil))
	}()
}