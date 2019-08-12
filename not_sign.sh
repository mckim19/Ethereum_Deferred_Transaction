cd go-ethereum
make clean
make geth
cd ..
rm -rf paralleltestwork
mkdir paralleltestwork
cp -r keystore/ paralleltestwork/
geth init genesis.json --datadir paralleltestwork/
