## ethereum_parallel_execution

이더리움의 실행 부분을 병렬로 처리하는 것이 목적입니다.

환경

    18.04.02 LTS live-server
    
## 이더리움 빌드방법

### 1. Go 설치
#### A.	https://github.com/golang/go/wiki/Ubuntu
```
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt-get update
$ sudo apt-get install golang-go
```
#### B.	편의를 위해 환경설정을 등록

source 명령어는 재부팅을 하지 않는 이상 각 쉘마다 다시 실행해야 한다.
```
$ cd $home
$ echo "PATH=\$PATH:/home/`logname`/ethereum_parallel_execution/go-ethereum/build/bin" >> ~/.bashrc
$ source .bashrc
```
### 2.	Ethereum Network Stats 설치
Eth-netstats는 이더리움 네트워크 상태를 추적하기 위한 인터페이스로 인터페이스는 웹으로 실행되며, 이더리움 노드와 통신하기 위하여 웹소켓을 사용한다. 따라서 노드에서 eth-netstats에 등록하기 위하여 서로 약속된 WS_SECRET이라는 환경변수를 이용한다. 만약 가상머신을 사용하고 있다면 포트 포워딩을 해줘야 한다. 가상머신 네트워크 어댑터에서 내부 3000번 포트와 연결해주는 외부 port를 설정해준다. 만약 어떻게 하는지 모르면 인터넷을 검색하거나 다른 사람에게 물어봐라.
```
$ cd $home
$ sudo apt install npm node-grunt-cli
$ git clone https://github.com/cubedro/eth-netstats
$ cd ~/eth-netstats
$ npm install
$ grunt   //java script 빌드.. 소스를 수정한 후 grunt를 해주면 실시간 반영됨
```
### 3.	이더리움 소스코드 다운로드 및 초기 컴파일 진행
```
$ cd $home
$ git clone https://github.com/ethereum/go-ethereum
$ cd go-ethereum
$ make geth
```
## 이더리움 테스트 환경 구축
단순히 병렬 처리가 가능한지를 보기 위한 것이므로 no-discover 옵션으로 public network를 구축해서 사용한다. Network Id는 940625로 사용한다.
```
참고 사이트 1 - geth 컨트랙트 호출: https://stackoverflow.com/questions/48184969/calling-smart-contracts-methods-using-web3-ethereum?rq=1
```
### 1.	이더리움 개인 네트워크 구축 준비
#### A.	데이터 폴더(workspace) 생성 및 계좌 생성
Account를 생성하라 그러면 Passphrase를 요구한다. 적당한 비밀번호로 설정하고 기억한다. Public address of the key 값이 출력이 되면 기억한다.
```
예시: Passphrase-> 2523, Public address of the key-> 0xcb2940b6766Dd4bfFF30616e4e1d3e911C8d803e
```
```
$ cd $home
$ mkdir paralleltestwork
$ geth --datadir paralleltestwork/ account new
```
#### B.	puppeth 모듈을 사용한 genesis.json 생성
만약 source .bashrc가 잘 적용이 되었다면 puppeth가 잘 동작할 것이다. puppeth를 실행하고 명령어대로 따르면 된다.
```
예시: network name=yoomeetestnet, what would you do=2, what would you do=1, which consensus engine=1, 
which accounts should be pre-funded=미리 생성해 놓은 계좌의 public key를 복사하여 넣음, 
pre-funded with 1 wei=yes, chain/network ID=940625, what would you like to do=2,2.
```
cp 명령어를 통해 genesis.json 파일이 생성되면 genesis.json 파일을 수정하여 추가 설정을 완료한다.
```
예시: difficulty=0x0100, balance="0x200000000000000000000"
```
```
$ puppeth
$ cp <networkname>.json genesis.json
```
### 2.	재부팅 시 이더리움 실행
```
$ cd ~/eth-netstats
$ nohup env WS_SECRET=Hello npm start & //백그라운드로 nohup(화면 없이) 실행
$ netstat -na | grep tcp | grep 3000 //netstat은 3000번 포트로 열림
$ cd $home
$ --ethstats yoom:Hello@localhost:3000	\
```
## 이더리움 콘솔 명령어 모음
### 1.	Geth 실행 옵션
```
$ geth --datadir paralleltestwork/ init genesis.json
$ geth --datadir paralleltestwork/ --networkid 940625 --rpc --rpcaddr "0.0.0.0" 
--rpcport 8600 --rpccorsdomain "*" --rpcapi "admin,db,eth,debug,miner,net,shh,txpool,personal,web3" 
--allow-insecure-unlock --nodiscover --port 30303 --unlock 0,1 --password password console
```
### 2.	어카운트 관련
```
$ personal.newAccount(“2523”)
$ eth.accounts / personal.listAccounts
$ eth.getBalance(eth.accounts[0])
```
### 3.	마이닝 관련
```
$ eth.coinbase
$ miner.setEtherbase(eth.accounts[3])
$ miner.start(1)
$ miner.stop()
$ eth.mining
$ eth.blockNumber
```
### 4.	트랜잭션 전송(contract을 거치지 않은 경우)
```
$ personal.unlockAccount(eth.accounts[0])
$ eth.sendTransaction({from:eth.accounts[0], to:eth.accounts[2], value:10000})
```
### 5.	트랜잭션 및 블록 정보 조회
```
$ eth.pendingTransactions
$ eth.getBlock(22)
$ eth.getTransaction("트랜잭션 주소")
$ eth.getTransactionReceipt("트랜잭션 주소") //배포한 contract의 주소를 보기 위해 주로 사용
$ eth.getCode("컨트랙트 주소") //배포한 contract의 바이트코드를 확인하기 위해 주로 사용
```
### 6. geth를 통한 Contract 생성방법
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
### 7. contract 함수 호출
#### A. contract 객체 생성
contract 함수를 호출하기 위해서는 contract의 abi와 컨트랙트 주소가 필요하다.
여기서는 vote 컨트랙트를 사용하였다. abi와 address는 6번을 통해 알아내야 한다.
```
$ abi = [{"constant":false,"inputs":[{"name":"candidate_num","type":"uint256"}],"name":"vote","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"candidate_num","type":"uint256"}],"name":"get_candidate","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"get_v","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]
$ c_address = "0xcef3434d109bb33b9dca073c4970ed174318eb0a"
$ c_instance = eth.contract(abi).at(c_address)
```
#### B. call
읽기만 하는 함수의 경우 geth의 call을 사용하여 호출한다. 트랜잭션을 발생시키지 않는다.
위에서 생성한 컨트랙트 객체를 가져다가 사용하였다.
```
$ c_instance.get_v.call()
$ c_instance.get_candidate.call()
```
#### C. sendTransaction
state를 변경시키는 경우 sendTransaction을 호출하여 트랜잭션을 생성해준다. 마이닝이 된 후 결과가 반영된다.
```
$ c_instance.vote.sendTransaction(0,{from: eth.accounts[0]})
```
## GIT 사용법
### 1.	GIT 제공방법 – gitlab을 사용할 것임
#### A.	Git 설치 및 초기 설정
```
$ sudo apt-get install git
$ git config --global user.name "John Doe"
$ git config --global user.email johndoe@example.com
```
#### B.	Gitlab 시 ssh로 접근하는 것이 편하므로 내 가상머신(컴퓨터)에서 ssh 키 생성 후 gitlab에 등록해줌
```
참고 - https://dejavuqa.tistory.com/139
```
```
$ ssh-keygen -t rsa -C "GitLab" -b 4096
```
#### C.	gtlab.com으로 들어가 로그인 후 프로젝트를 위한 git repository 생성. 생성했으면 git repository를 ssh 버전으로 git clone
```
$ git clone git@gitlab.com:yoomeeko/ethereum_parallel_execution.git
$ mv go_ethereum ethereum_parallel_execution.git/go_ethereum
$ mv genesis.json ethereum_parallel_execution.git/genesis.json
$ git add *
$ git commit -m “Ethereum 추가”
$ git push
```
### 2.	다른 사람의 git 설치 및 clone 방법
#### A.	Git 설치 및 초기 설정
```
$ sudo apt-get install git
$ git config --global user.name "John Doe"
$ git config --global user.email johndoe@example.com
```
#### B.	Gitlab 시 ssh로 접근하는 것이 편하므로 내 가상머신(컴퓨터)에서 ssh 키 생성 후 gitlab에 등록해줌
```
참고 - https://dejavuqa.tistory.com/139
```
```
$ ssh-keygen -t rsa -C "GitLab" -b 4096
```
#### C.	git repository를 ssh 버전으로 git clone
```
$ git clone git@gitlab.com:yoomeeko/ethereum_parallel_execution.git
```
### 3.	git 수정 후 commit 및 push 방법
```
$ git add *
$ git commit -m “added ~~”
$ git push
```
### 4.	git 최신버전 가져오기
```
$ git pull
```
## 스마트 컨트랙트
### 1. mutex 라이브러리
```
pragma solidity ^0.5.4;
library mutex {
    struct mutex_v
    {
        uint L;
    }
    function lock(mutex_v storage a) public {
        while(a.L == 0){}
        a.L=1;
    }
    function unlock(mutex_v storage a) public {
        a.L=0;
    }
}
```
### 2. voting_v1.sol
```
참조 사이트1 - 라이브러리 사용법: https://solidity-kr.readthedocs.io/ko/latest/contracts.html?highlight=library#libraries
참조 사이트2 - 라이브러리 링크 방법: https://medium.com/coinmonks/all-you-should-know-about-libraries-in-solidity-dd8bc953eae7
참조 사이트3 - geth console에서 contract 생성방법: https://medium.com/mercuryprotocol/dev-highlights-of-this-week-cb33e58c745f
--> npm install -g solc가 필요함
참조 사이트4 - remix 라이브러리 생성법: https://ethereum.stackexchange.com/questions/12299/how-does-solidity-online-compiler-link-libraries
```
```
pragma solidity ^0.5.4;
//We believe voters are innocent.
import {mutex} from "./mutex.sol";
contract C {
    using mutex for *;
    uint constant POPULATION_NUM=5;
    uint[3] candidate;
    uint v;
    bool final_flag;
    mutex.mutex_v x;
    mutex.mutex_v y;
    mutex.mutex_v z;
    
    function vote(uint candidate_num) public returns (bool)
    {
        /* x,y,z does not have to store sequence! */
        /* only v have to store sequence! */
        mutex.lock(x);
        if (v >= POPULATION_NUM)
        {
            mutex.unlock(x);
            return false;
        }
        v++;
        mutex.unlock(x);
        
        mutex.lock(y);
        if(v == POPULATION_NUM)
            final_flag = true;
        mutex.unlock(y);
        
        mutex.lock(z);
        candidate[candidate_num]+=1;
        mutex.unlock(z);
        //local_op(candidate[candidate],plus,1);       
        return true;
    }
}
```
