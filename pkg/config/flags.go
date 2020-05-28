package config

import (
	"flag"
	"time"

	"github.com/openshift/elasticsearch-proxy/pkg/util"
)

func newFlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("elasticsearch-proxy", flag.ExitOnError)

	flagSet.String("listening-address", ":8443", "<addr>:<port> to listen on for HTTPS clients")

	flagSet.String("tls-cert", "", "path to certificate file")
	flagSet.String("tls-key", "", "path to private key file")
	flagSet.String("tls-client-ca", "", "path to a CA file for admitting client certificates.")

	flagSet.String("metrics-listening-address", "", "<addr>:<port> to listen on for HTTPS metrics clients")
	flagSet.String("metrics-tls-cert", "", "path to certificate file from the metrics service")
	flagSet.String("metrics-tls-key", "", "path to private key file from the metrics service")

	flagSet.String("elasticsearch-url", "https://localhost:9200", "The default URL to Elasticsearch")

	flagSet.Bool("ssl-insecure-skip-verify", false, "skip validation of certificates presented when using HTTPS")
	flagSet.Bool("proxy-websockets", true, "enables WebSocket proxying")
	flagSet.Var(&util.StringArray{}, "openshift-ca", "paths to CA roots for the OpenShift API (may be given multiple times, defaults to /var/run/secrets/kubernetes.io/serviceaccount/ca.crt).")
	flagSet.Bool("request-logging", false, "Log requests to stdout")

	flagSet.Duration("upstream-flush", time.Duration(5)*time.Millisecond, "force flush upstream responses after this duration(useful for streaming responses). 0 to never force flush. Defaults to 5ms")
	flagSet.Var(&util.StringArray{}, "upstream-ca", "paths to CA roots for the Upstream (target) Server (may be given multiple times, defaults to system trust store).")
	flagSet.Duration("cache-expiry", time.Duration(5)*time.Minute, "cache expiration duration. The cache stores a specific set of OpenShift objects (projects, sar) used by the proxy.")

	//Auth flags
	flagSet.Var(&util.StringArray{}, "auth-backend-role", "A SAR to check to allow the given backend role(i.e. admin={'namespace':'default','verb':'get','resource':'pods/logs'}")
	flagSet.Var(&util.StringArray{}, "auth-whitelisted-name", "A name compared against cert CN for which a request will be passed through")
	flagSet.String("auth-admin-role", "", "The name of the only role that will be passed on the request if it is found in the list of roles")
	flagSet.String("auth-default-role", "", "The role given to every request unless it has the auth-admin-role")

	return flagSet
}
