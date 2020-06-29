WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
DATADIR=${WORKINGDIR}/paralleltestwork

${WORKINGDIR}/go-ethereum/build/bin/geth --datadir ${DATADIR} account new
