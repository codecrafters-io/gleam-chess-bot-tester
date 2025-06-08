wget -O go.tar.gz https://go.dev/dl/go1.23.7.linux-amd64.tar.gz
tar -xzvf go.tar.gz -C /usr/local
export PATH=$HOME/go/bin:/usr/local/go/bin:$PATH >> ~/.profile
echo export PATH=$HOME/go/bin:/usr/local/go/bin:$PATH >> ~/.profile
source ~/.profile

sudo apt-get update && sudo apt-get install -y cmake build-essential

sudo apt-get install -y g++ make

sudo apt-get install -y ninja-build

cd vcpkg && ./bootstrap-vcpkg.sh
