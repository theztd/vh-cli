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


### Sprava DNS

```bash
# Vypsani seznamu domen
vh-cli dns list
vh-cli dns zones

# Vypsani obsahu zony
vh-cli dns list -z fejk.net

# Vytvoreni zaznamu
vh-cli dns add -z fejk.net -t TXT -n pokus1 -v "Hodnota zaznamu" -c "Komentar"

# Smazani zaznamu
vh-cli dns rm -z fejk.net -id ID_ZAZNAMU
```



