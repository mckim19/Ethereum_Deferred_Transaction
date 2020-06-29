WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
#echo ${WORKINGDIR}
${WORKINGDIR}/go-ethereum/build/bin/geth --syncmode "full" --lightserv 20 --datadir ${WORKINGDIR}/paralleltestwork2/ --networkid 940625 --nodiscover --port 30304 console
