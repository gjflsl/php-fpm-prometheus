# php-fpm-prometheus

Simple [PHP-FPM](http://php.net/manual/en/install.fpm.php) status exporter for [Prometheus](https://prometheus.io/).

## Usage

```
$ ./php-fpm-prometheus --help
Usage of ./php-fpm-prometheus:
  -php-fpm-list string
    	PHP-FPM list file path
    	
$ ./php-fpm-prometheus -php-fpm-list php_fpm_list.txt

```

You can echo to node_exporter collector.textfile.directory.

[Grafana dashboard config](./php_fpm.json)

## Contributing

All contributions are welcome, but if you are considering significant changes, please open an issue beforehand and discuss it with us.

## License

MIT. See the `LICENSE` file for more information.

