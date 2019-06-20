package test

import (
	"client-go-helper/pkg/kubectl"
	"log"
	"testing"
)

var yaml =`
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: event-processor
  labels:
    draft: draft-app
    chart: "event-processor-1.20190215.1751-d-dev"
spec:
  replicas: 1
  strategy:
    rollingUpdate:
      maxSurge: 15%
      maxUnavailable: 15%
  template:
    metadata:
      labels:
        draft: draft-app
        app: event-processor
      annotations:
        prometheus.io/path: /prometheus
        prometheus.io/scrape: "true"
        
    spec:
      containers:
      - name: event-processor
        image: "core.harbor.172.16.20.63.nip.io/event-processor:1.20190215.1751-d"
        imagePullPolicy: IfNotPresent
        env:
        - name: JAVA_COMMAND
          value: java  -Xms256M -Xmx256M -Dspring.profiles.active=${PROFILE} -jar app.jar
        - name: PROFILE
          value: dev
        - name: EUREKA_INSTANCE_PREFER_IP_ADDRESS
          value: "false"
        - name: EUREKA_INSTANCE_HOSTNAME
          value: event-processor.dev.svc.cluster.local
        - name: EUREKA_CLIENT_SERVICEURL_DEFAULTZONE
          value: restful://eureka.dev.svc.cluster.local:8761/eureka/
        - name: MY_NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        - name: MY_POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        ports:
        - containerPort: 7212
        livenessProbe:
          httpGet:
            path: /health
            port: 7212
          initialDelaySeconds: 600
          periodSeconds: 30
          successThreshold: 1
          timeoutSeconds: 5
        readinessProbe:
          httpGet:
            path: /health
            port: 7212
          periodSeconds: 30
          successThreshold: 1
          timeoutSeconds: 5
        resources:
          limits:
            cpu: 300m
            memory: 512Mi
          requests:
            cpu: 60m
            memory: 409Mi
        volumeMounts:
        - mountPath: /app/log
          name: applog
        
      volumes:
      - hostPath:
          path: /app/log
          type: DirectoryOrCreate
        name: applog
      
      terminationGracePeriodSeconds: 20
`


func TestApply(t *testing.T) {
	e :=  kubectl.Apply(yaml,"local")
	if e !=nil {
		log.Print(e)
	}

}

func TestCreate(t *testing.T) {
	e :=  kubectl.Create(yaml,"local")
	if e !=nil {
		log.Print(e)
	}

}

