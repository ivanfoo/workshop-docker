FROM mongo:3.4

MAINTAINER genDevOps

# Add config
ADD config/mongod.conf /etc/mongod.conf

# Data volume
RUN mkdir /db && chmod 777 -R /db
VOLUME ["/db"]

# Ports
EXPOSE 27000

# RUN
CMD ["--config", "/etc/mongod.conf"]

