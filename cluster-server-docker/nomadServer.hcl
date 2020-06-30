# Increase log verbosity
log_level = "DEBUG"

# Setup data dir
data_dir = "/tmp/server1"

# Enable the server
server {
    enabled = true

    # Self-elect, should be 3 or 5 for production
    bootstrap_expect = 1
}

client {
  enabled       = true
  network_speed = 10
}

telemetry {
  collection_interval = "1s"
  disable_hostname = true
  prometheus_metrics = true
  publish_allocation_metrics = true
  publish_node_metrics = true
}