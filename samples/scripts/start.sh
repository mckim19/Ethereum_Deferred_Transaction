WORKINGDIR=/home/${USER}/Ethereum_Deferred_Transaction
#echo ${WORKINGDIR}
${WORKINGDIR}/go-ethereum/build/bin/geth --datadir ${WORKINGDIR}/paralleltestwork/ --networkid 940625 --nodiscover --port 30303 --unlock 0,1,2,3 --password ${WORKINGDIR}/samples/password --rpc --rpcaddr "0.0.0.0" --rpcport 8545  --rpccorsdomain "*" --rpcapi "admin,db,eth,debug,miner,net,shh,txpool,personal,web3" --allow-insecure-unlock console

