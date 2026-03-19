#!/bin/bash
# Start kubectl port-forwards for PipeCD services
# Only runs if the pipecd namespace exists (i.e. PipeCD has been deployed)

if ! kubectl get namespace pipecd &>/dev/null; then
  echo "PipeCD namespace not found, skipping port-forwards."
  exit 0
fi

echo "Starting PipeCD port-forwards..."

# Kill any existing port-forwards for these ports
pkill -f "kubectl port-forward.*svc/pipecd" 2>/dev/null || true

# Forward all PipeCD ports with 0.0.0.0 so Ona's proxy can reach them
kubectl port-forward --address 0.0.0.0 -n pipecd svc/pipecd 8080 &>/tmp/pf-8080.log &
kubectl port-forward --address 0.0.0.0 -n pipecd svc/pipecd 8443 &>/tmp/pf-8443.log &
kubectl port-forward --address 0.0.0.0 -n pipecd svc/pipecd-ops 9080 &>/tmp/pf-9080.log &

echo "Port-forwards started:"
echo "  8080 -> PipeCD Web UI"
echo "  8443 -> PipeCD API"
echo "  9080 -> PipeCD Ops"
