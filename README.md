# Vas-Hosting CLI

Command Line Interface pro Vas-Hosting management


## Installace

Aplikace je napsana v jazyce GO, staci tedy jen stahnout releasnutou binarku pro vas operacni system a zacit pouzivat.



## Autentizace

```bash
export VH_API_KEY="Tajny-Vygenerovany-Token"

# Defaultni hodnota je nastavena, nemate-li vlastni domenu, neni potreba nic menit
export VH_URL="https://centrum.vas-hosting.cz/api/v1/"
```

Pripadne lze vytvorit soubor $HOME/.vh/config.env s obsahem

```toml
[default]
VH_URL="https://centrum.vas-hosting.cz/api/v1/"
VH_API_KEY=Vygenerovany-Token

```

Nebo lze cestu ke konfiguracnimu souboru definovat take prez promenou **VH_CONFIG_PATH**

## Pouziti

### Automaticke doplnovani


```bash
# Ukazka na linuxu pro bash
vh-cli completion bash > /etc/profile.d/vh-cli.sh

```
### Sprava DNS

```bash
# Vypsani seznamu domen
vh-cli zone list

# Vypsani obsahu zony
vh-cli dns list -z fejk.net

# Vypsani obsahu zony s vyuzitim filtru
vh-cli dns list -z fejk.net -t A -n dev 

# Vypsani obsahu zony s pouzitim sablony
go run . dns l -z fejk.net -f ./examples/example-cloudflare-records.zone

# Vytvoreni zaznamu
vh-cli dns add -z fejk.net -t TXT -n pokus1 -v "Hodnota zaznamu" -c "Komentar"

# Smazani zaznamu
vh-cli dns del -z fejk.net -n nazev-zaznamu -t TXT
```

### Sprava SERVERU

```bash
# Vypis seznamu serveru
vh-cli servers list

# Vypis seznamu serveru s vyuzitim sablon
vh-cli servers list -f ./examples/example-ansible_iptables.yml

# Vypis seznamu serveru s vyuzitim sablon a zgroupovany dle definovaneho klice
vh-cli servers list -G hw -f ./examples/example-server_by_hw.html

# Vypis seznamu serveru s vyuzitim sablon a zgroupovany dle definovaneho klice vyfiltrovany dle labelu prod
go run . server list --filter-labels prod  -G hw -f examples/example-server_by_hw.html

# Vypis seznamu serveru ve formatu jak je obdrzen od VH API
vh-cli servers list-json

# Vypis informaci ke konkretnimu serveru
vh-cli servers list -n n1.fejk.net


# Reboot serveru (vyzaduje nasledne potvrzeni)
vh-cli servers reboot -n n1.fejk.net

```

### Debug mod

Pokud nastavim env promenou DEBUG na true, aplikace vypisuje informace o svem behu, ktere mohou pomoci s hledanim potizi...

```bash
DEBUG=true vh-cli server list -G hw
```

### Formatovani vystupu

Je mozne, ze budete chtit pouzit nastroj pro generovani vystupu v konkretnim formatu. Pro tyto ucely je k dispozici prepinac --template-file ktery definuje cestu k souboru ve formatu go/templates, kterym lze vygenerovat libovolny vystup. Pro zvyseni uzitecnosti lze v templatech pouzivat tyto funkce:

#### Contains
vraci true pokud vstupni list dat obsahuje hledany string

Priklad
```go
// Vytiskni jen polozky, ktere obsahuji label prod
{{ if Contains .Labels "prod" }}
{{ . }}
{{ end }}
```


#### Replace
nahradi ve vstupnich datech prvni vyskyt

Priklad
```go
// Zmeni v promene .Name vyskyty -prod na -production
{{ range . }}
{{ Replace .Name "-prod" "-production" }}
{{ end }}
```


#### ReplaceAll
nahradi ve vstupnich datech vsechny vyskyty

vstup
```
// nahradi vsechny vyskyty pismene "a" velkym pismenem "X"
{{ ReplaceAll .Name "a" "X" }}
```


#### Join
spoji pole definovanym delimiterem

vstup
```
// Spoji Vsechny labely do vystupniho formatu "prod, wordpress, debian11"
{{ Join .Labels ", " }}
```
