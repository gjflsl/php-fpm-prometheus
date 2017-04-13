# php-fpm-prometheus

Simple [PHP-FPM](http://php.net/manual/en/install.fpm.php) status exporter for [Prometheus](https://prometheus.io/).

## Usage

```
$./php-fpm-prometheus --help
Usage of /data/app/php-fpm-prometheus:
  -php-fpm-addr string
    	PHP-FPM address tcp://127.0.0.0.1:900 or unix:/path/to/unix/socket
    	
$./php-fpm-prometheus  -php-fpm-addr unix:/var/run/php5-fpm.sock
# HELP php_fpm_start_since Seconds since FPM start
# TYPE php_fpm_start_since counter
php_fpm{key="php_fpm_start_since",fgci_addr="/var/run/php5-fpm.sock"} 93940
# HELP php_fpm_accepted_conn Total of accepted connections
# TYPE php_fpm_accepted_conn counter
php_fpm{key="php_fpm_accepted_conn",fgci_addr="/var/run/php5-fpm.sock"} 5827
# HELP php_fpm_listen_queue Number of connections that have been initiated but not yet accepted
# TYPE php_fpm_listen_queue gauge
php_fpm{key="php_fpm_listen_queue",fgci_addr="/var/run/php5-fpm.sock"} 0
# HELP php_fpm_max_listen_queue Max. connections the listen queue has reached since FPM start
# TYPE php_fpm_max_listen_queue counter
php_fpm{key="php_fpm_max_listen_queue",fgci_addr="/var/run/php5-fpm.sock"} 0
# HELP php_fpm_listen_queue_length Maximum number of connections that can be queued
# TYPE php_fpm_listen_queue_length gauge
php_fpm{key="php_fpm_listen_queue_length",fgci_addr="/var/run/php5-fpm.sock"} 0
# HELP php_fpm_idle_processes Idle process count
# TYPE php_fpm_idle_processes gauge
php_fpm{key="php_fpm_idle_processes",fgci_addr="/var/run/php5-fpm.sock"} 2
# HELP php_fpm_active_processes Active process count
# TYPE php_fpm_active_processes gauge
php_fpm{key="php_fpm_active_processes",fgci_addr="/var/run/php5-fpm.sock"} 1
# HELP php_fpm_total_processes Total process count
# TYPE php_fpm_total_processes gauge
php_fpm{key="php_fpm_total_processes",fgci_addr="/var/run/php5-fpm.sock"} 3
# HELP php_fpm_max_active_processes Maximum active process count
# TYPE php_fpm_max_active_processes counter
php_fpm{key="php_fpm_max_active_processes",fgci_addr="/var/run/php5-fpm.sock"} 5
# HELP php_fpm_max_children_reached Number of times the process limit has been reached
# TYPE php_fpm_max_children_reached counter
php_fpm{key="php_fpm_max_children_reached",fgci_addr="/var/run/php5-fpm.sock"} 4
# HELP php_fpm_slow_requests Number of requests that exceed request_slowlog_timeout
# TYPE php_fpm_slow_requests counter
php_fpm{key="php_fpm_slow_requests",fgci_addr="/var/run/php5-fpm.sock"} 0
# HELP php_fpm_exporter_scrape_failures_total Number of errors while scraping php_fpm
# TYPE php_fpm_exporter_scrape_failures_total counter
php_fpm{key="php_fpm_exporter_scrape_failures_total",fgci_addr="/var/run/php5-fpm.sock"} 0

```

You can echo to node_exporter collector.textfile.directory.

[Grafana dashboard config](./php_fpm.json)

## Contributing

All contributions are welcome, but if you are considering significant changes, please open an issue beforehand and discuss it with us.

## License

MIT. See the `LICENSE` file for more information.

