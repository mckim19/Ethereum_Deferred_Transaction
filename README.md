## ethereum_parallel_execution

이더리움의 실행 부분을 병렬로 처리하는 것이 목적입니다.

operating system: 18.04.02 LTS live-server, golang: >=1.10
    
## 1. geth, solidity 설치 -> github 폴더에서 시작함을 전제

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
$ git clone https://github.com/mckim19/Ethereum_Deferred_Transaction.git
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
$ cd solidity
$ ./scripts/install_deps.sh
$ sudo apt remove --purge libz3-dev
$ mkdir build
$ cd build
$ cmake .. && make
$ cp solc/solc ../../../sol_file
```

## 2. 이더리움 테스트 환경 구축
단순히 병렬 처리가 가능한지를 보기 위한 것이므로 no-discover 옵션으로 private network를 구축해서 사용한다. Network Id는 940625로 사용한다.

### A.	이더리움 데이터 폴더 및 계좌 생성
데이터 폴더(workspace) 생성
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
puppeth 모듈이란 genesis.json 파일을 생성하는 모듈이다.(puppeth 모듈을 사용하여 genesis 파일 생성하는 것이 필요하면 사용, 아니면 github.com에 예제 genesis.json 파일 사용)
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
추후 password 파일을 생성하여 geth 실행 시 password를 주는 파일로 사용한다.
## 3. 이더리움 콘솔 명령어 모음
### A. sol 파일 컴파일을 통한 컨트랙트 abi, bytecode 도출
```
$ cd solidity_file
$ ./solc --abi voting_v2.sol
$ ./solc --bin voting_v2.sol
```
### B.	Geth init 및 run
스크립트를 실행하기 전 스크립트 안의 geth 실행파일 경로를 자신의 경로로 바꿔줌
3-C~3-H까지는 run.sh를 통해 실행된 geth의 콘솔에서 진행. 이는 ">" 기호로 표시
```
$ ./init.sh
$ ./run.sh
```
### C. 컨트랙트 배포
abi와 bytecode는 3-A를 통해 알아냄
```
> var abi = <abi>
> var bytecode = <bytecode>
> var contract = eth.contract(abi)
> var deploy = {from:eth.coinbase, data:bytecode, gas: 2000000}
> var instance = contract.new("DISQUALIFIED!", deploy)
```
### D. 배포된 컨트랙트 인스턴스 생성
abi는 배포된 컨트랙트의 abi이고 address는 배포된 컨트랙트의 주소
```
> var abi = <abi>
> var address = <contract address>
> var instance = eth.contract(abi).at(address)
```
### E. 컨트랙트 트랜잭션 호출
3-C 또는 3-D로부터 생성된 컨트랙트 인스턴스를 사용하여 트랜잭션 호출

```
> instance.add.sendTransaction({from: eth.accounts[0]})
> instance.read.call()
```
### F. 어카운트 관련
```
> personal.newAccount(“2523”)
> eth.accounts / personal.listAccounts
> eth.getBalance(eth.accounts[0])
```
### G.	마이닝 관련
```
> eth.coinbase
> miner.setEtherbase(eth.accounts[3])
> miner.start(1)
> miner.stop()
> eth.mining
> eth.blockNumber
```
### H.	트랜잭션 및 블록 정보 조회
```
$ eth.pendingTransactions
$ eth.getBlock("블록 번호")
$ eth.getTransaction("트랜잭션 주소")
$ eth.getTransactionReceipt("트랜잭션 주소") //배포한 contract의 주소를 보기 위해 주로 사용
$ eth.getCode("컨트랙트 주소") //배포한 contract의 바이트코드를 확인하기 위해 주로 사용
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
$ git add <files>
$ git commit -m “added ~~”
$ git push -u origin master
```
git 최신버전 가져오기
```
$ git pull
```
## 5. 유용 링크
```
1. CRDT: Conflict-free Replicated Data Types(modification에 대해 락을 사용하지 않고 프로토콜만을 사용하여 동기화가 가능하게 하는 기법 중의 하나)
https://medium.com/@amberovsky/crdt-conflict-free-replicated-data-types-b4bfc8459d26
```
