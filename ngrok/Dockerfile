FROM alpine
# https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip
RUN set -ex \
  && apk add --no-cache --virtual .build-deps wget \
  && apk add --no-cache ca-certificates \
  && cd /bin && pwd \
  && wget https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip \
  && unzip ngrok-stable-linux-amd64.zip \
  && rm ngrok-stable-linux-amd64.zip && cd - \
  && apk del .build-deps

ADD ngrok.yml /etc/ngrok.yml
ADD run.sh /bin
RUN chmod 755 /bin/run.sh
EXPOSE 4040

ENTRYPOINT ["/bin/run.sh"]
