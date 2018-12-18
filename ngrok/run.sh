#!/bin/ash

if [[ ! -x /bin/ngrok ]]; then
    echo "ngrok not found in /bin/"
fi

if [[ ! -f /etc/ngrok.yml ]]; then
    echo "ngrok.yml not found in /etc/"
fi
echo "you can use TARGET env var to override the default which is localhost:8080"

TARGET=${TARGET-:localhost:8080}
cmd="/bin/ngrok http --config /etc/ngrok.yml $TARGET"
echo $cmd
exec $cmd
