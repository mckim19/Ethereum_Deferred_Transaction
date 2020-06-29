WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
#echo ${WORKINGDIR}
${WORKINGDIR}/go-ethereum/build/bin/geth --datadir ${WORKINGDIR}/paralleltestwork2/ --nousb init ${WORKINGDIR}/samples/genesis.json
