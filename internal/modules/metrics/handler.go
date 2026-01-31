package metrics

import (
	"net/http"
)

func IngestMetric(w http.ResponseWriter, _ *http.Request){
	w.WriteHeader(http.StatusNotImplemented)
}