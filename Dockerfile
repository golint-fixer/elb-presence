# Docker Definition for ELB Presence Utility

FROM busybox:ubuntu-14.04
MAINTAINER Christian R. Vozar <cvozar@civisanalytics.com>

ADD cmd/elb-presence /elb-presence
RUN chmod +x /elb-presence

# Certificates here are needed for Go's crypto/x509 package which is utilized by
# the HTTPS client.
ADD ca-bundle.crt /etc/ssl/ca-bundle.pem

CMD ["/elb-presence"]
