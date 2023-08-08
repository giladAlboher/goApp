pipeline {
    agent any
    tool{
        go 'go'
    } // This assumes you have a Go installation called "go-1.12.7"
    stages {
        stage('Checkout') {
            steps {
                checkout scm
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
    
    post {
        always {
            cleanWs() // Clean up the workspace
        }
    }
}
