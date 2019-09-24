sudo chmod +x scripts/init/centos/gogs
sudo mkdir -p /home/gogs/log
sudo mkdir -p /home/gogs/gogs-repositories
sudo mkdir -p lock/subsys/
sudo cp scripts/init/centos/gogs /etc/rc.d/init.d/
sudo useradd git
sudo usermod -aG wheel git
sudo sed -i "s/# %wheel/ %wheel/g" /etc/sudoers
sudo chown -R git /var/app/
sudo chown -R git /var/app/current/custom/conf/app.ini
sudo chown -R git /home/gogs/
sudo service gogs start
