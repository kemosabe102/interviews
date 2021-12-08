# Description
Sensor-QC is designed to classify sensors based on their readings in a log file

# Build
To build the Sensor-QC application, run the `/sensor-qc/build/build.sh` script

# Deploy
To deploy, go to `/sensor-qc/deploy` and run 
`helm upgrade --install sensor-qc --namespace sensor-qc --create-namespace --atomic --timeout 30s .`