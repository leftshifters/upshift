brew install infer

wget https://github.com/facebook/infer/releases/download/v0.8.1/infer-linux64-v0.8.1.tar.xz
tar xf infer-linux64-v0.8.1.tar.xz
mkdir -r /usr/local/infer
mv infer-linux64-v0.8.1.tar.xz /usr/local/infer

sudo apt-get update
sudo apt-get upgrade

sudo apt-get install -y autoconf automake build-essential git libgmp-dev libmpc-dev libmpfr-dev m4 oracle-java7-installer python-software-properties unzip zlib1g-dev




opam update
opam pin add --yes --no-action infer .
opam install --deps-only infer


opam depext conf-pkg-config.1.0