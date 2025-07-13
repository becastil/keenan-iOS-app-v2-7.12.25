# Sydney Health Clone - Deployment Guide

## Overview

This guide covers deployment procedures for all Sydney Health Clone components across different environments.

## Environments

| Environment | Purpose | URL | Branch |
|------------|---------|-----|--------|
| Development | Active development | dev.sydneyhealth.com | develop |
| Staging | Pre-production testing | staging.sydneyhealth.com | staging |
| Production | Live environment | app.sydneyhealth.com | main |

## Prerequisites

### Required Access
- Kubernetes cluster access
- Docker registry credentials
- Database admin credentials
- SSL certificates
- Domain DNS control
- Monitoring dashboard access

### Required Tools
- `kubectl` configured for target clusters
- `helm` 3.0+
- `docker` CLI
- `gcloud` SDK (for GCP)
- Spinnaker CLI (optional)

## Backend Deployment

### 1. Build and Push Docker Images

```bash
# Set variables
export VERSION=$(git rev-parse --short HEAD)
export REGISTRY=gcr.io/sydney-health

# Build all services
cd backend
for service in gateway member benefits claims provider messaging; do
  cd services/$service
  docker build -t $REGISTRY/$service:$VERSION .
  docker tag $REGISTRY/$service:$VERSION $REGISTRY/$service:latest
  docker push $REGISTRY/$service:$VERSION
  docker push $REGISTRY/$service:latest
  cd ../..
done
```

### 2. Database Migration

```bash
# Connect to production database
kubectl port-forward -n sydney-health-prod svc/mysql 3306:3306

# Run migrations
migrate -path backend/migrations \
  -database "mysql://user:pass@localhost:3306/sydney_health" \
  up
```

### 3. Deploy to Kubernetes

#### Using Helm Charts

```bash
# Add custom values for environment
cat > values-production.yaml <<EOF
environment: production
image:
  tag: $VERSION
  pullPolicy: IfNotPresent

gateway:
  replicas: 5
  resources:
    requests:
      memory: "512Mi"
      cpu: "250m"
    limits:
      memory: "1Gi"
      cpu: "1000m"

database:
  host: mysql.sydney-health-prod.svc.cluster.local
  secretName: mysql-credentials

kafka:
  brokers:
    - kafka-0.kafka.sydney-health-prod.svc.cluster.local:9092
    - kafka-1.kafka.sydney-health-prod.svc.cluster.local:9092
    - kafka-2.kafka.sydney-health-prod.svc.cluster.local:9092
EOF

# Deploy using Helm
helm upgrade --install sydney-health ./helm/sydney-health \
  -f values-production.yaml \
  -n sydney-health-prod \
  --create-namespace
```

#### Using kubectl

```bash
# Apply configurations
kubectl apply -f backend/k8s/production/ -n sydney-health-prod

# Update image versions
kubectl set image deployment/gateway-service \
  gateway=$REGISTRY/gateway:$VERSION \
  -n sydney-health-prod

# Wait for rollout
kubectl rollout status deployment/gateway-service -n sydney-health-prod
```

### 4. Configure Ingress

```yaml
# backend/k8s/production/ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: sydney-health-ingress
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
    nginx.ingress.kubernetes.io/rate-limit: "100"
spec:
  tls:
  - hosts:
    - api.sydneyhealth.com
    secretName: api-tls-secret
  rules:
  - host: api.sydneyhealth.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: gateway-service
            port:
              number: 8080
```

## Web Application Deployment

### 1. Build Production Bundle

```bash
cd web
npm run build

# Test production build locally
npm start
```

### 2. Deploy to CDN

```bash
# Build Docker image
docker build -t $REGISTRY/web:$VERSION .
docker push $REGISTRY/web:$VERSION

# Deploy to Kubernetes
kubectl apply -f web/k8s/production/
kubectl set image deployment/web-app web=$REGISTRY/web:$VERSION -n sydney-health-prod
```

### 3. Configure CDN (CloudFlare)

```bash
# Upload static assets
aws s3 sync dist/static s3://sydney-health-static/web/

# Invalidate cache
aws cloudfront create-invalidation \
  --distribution-id E1234567890 \
  --paths "/*"
```

## Mobile App Deployment

### iOS Deployment

1. **Build Archive**
```bash
cd ios
xcodebuild archive \
  -workspace SydneyHealth.xcworkspace \
  -scheme SydneyHealth \
  -configuration Release \
  -archivePath build/SydneyHealth.xcarchive
```

2. **Export IPA**
```bash
xcodebuild -exportArchive \
  -archivePath build/SydneyHealth.xcarchive \
  -exportPath build/ \
  -exportOptionsPlist ExportOptions.plist
```

3. **Upload to App Store Connect**
```bash
xcrun altool --upload-app \
  --type ios \
  --file build/SydneyHealth.ipa \
  --username $APPLE_ID \
  --password $APP_SPECIFIC_PASSWORD
```

### Android Deployment

1. **Build Release APK**
```bash
cd android
./gradlew assembleRelease
```

2. **Sign APK**
```bash
jarsigner -verbose -sigalg SHA256withRSA -digestalg SHA-256 \
  -keystore release-key.keystore \
  app/build/outputs/apk/release/app-release-unsigned.apk \
  release-key
```

3. **Upload to Google Play Console**
```bash
# Using fastlane
fastlane supply --apk app/build/outputs/apk/release/app-release.apk
```

## Infrastructure Deployment

### Kafka Cluster

```bash
# Deploy Kafka using Strimzi operator
kubectl create namespace kafka
kubectl apply -f 'https://strimzi.io/install/latest?namespace=kafka'

# Create Kafka cluster
kubectl apply -f backend/k8s/kafka/kafka-cluster.yaml -n kafka
```

### Database Deployment

```bash
# Deploy MySQL using operator
kubectl apply -f https://raw.githubusercontent.com/mysql/mysql-operator/trunk/deploy/deploy-crds.yaml
kubectl apply -f https://raw.githubusercontent.com/mysql/mysql-operator/trunk/deploy/deploy-operator.yaml

# Create MySQL cluster
kubectl apply -f backend/k8s/mysql/mysql-cluster.yaml -n sydney-health-prod
```

### Monitoring Stack

```bash
# Deploy Prometheus operator
kubectl apply -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/main/bundle.yaml

# Deploy custom monitoring configuration
kubectl apply -f devops/observability/k8s/ -n monitoring
```

## CI/CD Pipeline

### Spinnaker Pipeline Execution

```bash
# Trigger pipeline via API
curl -X POST https://spinnaker.sydneyhealth.com/gate/pipelines/sydney-health/backend-pipeline \
  -H "Content-Type: application/json" \
  -d '{
    "type": "manual",
    "parameters": {
      "branch": "main",
      "version": "'$VERSION'"
    }
  }'
```

### Manual Deployment Steps

1. **Pre-deployment Checks**
```bash
# Check cluster health
kubectl get nodes
kubectl top nodes

# Check current deployments
kubectl get deployments -n sydney-health-prod

# Run smoke tests against staging
./scripts/smoke-tests.sh staging
```

2. **Deploy Services**
```bash
# Deploy in order
./scripts/deploy.sh gateway $VERSION
./scripts/deploy.sh member $VERSION
./scripts/deploy.sh benefits $VERSION
./scripts/deploy.sh claims $VERSION
./scripts/deploy.sh provider $VERSION
./scripts/deploy.sh messaging $VERSION
```

3. **Post-deployment Verification**
```bash
# Check pod status
kubectl get pods -n sydney-health-prod

# Check logs
kubectl logs -l app=gateway-service -n sydney-health-prod

# Run health checks
./scripts/health-check.sh production

# Run smoke tests
./scripts/smoke-tests.sh production
```

## Rollback Procedures

### Automatic Rollback

Configured in Spinnaker pipeline:
- Failed health checks trigger automatic rollback
- Performance degradation triggers rollback
- Error rate > 5% triggers rollback

### Manual Rollback

```bash
# Get current revision
kubectl rollout history deployment/gateway-service -n sydney-health-prod

# Rollback to previous version
kubectl rollback undo deployment/gateway-service -n sydney-health-prod

# Rollback to specific revision
kubectl rollback undo deployment/gateway-service --to-revision=2 -n sydney-health-prod
```

## Monitoring Post-Deployment

### Key Metrics to Monitor

1. **Application Metrics**
   - Request rate
   - Error rate (target < 0.1%)
   - Response time (p95 < 500ms)
   - Active users

2. **Infrastructure Metrics**
   - CPU usage (< 80%)
   - Memory usage (< 90%)
   - Disk I/O
   - Network throughput

3. **Business Metrics**
   - Claims processed
   - Messages sent
   - Provider searches
   - Login success rate

### Dashboards

Access monitoring dashboards:
- Grafana: https://grafana.sydneyhealth.com
- Jaeger: https://jaeger.sydneyhealth.com
- Kibana: https://kibana.sydneyhealth.com

## Security Considerations

### Pre-deployment Security Checks

```bash
# Scan Docker images
docker scan $REGISTRY/gateway:$VERSION

# Run security tests
./scripts/security-scan.sh

# Check secrets
kubesec scan backend/k8s/production/secrets.yaml
```

### SSL Certificate Management

```bash
# Check certificate expiration
kubectl get certificate -n sydney-health-prod

# Manually renew if needed
kubectl delete certificate api-tls-cert -n sydney-health-prod
# cert-manager will automatically create new certificate
```

## Disaster Recovery

### Backup Procedures

```bash
# Database backup
kubectl exec -it mysql-0 -n sydney-health-prod -- \
  mysqldump -u root -p sydney_health > backup-$(date +%Y%m%d).sql

# Kubernetes configuration backup
velero backup create sydney-health-backup --include-namespaces sydney-health-prod
```

### Restore Procedures

```bash
# Restore database
kubectl exec -it mysql-0 -n sydney-health-prod -- \
  mysql -u root -p sydney_health < backup-20240115.sql

# Restore Kubernetes resources
velero restore create --from-backup sydney-health-backup
```

## Troubleshooting Deployment Issues

### Common Issues

1. **Pods not starting**
```bash
# Check pod events
kubectl describe pod <pod-name> -n sydney-health-prod

# Check resource limits
kubectl top pods -n sydney-health-prod
```

2. **Service unavailable**
```bash
# Check service endpoints
kubectl get endpoints -n sydney-health-prod

# Check ingress
kubectl describe ingress sydney-health-ingress -n sydney-health-prod
```

3. **Database connection issues**
```bash
# Test connection from pod
kubectl exec -it gateway-service-xxx -n sydney-health-prod -- \
  nc -zv mysql.sydney-health-prod.svc.cluster.local 3306
```

### Emergency Procedures

1. **Full rollback**
```bash
./scripts/emergency-rollback.sh
```

2. **Scale down to reduce load**
```bash
kubectl scale deployment --all --replicas=1 -n sydney-health-prod
```

3. **Enable maintenance mode**
```bash
kubectl apply -f backend/k8s/maintenance-mode.yaml
```

## Contact Information

### Escalation Path

1. On-call engineer: Check PagerDuty
2. Team lead: #sydney-health-oncall Slack
3. Infrastructure team: infrastructure@sydneyhealth.com
4. Security team: security@sydneyhealth.com

### Runbooks

- [Gateway Service Runbook](./runbooks/gateway.md)
- [Database Runbook](./runbooks/database.md)
- [Kafka Runbook](./runbooks/kafka.md)
- [Incident Response](./runbooks/incident-response.md)