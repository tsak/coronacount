# Corona Count

Please panic!

Checking major news sites for number of mentions of the phrase "Corona", "Covid-19" or "Sars-Cov-2".

See for yourself at [cc.tsak.net](https://cc.tsak.net) (this is now a static copy of the service originally running and updating)

## Build

```bash
go build
```

## Run

```bash
Usage of ./coronocount:
  -d	Debug mode
  -i int
    	Scrape interval (default 30)
  -l string
    	Address and port to listen and serve on (default "localhost:8080")
  -s string
    	File to load URLs from (default "sites.txt")
```

## Run as a systemd service

See [coronacount.service](coronacount.service) systemd service definition.

To install (tested on Ubuntu 16.04):

1. `adduser coronacount`
2. copy `coronacount` binary as well as `sites.txt` and `template.html` to `/home/coronacount`
3. place systemd service script in `/lib/systemd/system/`
4. `sudo systemctl enable coronacount.service`
5. `sudo systemctl start coronacount`
6. `sudo journalctl -f -u coronacount`

The last command will show if the service was started.

