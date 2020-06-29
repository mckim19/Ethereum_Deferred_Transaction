WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
DATADIR=${WORKINGDIR}/paralleltestwork1
NETWORKID=940625
PORT=30304
RPCPORT=8546

${WORKINGDIR}/go-ethereum/build/bin/geth --syncmode "light" --datadir ${DATADIR} --networkid ${NETWORKID} --nodiscover --port ${PORT} --rpcport ${RPCPORT} --unlock 0 --password ${DATADIR}/scripts/password --allow-insecure-unlock console

