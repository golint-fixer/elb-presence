# ELB Presence

A lightweight utility for Elastic Load-Balancer membership. It is most useful for pairing with a Docker container to announce the inclusion of the host instance within the specified Elastic Load-Balancer.

This utility takes inspiration from the CoreOS utility of the [same name](https://github.com/coreos/elb-presence). The CoreOS utility is a full Ubuntu distribution, a simple Python script utilizing the Boto library with the image clocking in at 211MB. This utility and its acompanying Docker image utilizes Busybox and a pre-compiled Go binary clocking in at around 9.6MB.

# Usage

The best place to start is the [CoreOS example deployment with Fleet](https://coreos.com/docs/launching-containers/launching/fleet-example-deployment/). Whereever you would use the CoreOS elb-presence sidekick service unit you can utilize this utility.

# Configuration

The container expects three environment variables to be set.
- `AWS_ACCESS_KEY` - AWS Access Key with permissions to add/remove instances from the designated Elastic Load-Balancer.
- `AWS_SECRET_KEY` - AWS Secret Key with permissions to add/remove instances from the designated Elastic Load-Balancer.
- `ELB_NAME` - The name of your Elastic Load-Balancer.

Example Docker run:
- `$ docker build -t elb-presence .`
- `$ docker run -e "AWS_ACCESS_KEY=XXXXXXXXXXXXXXXXXXXX" -e "AWS_SECRET_KEY=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx" -e "ELB_NAME=your-elb-name" elb-presence`
