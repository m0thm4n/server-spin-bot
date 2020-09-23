############################################################
# Dockerfile that builds a Soace Engineers Gameserver
############################################################
FROM cm2network/steamcmd:root

LABEL maintainer="nathan.moritz@protonmail.com"

ENV STEAMAPPID 298740
ENV STEAMAPP space-engineers
ENV STEAMAPPDIR "${HOMEDIR}/${STEAMAPP}-dedicated"
ENV DLURL https://github.com/m0thm4n/server-spin-bot

RUN set -x
RUN apt update
RUN apt upgrade -y
RUN apt install -y --no-install-recommends --no-install-suggests wget=1.20.1-1.1
RUN wget "${DLURL}/tree/master/scripts/entry.sh" -O "${HOMEDIR}/entry.sh"
RUN mkdir -p "${STEAMAPPDIR}"
RUN chmod 755 "${HOMEDIR}/entry.sh" "${STEAMAPPDIR}"
RUN chown "${USER}:${USER}" "${HOMEDIR}/entry.sh" "${STEAMAPPDIR}"
RUN rm -rf /var/lib/apt/lists/*

ENV PORT=27016 \
        QUERYPORT=27165 \
        RCONPORT=21114 \
        FIXEDMAXPLAYERS=80 \
        RANDOM=NONE

USER ${USER}

WORKDIR ${HOMEDIR}

ENTRYPOINT "${HOMEDIR}/entry.sh"

# Expose ports
EXPOSE 27016/udp \
	27165/tcp \
	27165/udp \
	21114/tcp \
	21114/udp