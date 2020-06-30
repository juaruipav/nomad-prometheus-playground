# Increase log verbosity
log_level = "DEBUG"

# Setup data dir
data_dir = "/tmp/client"

# Enable the client
client {
    enabled = true

    # For demo assume we are talking to server1. For production,
    # this should be like "nomad.service.consul:4647" and a system
    # like Consul used for service discovery.
   servers = ["${NOMAD_MASTER_ADDR}"]
}

consul {
	address = "${CONSUL_MASTER_ADDR}:8500"
}

# Modify our port to avoid a collision with server1
#ports {
   # http = 5656
#}

telemetry {
  collection_interval = "1s"
  prometheus_metrics = true
  publish_allocation_metrics = true
  publish_node_metrics = true
  use_node_name = false
  disable_hostname = false
}