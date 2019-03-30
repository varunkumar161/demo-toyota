pipeline {
    agent any
    
    stages {
        stage ('Build') {
		steps {
		      dir ('./'){
			sh '''go build file_read.go'''
		      }
        
			}
		}
	}
}
