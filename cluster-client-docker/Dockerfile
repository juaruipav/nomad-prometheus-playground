FROM hashi-base-image:0.1.0

# RUN nomad agent -config client.hcl &
COPY nomadClient.hcl.tpl .
COPY consulClient.hcl.tpl .

COPY launchServices.sh .
RUN chmod +x launchServices.sh
CMD ./launchServices.sh