#!/bin/bash
bash "${STEAMCMDDIR}/steamcmd.sh" '+login anonymous' \
		'+force_install_dir' "${STEAMAPPDIR}" \
		'+app_update' "${STEAMAPPPID}" \
		'+quit'

sed -i -e 's/Port=21114/'"Port=${RCONPORT}"'/g' "$PSTEAMAPPDIR}/SEGame/ServerConfg/Rcon.cfg"

wine "${STEAMAPPDIR}/SpaceEngineersDedicated.exe"
bash	Port="${PORT}" \
	QueryPort="${QUERYPORT}" \
	RCONPORT="${RCONPORT}" \
	FIXEDMAXPLAYERS="${FIXEDMAXPLAYERS}" \
	RANDOM="${RANDOM}"
