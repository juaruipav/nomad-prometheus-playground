FROM hashi-base-image:0.1.0

COPY nomadServer.hcl .
COPY consulServer.hcl .

RUN mkdir -p /opt/consul
COPY launchServices.sh .
RUN chmod +x launchServices.sh
CMD ./launchServices.sh
