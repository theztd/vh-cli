# Vas-Hosting CLI

Command Line Interface pro Vas-Hosting management


## Install

Aplikace je napsana v jazyce GO, staci tedy jen stahnout vybuildenou binarku pro vas operacni system a zacit pouzivat.



## Autentizace

```bash
export VH_API_KEY="Tajny-Vygenerovany-Token"

# Defaultni hodnota je nastavena, nemate-li vlastni domenu, neni potreba nic menit
export VH_URL="https://centrum.vas-hosting.cz/api/v1/"
```

Pripadne lze vytvorit soubor $HOME/.vh/config.ini s obsahem

```toml
[default]
VH_URL="https://centrum.vas-hosting.cz/api/v1/"
VH_API_KEY=Vygenerovany-Token

```

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

# Vypsani obsahu zony dle sablony
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

# Vypis informaci ke konkretnimu serveru
vh-cli servers list -n n1.fejk.net


# Reboot serveru (vyzaduje nasledne potvrzeni)
vh-cli servers reboot -n n1.fejk.net

```



### Formatovani vystupu

Je mozne, ze budete chtit pouzit nastroj pro generovani vystupu v konkretnim formatu. Pro tyto ucely je k dispozici prepinac --template-file ktery definuje cestu k souboru ve formatu go/templates, kterym lze vygenerovat libovolny vystup.
