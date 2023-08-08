pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Build') {
            steps {
                script {
                    sh 'go build -o helloWorldApp'
                }
            }
        }
        
        stage('Test') {
            steps {
                script {
                    sh 'go test ./...'
                }
            }
        }
        
        stage('Deploy') {
            steps {
                script {
                    // Assuming you have some deployment steps here
                    // like copying the binary to a server or container
                    // and starting the application
                    sh './helloWorldApp'
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
