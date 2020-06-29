WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
#echo ${WORKINGDIR}
${WORKINGDIR}/go-ethereum/build/bin/geth --syncmode "light" --datadir ${WORKINGDIR}/paralleltestwork1/ --networkid 940625 --nodiscover --port 30303 --rpcport 8544 --unlock 0,1 --password ${WORKINGDIR}/samples/password --allow-insecure-unlock console

