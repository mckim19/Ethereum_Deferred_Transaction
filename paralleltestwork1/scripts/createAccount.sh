WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
DATADIR=${WORKINGDIR}/paralleltestwork1

${WORKINGDIR}/go-ethereum/build/bin/geth --datadir ${DATADIR} account new
