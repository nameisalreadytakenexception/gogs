cp scripts/init/centos/gogs /etc/rc.d/init.d/
useradd -r git
service gogs start
service gogs enable #??
