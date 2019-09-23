sudo cp scripts/init/centos/gogs /etc/rc.d/init.d/
sudo useradd -r git
sudo chown -R git /var/app/current
service gogs start
service gogs enable #??
