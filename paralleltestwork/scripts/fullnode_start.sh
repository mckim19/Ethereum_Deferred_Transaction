WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
DATADIR=${WORKINGDIR}/paralleltestwork
NETWORKID=940625
PORT=30303
RPCPORT=8545

${WORKINGDIR}/go-ethereum/build/bin/geth --syncmode "full" --maxpeers 20 --lightserv 20 --datadir ${DATADIR} --networkid ${NETWORKID} --nodiscover --port ${PORT} --rpcport ${RPCPORT} --unlock 0 --password ${DATADIR}/scripts/password --allow-insecure-unlock console
