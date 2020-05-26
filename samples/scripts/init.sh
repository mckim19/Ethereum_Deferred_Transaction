WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
#echo ${WORKINGDIR}
${WORKINGDIR}/go-ethereum/build/bin/geth --datadir ${WORKINGDIR}/paralleltestwork/ --nousb init ${WORKINGDIR}/samples/genesis.json
