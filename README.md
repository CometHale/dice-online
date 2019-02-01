# Dice Online
A web-based dice rolling game (Service Engineer Code Challenge)

## Dice Game Explanation

At the beginning of the game, the user chooses a number between 1 and 6. 
The goal of the game is to successfully get a dice roll for that number. 

i.e. if the user chose 5 at the beginning of the game, the user will only win if
they roll a 5 on the dice.

## Stack

Golang
Vue.js

## Example Configuration File (JSON)

{
    "PostgreSQL": {
        "Username": "diceonline",
        "Password": "password",
        "Name": "diceonline",
        "Host": "localhost",
        "Port": "5432",
        "SSLMode": "disable"
    },
    "Server": {
        "Hostname": "127.0.0.1",
        "UseHTTP": true,
        "UseHTTPS": false,
        "HTTPPort": 3000,
        "HTTPSPort": 443,
        "CertFile": "tls/server.crt",
        "KeyFile": "tls/server.key"
    }
}

# Setup

## Database

1. Setup a database server per the settings in the configuration files
2. Run the provided sql scripts (in app/shared/scripts) to create the necessary tables


## Improvements To Be Made

1. Security

Instead of passing important data like user IDs around via JSON, develop a Session system.

2. Sessions and Cookies
