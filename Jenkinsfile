pipeline {
    agent any

    tools{
        go 'go'
    } // This assumes you have a Go installation called "go-1.12.7"
    
   
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('only works with branch dev'){
            echo(message: 'This stage only works with branch dev', color: 'RED')
            when {
                branch 'dev'
            }
        }
        
        stage('Build') {
            steps {
                script {
                    sh 'go build'
                }
            }
        }
        
        stage('Test') {
            steps {
                script {
                    sh 'go test'
                }
            }
        }
    }
}