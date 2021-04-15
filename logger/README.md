# Logging
```go
import "https://github.com/onelazyugy/golang-common/logger"
import "https://github.com/onelazyugy/golang-common/constants"

func Handler(w http.ResponseWriter, r *http.Request){
    //Populate related fields
    //Corelation id is a must
    fields := map[string]interface{}{
		"request_id": "Create Item", "user_ip": "TEST Item Create",
		"nep-correlation-id": r.Header.Get(constants.CorrelationID),
	}
    //Get a new logger with fields
	log := logger.New("ENV_PROFILE").WithFields(fields)
	log.Info("Creating items...")
    /*JSON Output:   
        {
            "level": "info",
            "msg": "Creating items...",
            "nep-correlation-id": 11235,
            "request_id": "Create Item",
            "time": "2021-02-25T17:23:12-05:00",
            "user_ip": "TEST Item Create"
        }
    */
    /*Pretty Print: INFO   [2021-02-25T17:24:43-05:00] Creating items...       nep-correlation-id=11235 request_id="Create Item" user_ip="TEST Item Create" */
}

//Always log at levels do not use print 
log.Info()
log.Debug()
log.Error()
log.Fatal()
```

The logger uses logrus logging package which is thread safe and compatable with go's 
standard logger. All standard log functions are exposed. [Logrus Documentation](https://pkg.go.dev/github.com/sirupsen/logrus)

Logs can be found [here](https://console.cloud.google.com/logs/query?folder=true&organizationId=1052607630679&project=ret-shasta-cug01-dev). Run the following query: 
```
resource.type="k8s_container" resource.labels.cluster_name="xxx" resource.labels.namespace_name="xxxx" resource.labels.container_name="xxxx"
```

How to [Query](https://cloud.google.com/logging/docs/view/logging-query-language) Logs?
