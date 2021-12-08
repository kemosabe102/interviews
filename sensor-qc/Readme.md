# Description
Sensor-QC is designed to classify sensors based on their readings

# Build
To build the Sensor-QC application, run the `/sensor-qc/build/build.sh` script

# Deploy using Helm
To deploy, go to `/sensor-qc/deploy` and run 
`helm upgrade --install sensor-qc --namespace sensor-qc --create-namespace --atomic --timeout 30s .`

## Define Log File Path
The `LOG_FILE_PATH` environment variable can be updated using the `Values.env.logFilePath` Helm values entry 