#geth --datadir paralleltestwork/ --networkid 940625 --nodiscover --port 30303 --unlock 0,1 --password password --ethstats node1:hello@localhost:3000 console 
/home/${USER}/ethereum_parallel_execution/go-ethereum/build/bin/geth --datadir paralleltestwork/ --networkid 940625 --nodiscover --port 30303 --unlock 0,1 --password password --rpc --rpcaddr "0.0.0.0" --rpcport 8545  --rpccorsdomain "*" --rpcapi "admin,db,eth,debug,miner,net,shh,txpool,personal,web3" --allow-insecure-unlock console

