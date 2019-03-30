pipeline {
    agent any
    
    stages {
        stage ('Build') {
		steps {
		      dir ('./'){
			sh '''go build file_read.go'''
		      }
        
        		dir ('source/terraform/dev') {
				sh 'terraform init && terraform apply -auto-approve'				}
			}
		}
	}
}
