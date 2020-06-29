WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
DATADIR=${WORKINGDIR}/paralleltestwork2
NETWORKID=940625
PORT=30305
RPCPORT=8547

${WORKINGDIR}/go-ethereum/build/bin/geth --syncmode "light" --datadir ${DATADIR} --networkid ${NETWORKID} --nodiscover --port ${PORT} --rpcport ${RPCPORT} --unlock 0 --password ${DATADIR}/scripts/password --allow-insecure-unlock console

