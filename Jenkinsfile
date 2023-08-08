pipeline {
    agent any
    tool(go: 'go-1.12.7') // This assumes you have a Go installation called "go-1.12.7"
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Build') {
            steps {
                script {
                    sh 'go run main.go'
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
