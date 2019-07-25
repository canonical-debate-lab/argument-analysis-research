module github.com/canonical-debate-lab/argument-analysis-research

go 1.12

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	github.com/Obaied/RAKE.Go v0.0.0-20181207222342-f0f99b2097df
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/google/uuid v1.1.1
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/mingrammer/commonregex v1.0.0 // indirect
	github.com/montanaflynn/stats v0.5.0 // indirect
	github.com/neurosnap/sentences v1.0.6 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829 // indirect
	github.com/seibert-media/golibs v1.0.3
	github.com/sethgrid/pester v0.0.0-20190127155807-68a33a018ad0
	github.com/stretchr/testify v1.3.0 // indirect
	go.etcd.io/bbolt v1.3.3
	go.opencensus.io v0.22.0
	go.uber.org/zap v1.10.0
	gonum.org/v1/gonum v0.0.0-20190710053202-4340aa3071a0 // indirect
	gopkg.in/djherbis/stow.v3 v3.0.0
	gopkg.in/jdkato/prose.v2 v2.0.0-20180825173540-767a23049b9e
	gopkg.in/neurosnap/sentences.v1 v1.0.6 // indirect
)

replace github.com/kelseyhightower/envconfig => github.com/seibert-media/envconfig v1.4.0
