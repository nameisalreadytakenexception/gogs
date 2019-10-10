provider "aws" {
  region     = "eu-central-1"

}
resource "aws_elastic_beanstalk_application" "gogstf-TNUMBER" {
  name = "gogstf-TNUMBER"
}
resource "aws_security_group" "gogssg-TNUMBER" {
  name        = "gogssg-TNUMBER"
  description = "Allow inbound traffic from provided Security Groups"
  vpc_id      = "vpc-0b6f211bbe6b29305"
  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "TCP"
    #security_groups = ["gogssg"]
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = 80 #3000
    to_port   = 80 #3000
    protocol  = "TCP"
    #security_groups = ["gogssg"]
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = 3306
    to_port   = 3306
    protocol  = "TCP"
    #security_groups = ["gogssg"]
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    ita_group = "Lv-428"
  }
}
resource "aws_elastic_beanstalk_environment" "gogs" {
  name                = "gogsenvtf-TNUMBER"
  application         = "${aws_elastic_beanstalk_application.gogstf-TNUMBER.name}"
  solution_stack_name = "64bit Amazon Linux 2018.03 v2.13.0 running Go 1.13"
  tags = {
    ita_group = "Lv-428"
  }
  setting {
    namespace = "aws:ec2:vpc"
    name      = "VPCId"
    value     = "vpc-0b6f211bbe6b29305"
  }
  setting {
    namespace = "aws:ec2:vpc"
    name      = "Subnets"
    value     = "subnet-0336e2c2d9e698fb8"
  }
  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "SecurityGroups"
    value     = "${aws_security_group.gogssg-TNUMBER.id}"
  }
  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "InstanceType"
    value     = "t2.micro"
  }
  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "EC2KeyName"
    value     = "gogs_key_pair"
  }
  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "IamInstanceProfile"
    value     = "aws-elasticbeanstalk-ec2-role"
  }
  setting {
    namespace = "aws:elasticbeanstalk:environment"
    name      = "ServiceRole"
    value     = "aws-elasticbeanstalk-service-role"
  }
#   provisioner "local-exec" {
#     command = "sudo touch /etc/passwd-s3fs"
#   }
#   provisioner "local-exec" {
#     command = "sudo bash -c 'echo "AWS_ACCESS_KEY_ID:AWS_SECRET_ACCESS_KEY" > /etc/passwd-s3fs'"
#   }
  provisioner "local-exec" {
    command = "sudo chmod 640 /etc/passwd-s3fs"
  }
  provisioner "file" {
    source      = "passwd-s3fs"
    destination = "/etc/passwd-s3fs"
  }
}
