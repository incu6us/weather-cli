weather-cli
---

## To install dependencies
```bash
make bin-install
```

### Other `make` commands
    * `make test` - run tests
    * `make lint` - run linter

## Before running the application
Create a config.yaml file in the root directory of the project. You can use the config.yaml.dist file as a template.

## Application
    * global options:
        --debug value, -d value   enable debug mode, which will print data from weather clients (default: false)
        --config value, -c value  (default: "config.yaml")
        --help, -h                show help

    * commands:
        run    run the weather-cli application
        help, h  Shows a list of commands or help for one command

    * run command options:
        --country value  choose country (default: "Ukraine")
        --city value     choose city (default: "Kyiv")
        --help, -h       show help

## Application output
```bash
time=2024-12-12T14:12:01.470+02:00 level=INFO msg="Runnig the weather-cli application"
time=2024-12-12T14:12:01.739+02:00 level=INFO msg="Client: weatherapi"
time=2024-12-12T14:12:01.739+02:00 level=INFO msg="Temperature: 0.100000"
time=2024-12-12T14:12:01.739+02:00 level=INFO msg="All Details: &{{1734003900 2024-12-12 13:45 0.1 32.3 1 {Sunny //cdn.weatherapi.com/weather/64x64/day/113.png 1000} 10.5 16.9 317 NW 1021 30.14 0 0 73 0 -4.6 23.8 -4.6 23.8 0.2 32.3 -4.1 24.7 10 6 0.6 14.7 23.7}}"
```
