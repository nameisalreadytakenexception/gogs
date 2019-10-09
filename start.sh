sudo chmod +x scripts/init/centos/gogs
sudo mkdir -p /home/gogs/log
sudo mkdir -p /home/gogs/gogs-repositories
sudo mkdir -p lock/subsys/
sudo cp scripts/init/centos/gogs /etc/rc.d/init.d/
sudo useradd git
sudo chown -R git /var/app/
sudo chown -R git /var/app/current/custom/conf/app.ini
sudo chown -R git /home/gogs/
sudo service gogs start

sudo touch /home/gogs/starttt.sh
sudo echo "#!/bin/sh" > /home/gogs/starttt.sh
sudo echo "sudo service gogs start" >> /home/gogs/starttt.sh
#sudo echo "sudo chown -R git /var/app/current/custom/conf/app.ini" >> /home/gogs/starttt.sh
sudo echo "sudo sed -i "s+http://127.0.0.1:5000+http://127.0.0.1:3000+g" /etc/nginx/conf.d/elasticbeanstalk/00_application.conf && sudo service nginx restart" >> /home/gogs/starttt.sh
sudo chmod +x /home/gogs/starttt.sh
(crontab -l 2>/dev/null; echo "* * * * * /home/gogs/starttt.sh") | crontab -
