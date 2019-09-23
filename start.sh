sudo chmod +x scripts/init/centos/gogs
sudo mkdir log
sudo mkdir -p lock/subsys/
sudo cp scripts/init/centos/gogs /etc/rc.d/init.d/
#sudo useradd -r git
sudo useradd git
sudo chown -R git /var/app/
#sudo chown -R ec2-user /var/app/
sudo service gogs start
