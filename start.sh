sudo chmod +x scripts/init/centos/gogs
sudo mkdir log
sudo mkdir -p lock/subsys/
sudo cp scripts/init/centos/gogs /etc/rc.d/init.d/
sudo useradd -r git
sudo chown -R git /var/app/
service gogs start
service gogs enable #??
