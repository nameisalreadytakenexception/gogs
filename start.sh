sudo yum update -y
sudo yum install automake fuse fuse-devel gcc-c++ git libcurl-devel libxml2-devel make openssl-devel -y
git clone https://github.com/s3fs-fuse/s3fs-fuse.git
cd s3fs-fuse/
./autogen.sh
./configure --prefix=/usr --with-openssl
make
sudo make install
cd ..
sudo touch /etc/passwd-s3fs
sudo bash -c "echo 's3fspass' | base64 -d > /etc/passwd-s3fs"
sudo chmod 640 /etc/passwd-s3fs
sudo chmod +x scripts/init/centos/gogs
sudo mkdir -p /home/gogs/log
sudo mkdir -p /home/gogs/gogs-repositories
# sudo mkdir -p /home/gogs/gogs-repositories-s3fs
sudo mkdir -p lock/subsys/
sudo cp scripts/init/centos/gogs /etc/rc.d/init.d/
sudo useradd git
sudo chown -R git /var/app/
sudo s3fs gogs-storage -o use_cache=/tmp -o allow_other -o multireq_max=5 /home/gogs/gogs-repositories
# sudo cp /home/gogs/gogs-repositories-s3fs/* /home/gogs/gogs-repositories/
sudo chmod 
sudo chown -R git /home/gogs/gogs-repositories
sudo chown -R git /home/gogs/log
sudo bash -c 'echo "/usr/bin/s3fs gogs-storage -o use_cache=/tmp -o allow_other -o uid=1001 -o mp_umask=002 -o multireq_max=5 /home/gogs/gogs-repositories"'
sudo service gogs start

sudo touch /home/gogs/start1.sh
sudo echo "#!/bin/sh" > /home/gogs/start1.sh
sudo echo "sudo service gogs start" >> /home/gogs/start1.sh
#sudo echo "sudo chown -R git /var/app/current/custom/conf/app.ini" >> /home/gogs/start1.sh
sudo echo "sudo sed -i "s+http://127.0.0.1:5000+http://127.0.0.1:3000+g" /etc/nginx/conf.d/elasticbeanstalk/00_application.conf && sudo service nginx restart" >> /home/gogs/start1.sh
sudo chmod +x /home/gogs/start1.sh
(crontab -l 2>/dev/null; echo "* * * * * /home/gogs/start1.sh") | crontab -
# sudo touch /home/gogs/start2.sh
# sudo echo "#!/bin/sh" > /home/gogs/start2.sh
# sudo echo "sudo cp /home/gogs/gogs-repositories/* /home/gogs/gogs-repositories-s3fs/ -n" >> /home/gogs/start2.sh
# sudo echo "sudo cp /home/gogs/gogs-repositories-s3fs/* /home/gogs/gogs-repositories/ -n" >> /home/gogs/start2.sh
# sudo chmod +x /home/gogs/start2.sh
# (crontab -l 2>/dev/null; echo "* * * * * /home/gogs/start2.sh") | crontab -
