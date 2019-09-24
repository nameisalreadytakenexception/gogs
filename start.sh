# sudo chmod +x scripts/init/centos/gogs
# sudo mkdir -p /home/gogs/log
# sudo mkdir -p /home/gogs/gogs-repositories
# sudo mkdir -p lock/subsys/
# sudo cp scripts/init/centos/gogs /etc/rc.d/init.d/
# #sudo mv app.ini custom/conf/app.ini
# sudo useradd git
# sudo chown -R git /var/app/
# sudo chown -R git /var/app/current/custom/conf/app.ini
# sudo chown -R git /home/gogs/
# sudo service gogs start

sudo chmod +x scripts/init/centos/gogs
sudo useradd git
sudo mkdir -p /home/git/gogs/log
sudo mkdir -p /home/git/gogs/gogs-repositories
sudo mkdir -p /home/git/gogs/lock/subsys/
sudo cp scripts/init/centos/gogs /etc/rc.d/init.d/
sudo cp -r /var/app/current/* /home/git/gogs
sudo service gogs start
