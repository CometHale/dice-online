# Dice Online
A web-based dice rolling game (Service Engineer Code Challenge)



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

