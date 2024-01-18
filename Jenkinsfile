pipeline {
    agent any

    stages {
        stage ('Diployment'){
            steps {
                script {
                    sh 'echo "Deploying..."'
                    sh 'sudo systemctl restart mobilemart.service'
                    sh 'echo "Deployment complete."'
                }
            }
        }
    }
    
}