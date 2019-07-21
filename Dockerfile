FROM plugins/base:linux-amd64

LABEL maintainer="Michael Müller <io@digitalwerber.de>" \
  org.label-schema.name="Drone Teams" \
  org.label-schema.vendor="Michael Müller" \
  org.label-schema.schema-version="1.0"

ADD ./drone-teams /bin/
ENTRYPOINT /bin/drone-teams
