#!/bin/bash
bash "${STEAMCMDDIR}/steamcmd.sh +login anonymous +force_install_dir "${STEAMAPPDIR}" +app_update "${STEAMAPPID}" +quit"

# Change rcon port on first launch, because the default config overwrites the commandline parameter (you can comment this out if it has done it's purpose)
sed -i -e 's/Port=21114/'"Port=${RCONPORT}"'/g' "${STEAMAPPDIR}/SEGame/ServerConfig/Rcon.cfg"

wine "${STEAMAPPDIR}/SpaceEngineersDedicated.exe"
bash 	Port="${PORT}" \
		QueryPort="${QUERYPORT}" \
		RCONPORT="${RCONPORT}" \
		FIXEDMAXPLAYERS="${FIXEDMAXPLAYERS}" \
		RANDOM="${RANDOM}"

# bash "${STEAMAPPDIR}/SeGameServer.sh" \
# 			Port="${PORT}" \
# 			QueryPort="${QUERYPORT}" \
# 			RCONPORT="${RCONPORT}" \
# 			FIXEDMAXPLAYERS="${FIXEDMAXPLAYERS}" \
# 			RANDOM="${RANDOM}"