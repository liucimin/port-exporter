package main


import (

	"docker-interface-exporter/pkg/containers"
	"docker-interface-exporter/pkg/tools"
	"flag"
	"fmt"
	"os"
	"github.com/golang/glog"
	"context"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
)





func init() {
	prometheus.MustRegister(version.NewCollector("collectd_exporter"))
}

func main() {


	flag.Parse()

	var (
		listenAddress = flag.String("web.listen-address", ":9103", "Address on which to expose metrics and web interface.")
		metricsPath      = flag.String("web.telemetry-path", "/metrics", "Path under which to expose Prometheus metrics.")

		)

	glog.Infoln("Starting collectd_exporter", version.Info())
	glog.Infoln("Build context", version.BuildContext())

	c := newCollectdCollector()
	prometheus.MustRegister(c)

	startCollectdServer(context.Background(), c)

	if *collectdPostPath != "" {
		http.HandleFunc(*collectdPostPath, c.collectdPost)
	}

	http.Handle(*metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Collectd Exporter</title></head>
             <body>
             <h1>Collectd Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
