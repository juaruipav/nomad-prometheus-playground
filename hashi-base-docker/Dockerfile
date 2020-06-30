FROM ubuntu:18.04

ARG NOMAD_VER="0.11.3"
ARG CONSUL_VER="1.8.0"

RUN apt-get update && \
    apt-get install -y --no-install-recommends tar gettext-base  python3 ssh iproute2 net-tools wget unzip && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p /opt/nomad/bin && mkdir -p /opt/consul/bin \
    mkdir -p /opt/docker

########################
## NOMAD installation ##
########################
    
RUN cd /opt/nomad/bin && \
    wget --no-check-certificate https://releases.hashicorp.com/nomad/${NOMAD_VER}/nomad_${NOMAD_VER}_linux_amd64.zip && \
    unzip nomad_${NOMAD_VER}_linux_amd64.zip

########################
## CONSUL installation ##
#######################

    
RUN cd /opt/consul/bin && \
    wget --no-check-certificate https://releases.hashicorp.com/consul/${CONSUL_VER}/consul_${CONSUL_VER}_linux_amd64.zip && \
    unzip consul_${CONSUL_VER}_linux_amd64.zip


ENV PATH=$PATH:/opt/nomad/bin:/opt/docker/docker:/opt/consul/bin

