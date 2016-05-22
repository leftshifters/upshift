# Details here - https://gitlab.com/gitlab-org/gitlab-ci-multi-runner/blob/master/docs/install/linux-repository.md
# Install runner
curl -L https://packages.gitlab.com/install/repositories/runner/gitlab-ci-multi-runner/script.deb.sh | sudo bash
sudo apt-get install gitlab-ci-multi-runner
sudo gitlab-ci-multi-runner register
# Add details from http://code.leftshift.io/ci or your gitlab installation

# Dockerfile for Android - https://gist.github.com/eluleci/f9904382c1496fb81079
apt-get update
apt-get install zip unzip
apt-get install -y curl
dpkg --add-architecture i386
apt-get update
apt-get install -y libc6:i386 libstdc++6:i386 lib32z1 libsdl1.2debian:i386
apt-get install -y git-core
touch /var/opt/env
chmod 777 /var/opt/env
echo "debconf shared/accepted-oracle-license-v1-1 select true" | /usr/bin/debconf-set-selections
echo "debconf shared/accepted-oracle-license-v1-1 seen true" | /usr/bin/debconf-set-selections
apt-get install -q -y software-properties-common
apt-get install -q -y python-software-properties
add-apt-repository ppa:webupd8team/java -y
apt-get update
echo oracle-java7-installer shared/accepted-oracle-license-v1-1 select true | /usr/bin/debconf-set-selections
apt-get install oracle-java7-installer oracle-java8-installer -y
cd /usr/local/ && wget http://dl.google.com/android/android-sdk_r24.4.1-linux.tgz && tar xf android-sdk_r24.4.1-linux.tgz
rm android-sdk_r24.4.1-linux.tgz
/usr/local/android-sdk-linux/tools/android update sdk --no-ui
rm -rf /usr/local/android-sdk_r24.2-linux.tgz
cd /usr/local/ && wget http://services.gradle.org/distributions/gradle-2.13-all.zip && unzip -o gradle-2.13-all.zip
rm -rf /usr/local/gradle-2.13-all.zip
cd && curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift && chmod +x upshift && ./upshift install
rm upshift

sudo apt-get install mailutil postfix
vi /etc/postfix/main.cf
- Add
smtp_sasl_auth_enable = yes
relayhost = smtp.mailgun.org
smtp_sasl_security_options = noanonymous
smtp_sasl_password_maps=hash:/etc/postfix/sasl_passwd
- Remove
relayhost = 

echo 'smtp.mailgun.org postmaster@vercingetorix.mailgun.org:0i0u3gpisfs5' > /etc/postfix/sasl_passwd
chmod 600 /etc/postfix/sasl_passwd
postmap /etc/postfix/sasl_passwd
service postfix restart
echo "test super test" | mail -s "test subject" "sudhanshu@leftshift.io"


vi /var/opt/env
- Add
DEBIAN_FRONTEND="noninteractive"
JAVA_HOME="/usr/lib/jvm/java-7-oracle"
JAVA8_HOME="/usr/lib/jvm/java-8-oracle"
PATH="$JAVA_HOME:$PATH"
ANDROID_HOME="/usr/local/android-sdk-linux"
PATH="$PATH:$ANDROID_HOME/tools"
PATH="$PATH:$ANDROID_HOME/platform-tools"
GRADLE_HOME="/usr/local/gradle-2.13"
PATH="$PATH:$GRADLE_HOME/bin"

vi /etc/profile
- Add
source /var/opt/env


# To find packages while are missing
android list sdk -a --no-ui
android update sdk -a --no-ui --filter 5,6,7



for remote in `git branch -r | grep -v master `; do git checkout --track $remote ; done

