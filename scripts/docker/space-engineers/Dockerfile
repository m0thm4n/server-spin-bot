############################################################
# Dockerfile that builds a Soace Engineers Gameserver
############################################################
FROM cm2network/steamcmd:root

LABEL maintainer="nathan.moritz@protonmail.com"

ENV STEAMAPPID 298740
ENV STEAMAPP space-Engineers
ENV STEAMAPPDIR "${HOMEDIR}/${STEAMAPP}-didicated"
ENV DLURL https://raw.githubusercontent.com/m0thm4n/server-spin-bot

RUN set -x \
        && apt update \
        && apt install -y --no-install-recommends --no-install-suggests \ 
            wget=1.20.1-1.1 \
        && wget "${DLURL}/tree/master/scripts/entry.sh" -O "${HOMEDIR}/entry.sh" \
        && mkdir -p "${STEAMAPPDIR}" \
        && chmod 755 "${HOMEDIR}/entry.sh" "${STEAMAPPDIR}" \
        && chown "${USER}:${USER}" "${HOMEDIR}/entry.sh" "${STEAMAPPDIR}" \

ENV PORT=27016 \
        QUERYPORT=27165 \
        RCONPORT=21114 \
        FIXEDMAXPLAYERS=80 \
        RANDOM=NONE

USER ${USER}

WORKDIR ${HOMEDIR}

CMD ["bash", "entry.sh"]

# Expose ports
EXPOSE 27016/udp \
	27165/tcp \
	27165/udp \
	21114/tcp \
	21114/udp