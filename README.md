## ethereum_parallel_execution

이더리움의 실행 부분을 병렬로 처리하는 것이 목적입니다.

operating system: 18.04.02 LTS live-server
golang: >=1.10
    
## 1. geth, solidity 설치 -> github을 clone한 기본폴더에서 시작함을 전제

### A. Go 설치
golang 설치. golang version은 1.10 이상이다.
```
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt-get update
$ sudo apt-get install golang-go
```
편의를 위해 환경설정을 등록.
```
$ cd $home
$ echo "PATH=\$PATH:/home/`logname`/ethereum_parallel_execution/go-ethereum/build/bin" >> ~/.bashrc
$ source .bashrc
```
### B. github clone
```
$ git clone <github address>
```
### C. geth 컴파일
```
$ cd go-ethereum
$ make all
혹은 geth만 빌드하고 싶으면
$ make geth
```
### D. solidity 라이브러리 컴파일
solc 컴파일러를 컨트랙트 코드가 위치한 폴더에 두어 해당 폴더에서 컴파일하도록 함. 환경변수 등록하여 사용하여도 무방
```
$ cd ethereum_parallel_execution/solidity
$ ./scripts/install_deps.sh
$ sudo apt remove --purge libz3-dev
$ mkdir build
$ cd build
$ cmake .. && make
$ cp solc/solc ../../../sol_file
```

## 2. 이더리움 테스트 환경 구축
단순히 병렬 처리가 가능한지를 보기 위한 것이므로 no-discover 옵션으로 public network를 구축해서 사용한다. Network Id는 940625로 사용한다.
```
참고 사이트 1 - geth 컨트랙트 호출: https://stackoverflow.com/questions/48184969/calling-smart-contracts-methods-using-web3-ethereum?rq=1
```
### A.	이더리움 데이터 폴더 및 계좌 생성
데이터 폴더(workspace) 생성(github에 올라와있는 paralleltestwork는 무시해도 무방)
```
$ cd go-ethereum
$ mkdir paralleltestwork
```
계좌 생성: 미리 계좌를 생성하여 genesis 파일에서 초기에 코인을 보유할 수 있게 함.
```
geth --datadir paralleltestwork/ account new
```
적당한 비밀번호를 기억한다. Public address of the key 값이 출력이 되면 기억한다.
```
예시: Passphrase-> 2523, Public address of the key-> 0xcb2940b6766Dd4bfFF30616e4e1d3e911C8d803e
```
### B.	puppeth 모듈을 사용한 genesis.json 생성
puppeth 모듈이란 genesis.json 파일을 생성하는 모듈이다.(만약 필요하면 사용하고 github.com에 예제 genesis.json 파일이 올라와 있어 이를 사용해도 무방)
```
$ cd go-ethereum
$ puppeth
$ cp <networkname>.json genesis.json
```
```
예시: network name=yoomeetestnet, what would you do=2, what would you do=1, which consensus engine=1, 
which accounts should be pre-funded=미리 생성해 놓은 계좌의 public key를 복사하여 넣음, 
pre-funded with 1 wei=yes, chain/network ID=940625, what would you like to do=2,2.
```
cp 명령어를 통해 genesis.json 파일이 생성되면 genesis.json 파일을 수정하여 추가 설정을 완료한다.
```
예시: difficulty=0x0100, balance="0x200000000000000000000"
```
## 3. 이더리움 콘솔 명령어 모음
### A.	Geth 실행 옵션
```
$ geth --datadir paralleltestwork/ init genesis.json
$ geth --datadir paralleltestwork/ --networkid 940625 --rpc --rpcaddr "0.0.0.0" \
--rpcport 8600 --rpccorsdomain "*" --rpcapi "admin,db,eth,debug,miner,net,shh,txpool,personal,web3" \
--allow-insecure-unlock --nodiscover --port 30303 --unlock 0,1 --password password console
```
### B.	어카운트 관련
```
$ personal.newAccount(“2523”)
$ eth.accounts / personal.listAccounts
$ eth.getBalance(eth.accounts[0])
```
### C.	마이닝 관련
```
$ eth.coinbase
$ miner.setEtherbase(eth.accounts[3])
$ miner.start(1)
$ miner.stop()
$ eth.mining
$ eth.blockNumber
```
### D.	트랜잭션 전송(contract을 거치지 않은 경우)
```
$ personal.unlockAccount(eth.accounts[0])
$ eth.sendTransaction({from:eth.accounts[0], to:eth.accounts[2], value:10000})
```
### E.	트랜잭션 및 블록 정보 조회
```
$ eth.pendingTransactions
$ eth.getBlock(22)
$ eth.getTransaction("트랜잭션 주소")
$ eth.getTransactionReceipt("트랜잭션 주소") //배포한 contract의 주소를 보기 위해 주로 사용
$ eth.getCode("컨트랙트 주소") //배포한 contract의 바이트코드를 확인하기 위해 주로 사용
```
### F. geth를 통한 Contract 생성방법 (TODO: 편의를 위해 javascript(+nodejs)로 변경함, update 필요)
```
remix를 사용하여 contract를 deploy하면 deploy 할 때마다 라이브러리 contract도 계속 새로 생성되는 단점이 존재한다. 
mutex 컨트랙트가 난무하는 것을 방지하기 위해 geth에서 contract을 deploy하는 것을 선택하였다.
준비사항은 원하는 위치에 .sol 소스파일을 위치하는 것이다. 편의를 위해 paralleltestwork(작업폴더)안에 solidity_file이란 
폴더를 생성하여 컴파일을 원하는 소스코드와 라이브러리 소스코드를 같이 위치시킨다.
```
```
$ cd ~/paralleltestwork/solidity_file
$ solc --abi voting_v2.sol
$ solc --bin voting_v2.sol
$ var contract = eth.contract([{"constant":false,"inputs":[{"name":"candidate_num","type":"uint256"}],"name":"vote","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"candidate_num","type":"uint256"}],"name":"get_candidate","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"get_v","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}])
$ var bytecode = '0x608060405234801561001057600080fd5b5061048f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630121b93f1461004657806395fc395d1461008c578063fbfbfd97146100ce575b600080fd5b6100726004803603602081101561005c57600080fd5b81019080803590602001909291905050506100ec565b604051808215151515815260200191505060405180910390f35b6100b8600480360360208110156100a257600080fd5b8101908080359060200190929190505050610439565b6040518082815260200191505060405180910390f35b6100d6610450565b6040518082815260200191505060405180910390f35b600073efe38f307df41975ba058dfe2824ed53dd36be00637308809e60056040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561013e57600080fd5b505af4158015610152573d6000803e3d6000fd5b505050506005600354106101d15773efe38f307df41975ba058dfe2824ed53dd36be0063b5fba83d60056040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156101b057600080fd5b505af41580156101c4573d6000803e3d6000fd5b5050505060009050610434565b60036000815480929190600101919050555073efe38f307df41975ba058dfe2824ed53dd36be0063b5fba83d60056040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561023357600080fd5b505af4158015610247573d6000803e3d6000fd5b5050505073efe38f307df41975ba058dfe2824ed53dd36be00637308809e60066040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561029b57600080fd5b505af41580156102af573d6000803e3d6000fd5b50505050600560035414156102da576001600460006101000a81548160ff0219169083151502179055505b73efe38f307df41975ba058dfe2824ed53dd36be0063b5fba83d60066040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561032a57600080fd5b505af415801561033e573d6000803e3d6000fd5b5050505073efe38f307df41975ba058dfe2824ed53dd36be00637308809e60076040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561039257600080fd5b505af41580156103a6573d6000803e3d6000fd5b505050506001600083600381106103b957fe5b016000828254019250508190555073efe38f307df41975ba058dfe2824ed53dd36be0063b5fba83d60076040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561041757600080fd5b505af415801561042b573d6000803e3d6000fd5b50505050600190505b919050565b600080826003811061044757fe5b01549050919050565b600060035490509056fea265627a7a7230582042e82287105bbb594545c04327f3a4d14702d2c56f6fe113e6c480cc6bddfa8d64736f6c63430005090032'
$ var deploy = {from:eth.coinbase, data:bytecode, gas: 2000000}
$ var contract_instance = contract.new("DISQUALIFIED!", deploy)
```
### G. contract 함수 호출 (TODO: 편의를 위해 javascript(+nodejs)로 변경함, update 필요)
#### i. contract 객체 생성
contract 함수를 호출하기 위해서는 contract의 abi와 컨트랙트 주소가 필요하다.
여기서는 vote 컨트랙트를 사용하였다. abi와 address는 6번을 통해 알아내야 한다.
```
$ abi = [{"constant":false,"inputs":[{"name":"candidate_num","type":"uint256"}],"name":"vote","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"candidate_num","type":"uint256"}],"name":"get_candidate","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"get_v","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]
$ c_address = "0xcef3434d109bb33b9dca073c4970ed174318eb0a"
$ c_instance = eth.contract(abi).at(c_address)
```
#### ii. call
읽기만 하는 함수의 경우 geth의 call을 사용하여 호출한다. 트랜잭션을 발생시키지 않는다.
위에서 생성한 컨트랙트 객체를 가져다가 사용하였다.
```
$ c_instance.get_v.call()
$ c_instance.get_candidate.call()
```
#### iii. sendTransaction
state를 변경시키는 경우 sendTransaction을 호출하여 트랜잭션을 생성해준다. 마이닝이 된 후 결과가 반영된다.
```
$ c_instance.vote.sendTransaction(0,{from: eth.accounts[0]})
```
## 4. GIT 사용법
### A.	GIT 제공방법 – gitlab을 사용할 것임
Git 설치 및 초기 설정
```
$ sudo apt-get install git
$ git config --global user.name "John Doe"
$ git config --global user.email johndoe@example.com
```
Gitlab 시 ssh로 접근하는 것이 편하므로 내 가상머신(컴퓨터)에서 ssh 키 생성 후 gitlab에 등록해줌
```
참고 - https://dejavuqa.tistory.com/139
```
```
$ ssh-keygen -t rsa -C "GitLab" -b 4096
```
gitlab.com으로 들어가 로그인 후 프로젝트를 위한 git repository 생성. 생성했으면 git repository를 ssh 버전으로 git clone
```
$ git clone git@gitlab.com:yoomeeko/ethereum_parallel_execution.git
$ mv go_ethereum ethereum_parallel_execution.git/go_ethereum
$ mv genesis.json ethereum_parallel_execution.git/genesis.json
$ git add *
$ git commit -m “Ethereum 추가”
$ git push
```
### B. git 명령어
git 수정 후 commit 및 push 방법
```
$ git add *
$ git commit -m “added ~~”
$ git push
```
git 최신버전 가져오기
```
$ git pull
```

## Block explorer
   block explorer는 블록정보, 블록 안에 담겨 있는 트랜잭션 정보, account 정보를 ui로 예쁘게 볼 수 있는 툴을 말한다. 
   아직 이더리움에서 private chain을 위해 공식적으로 지원하는 툴은 없지만, 오픈소스가 굉장히 많다. 
   우리는 그중에서도 Carsenk의 오픈소스 explorer를 사용할 것이다.
   환경 설정하는 데 에러가 나서 실제로는 못 실행해 보았지만, 나중에 이용하면 좋을 것 같다.
   ```
   참고 github: https://github.com/carsenk/explorer
   ```
## 분산락 관련 링크
```
1. CRDT: Conflict-free Replicated Data Types(modification에 대해 락을 사용하지 않고 프로토콜만을 사용하여 동기화가 가능하게 하는 기법 중의 하나)
https://medium.com/@amberovsky/crdt-conflict-free-replicated-data-types-b4bfc8459d26
```
