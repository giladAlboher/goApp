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

        stage('dev'){
            when {
                branch 'dev'
            }
            steps {
                script {
                    echo('This stage only works with branch dev')
                }
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