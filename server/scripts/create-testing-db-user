#!/bin/bash

ENV_PATH="./server/.env"

# Load .env file
if [ -f "$ENV_PATH" ]; then
    export $(cat "$ENV_PATH" | xargs)
else
    echo ".env file not found at $ENV_PATH"
    exit 1
fi

# Connect to MySQL and run the user creation and privileges commands
mysql -u root -p"$ROOT_DB_PW" -h "$MYSQL_HOST" -e "
CREATE USER '$LOCAL_KLEIO_USER'@'localhost' IDENTIFIED BY '$LOCAL_KLEIO_PW';
GRANT ALL PRIVILEGES ON *.* TO '$LOCAL_KLEIO_USER'@'localhost' WITH GRANT OPTION;
FLUSH PRIVILEGES;
"

echo "MySQL user '$LOCAL_KLEIO_USER' created and granted all privileges."
