###################################
# variables
###################################

variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "private_key_path" {}
variable "key_name" {
 default="ansible"
}

##################################
# Providers
##################################

provider "aws"{
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "us-east-2"
}
##################################
# Resources
##################################
resource "aws_instance" "toyota-demo"{

 ami          = "ami-0c55b159cbfafe1f0"
 instance_type= "t2.micro"
 associate_public_ip_address = true
 key_name     = "${var.key_name}"

 connection {
	user = "ubuntu"
	private_key = "${file(var.private_key_path)}"
	}
	provisioner "remote-exec" {
          inline = [
	     "sudo add-apt-repository ppa:longsleep/golang-backports",
	     "sudo apt-get update",
	     "sudo apt-get install golang-go"
          ]
       }
   }
####################################
# OUTPUT
####################################

output "aws_instance_public_ip"{
            value ="${aws_instance.toyota-demo.public_ip}"
 }
