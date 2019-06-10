module client-go-helper

go 1.12

replace (
	github.com/renstrom/dedent => github.com/lithammer/dedent v1.1.0
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421 => github.com/golang/oauth2 v0.0.0-20190226205417-e64efc72b421
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2 => github.com/golang/text v0.3.1-0.20180807135948-17ff2d5776d2
	google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19 => github.com/google/go-genproto v0.0.0-20190307195333-5fe7a883aa19
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/MakeNowJust/heredoc v0.0.0-20171113091838-e9091a26100e // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.4.2 // indirect
	github.com/Masterminds/sprig v2.18.0+incompatible // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/Sirupsen/logrus v1.0.6
	github.com/cenkalti/backoff v2.1.1+incompatible
	github.com/chai2010/gettext-go v0.0.0-20170215093142-bf70f2a70fb1 // indirect
	github.com/coreos/etcd v3.3.12+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/cyphar/filepath-securejoin v0.2.2 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.0 // indirect
	github.com/emicklei/go-restful v2.9.1+incompatible // indirect
	github.com/emicklei/go-restful-swagger12 v0.0.0-20170926063155-7524189396c6 // indirect
	github.com/evanphx/json-patch v4.1.0+incompatible
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/color v1.7.0
	github.com/ghodss/yaml v1.0.0
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-openapi/spec v0.18.0 // indirect
	github.com/go-openapi/strfmt v0.18.0 // indirect
	github.com/go-openapi/validate v0.18.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf // indirect
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/gorm v1.9.4
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/mitchellh/go-homedir v1.0.0
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20120707110453-a547fc61f48d // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/common v0.2.0
	github.com/renstrom/dedent v1.1.0 // indirect
	github.com/russross/blackfriday v1.5.2 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.2
	github.com/src-d/go-git v4.7.0+incompatible
	github.com/technosophos/moniker v0.0.0-20180509230615-a5dbd03a2245 // indirect
	github.com/ugorji/go v1.1.4 // indirect
	golang.org/x/net v0.0.0-20190320064053-1272bf9dcd53 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	google.golang.org/grpc v1.19.1 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.8.2
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/square/go-jose.v2 v2.3.0 // indirect
	gopkg.in/src-d/go-billy.v4 v4.3.0 // indirect
	gopkg.in/src-d/go-git.v4 v4.10.0
	gopkg.in/urfave/cli.v1 v1.20.0
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20181213150558-05914d821849
	k8s.io/apiextensions-apiserver v0.0.0-20181213153335-0fe22c71c476 // indirect
	k8s.io/apimachinery v0.0.0-20181127025237-2b1284ed4c93
	k8s.io/apiserver v0.0.0-20181213151703-3ccfe8365421 // indirect
	k8s.io/cli-runtime v0.0.0-20181213153952-835b10687cb6 // indirect
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/helm v2.13.1+incompatible
	k8s.io/klog v0.2.0 // indirect
	k8s.io/kube-openapi v0.0.0-20190320154901-5e45bb682580 // indirect
	k8s.io/kubernetes v1.13.5-beta.0.0.20190322001621-1a91ffde19dd
	k8s.io/utils v0.0.0-20190308190857-21c4ce38f2a7 // indirect
	sigs.k8s.io/kustomize v2.0.3+incompatible // indirect
	sigs.k8s.io/yaml v1.1.0
	vbom.ml/util v0.0.0-20180919145318-efcd4e0f9787 // indirect

)
